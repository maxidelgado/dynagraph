package main

import (
	"context"
	"github.com/maxidelgado/dynagraph/internal/common"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/maxidelgado/dynagraph/client"
	"github.com/maxidelgado/dynagraph/utils"
)

func main() {
	sess := connectAWS("us-east-1")
	db := dynamodb.New(sess)
	c, _ := client.New(db, "yourTableName")
	ctx := context.Background()

	var err error

	// create user
	u := ExampleUser{
		FirstName: "first_name",
		LastName:  "last_name",
		Address:   "address",
	}
	p := ExampleProduct{
		Name:        "name",
		Description: "description",
	}
	u.Id, err = c.Node(ctx).Put(u)
	if err != nil {
		log.Fatal(err)
	}

	// create product
	p.Id, err = c.Node(ctx).Put(p)
	if err != nil {
		log.Fatal(err)
	}

	// create order tx
	o := ExampleOrder{}
	o.Id = common.Id(o)

	inputs := utils.Operations{}
	inputs, _ = inputs.AppendNode(o.Id, o)
	inputs, _ = inputs.AppendEdge(o.Id, u)
	inputs, _ = inputs.AppendEdge(o.Id, p)
	err = c.Transaction(ctx).Put(inputs).Run()
	if err != nil {
		log.Fatal(err)
	}

	// update user
	err = c.Node(ctx, u.Id).Update(ExampleUser{FirstName: "updated_first_name"})
	if err != nil {
		log.Fatal(err)
	}

	// get order edges
	var edges []struct {
		Id   string
		Type string
	}
	err = c.Query(ctx).All(utils.Query{ID: utils.ID{Id: o.Id, Type: "edge"}}, &edges)
	if err != nil {
		log.Fatal(err)
	}

	// delete example by using transactions
	keys := utils.IDs{}
	for _, edge := range edges {
		keys = append(keys, utils.ID{Id: edge.Id, Type: edge.Type})
	}
	keys = keys.AppendKeys(utils.ID{Id: p.Id, Type: "node:exampleproduct"})
	keys = keys.AppendKeys(utils.ID{Id: o.Id, Type: "prop:exampleorder"})
	err = c.Transaction(ctx).Delete(keys).Run()
	if err != nil {
		log.Fatal(err)
	}
	// delete individual node
	err = c.Node(ctx).Delete(utils.ID{Id: u.Id, Type: "node:exampleuser"})
	if err != nil {
		log.Fatal(err)
	}

	/*
		u.Id, err = c.Node().Put(u)
		u.Id, err = c.Node("1234").Put(u)
		u.Id, err = c.Node("pisame_esta").Put(u)

		err = c.Node("1234").Update(u)
		err = c.Node().Update(u) // error
		err = c.Node("").Update(u) // error
		err = c.Node("no_existe").Update(u) // error

		err = c.Node(u.Id).Edge(p) // OK
		err = c.Node("no_existe").Edge(p) // error

		err = c.Node(u.Id).Prop(p) // OK
		err = c.Node("no_existe").Prop(p) // error

		err = c.Node(u.Id).Ref(p.Id) // OK
		err = c.Node("no_existe").Ref(p.Id) // error

		err = c.Node().Delete(dynagraph.Query{"1234", "untipo"}) // OK
		err = c.Node().Delete(dynagraph.Query{"1234", ""}) // error
		err = c.Node().Delete(dynagraph.Query{}) // error

		err = c.Query(dynagraph.Query{"1234", "untipo"}).One(&result)
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
