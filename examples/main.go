package main

import (
	"context"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/davecgh/go-spew/spew"
	"github.com/maxidelgado/dynagraph"
	"github.com/maxidelgado/dynagraph/pkg/utils"
)

func init() {
	os.Setenv("AWS_REGION", "us-east-1")
	os.Setenv("AWS_PROFILE", "YOUR_AWS_PROFILE")
}

type Person struct {
	Id   string
	Name string

	Addresses []Address `dynamo:"-"`
}

type Address struct {
	Id     string
	Street string
	Number int
}

func main() {
	var (
		sess = connectAWS("us-east-1")
		db   = dynamodb.New(sess)
		c, _ = dynagraph.New(db, "YOUR_TABLE_NAME")
		ctx  = context.Background()
	)

	foo := Person{
		Name: "foo",
	}

	bar := Person{
		Name: "bar",
	}

	addr := Address{
		Street: "fake street",
		Number: 123,
	}

	addr2 := Address{
		Street: "another fake street",
		Number: 456,
	}

	// create persons
	foo.Id, _ = c.Node(ctx).Put(foo)
	bar.Id, _ = c.Node(ctx).Put(bar)

	// create addressesByPerson
	addr.Id, _ = c.Node(ctx).Put(addr)
	addr2.Id, _ = c.Node(ctx).Put(addr2)

	// add edges between them
	c.Node(ctx, foo.Id).Edge(addr)
	c.Node(ctx, foo.Id).Edge(addr2)
	c.Node(ctx, bar.Id).Edge(addr)

	// add reverse edges
	c.Node(ctx, addr.Id).Edge(foo)
	c.Node(ctx, addr.Id).Edge(bar)
	c.Node(ctx, addr2.Id).Edge(foo)

	// get all saved persons
	var allPersons []Person
	c.Query(ctx).All(utils.Query{
		ID:    utils.ID{Id: "person", Type: "node:person"},
		Index: utils.ByType,
	}, &allPersons)

	// get persons by address
	var personsByAddress []Person
	c.Query(ctx).All(utils.Query{
		ID: utils.ID{Id: addr.Id, Type: "edge:person:"},
	}, &personsByAddress)

	// get addressesByPerson by person
	c.Query(ctx).All(utils.Query{
		ID: utils.ID{Id: foo.Id, Type: "edge:address:"},
	}, &foo.Addresses)

	// get one person
	var onePerson Person
	c.Query(ctx).One(utils.ID{Id: bar.Id, Type: "node:person"}, &onePerson)

	spew.Dump(allPersons)
	/*
		 [
			 {
			  Id: "person:73b1b881-be62-41d8-a3da-2ac5a6f25d6f",
			  Name: "bar",
			  Addresses: <nil>
			 },
			 {
			  Id: "person:7a29aaf9-2a28-4dcf-9f6b-2f50ffa9a5c7",
			  Name: "foo",
			  Addresses: <nil>
			 }
		]
	*/
	spew.Dump(personsByAddress)
	/*
		[
			{
			  Id: "address:fa547273-e265-4a47-a9ef-39c7538e4be8",
			  Name: "bar",
			  Addresses: <nil>
			 },
			 {
			  Id: "address:fa547273-e265-4a47-a9ef-39c7538e4be8",
			  Name: "foo",
			  Addresses: <nil>
			 }
		]
	*/
	spew.Dump(foo.Addresses)
	/*
		[
			 {
			  Id: "person:7a29aaf9-2a28-4dcf-9f6b-2f50ffa9a5c7",
			  Street: "another fake street",
			  Number: 456
			 },
			 {
			  Id: "person:7a29aaf9-2a28-4dcf-9f6b-2f50ffa9a5c7",
			  Street: "fake street",
			  Number: 123
			 }
		]
	*/
	spew.Dump(onePerson)
	/*
		{
		 Id: "person:73b1b881-be62-41d8-a3da-2ac5a6f25d6f",
		 Name: "bar",
		 Addresses: <nil>
		}
	*/
}

func connectAWS(region string) *session.Session {
	sess, err := session.NewSession(
		&aws.Config{
			Region: aws.String(region),
		},
	)
	if err != nil {
		panic(err)
	}
	return sess
}
