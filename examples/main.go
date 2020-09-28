package main

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/maxidelgado/dynagraph"
)

func main() {
	sess := session.Must(session.NewSession(&aws.Config{Region: aws.String("us-east-1")}))
	g := dynagraph.New(sess, dynagraph.Config{TableName: "SOME TABLE"})

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
	u.Id, err = g.Node().Put(u)
	if err != nil {
		log.Fatal(err)
	}

	// create product
	p.Id, err = g.Node().Put(p)
	if err != nil {
		log.Fatal(err)
	}

	// create order tx
	o := ExampleOrder{}
	o.Id = dynagraph.NewId(o)

	inputs := dynagraph.WriteItemsInput{}
	inputs, _ = inputs.AppendNode(o.Id, o)
	inputs, _ = inputs.AppendEdge(o.Id, u)
	inputs, _ = inputs.AppendEdge(o.Id, p)
	err = g.Transaction().Put(inputs).Run()
	if err != nil {
		log.Fatal(err)
	}

	// update user
	err = g.Node(u.Id).Update(ExampleUser{FirstName: "updated_first_name"})
	if err != nil {
		log.Fatal(err)
	}

	// get order edges
	var edges []struct {
		Id   string
		Type string
	}
	err = g.Query(dynagraph.Filter{Id: o.Id, Type: "edge"}).All(&edges)
	if err != nil {
		log.Fatal(err)
	}

	// delete example by using transactions
	keys := dynagraph.KeysInput{}
	for _, edge := range edges {
		keys = append(keys, dynagraph.Filter{Id: edge.Id, Type: edge.Type})
	}
	keys = keys.AppendKeys(dynagraph.Filter{Id: p.Id, Type: "node:exampleproduct"})
	keys = keys.AppendKeys(dynagraph.Filter{Id: o.Id, Type: "prop:exampleorder"})
	err = g.Transaction().Delete(keys).Run()
	if err != nil {
		log.Fatal(err)
	}
	// delete individual node
	err = g.Node().Delete(dynagraph.Filter{Id: u.Id, Type: "node:exampleuser"})
	if err != nil {
		log.Fatal(err)
	}

	/*
		u.Id, err = g.Node().Put(u)
		u.Id, err = g.Node("1234").Put(u)
		u.Id, err = g.Node("pisame_esta").Put(u)

		err = g.Node("1234").Update(u)
		err = g.Node().Update(u) // error
		err = g.Node("").Update(u) // error
		err = g.Node("no_existe").Update(u) // error

		err = g.Node(u.Id).Edge(p) // OK
		err = g.Node("no_existe").Edge(p) // error

		err = g.Node(u.Id).Prop(p) // OK
		err = g.Node("no_existe").Prop(p) // error

		err = g.Node(u.Id).Ref(p.Id) // OK
		err = g.Node("no_existe").Ref(p.Id) // error

		err = g.Node().Delete(dynagraph.Filter{"1234", "untipo"}) // OK
		err = g.Node().Delete(dynagraph.Filter{"1234", ""}) // error
		err = g.Node().Delete(dynagraph.Filter{}) // error

		err = g.Query(dynagraph.Filter{"1234", "untipo"}).One(&result)
	*/
}
