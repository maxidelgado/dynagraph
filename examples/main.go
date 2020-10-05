package main

import (
	"context"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/maxidelgado/dynagraph"
	"github.com/maxidelgado/dynagraph/pkg/utils"
)

func init() {
	os.Setenv("AWS_REGION", "us-east-1")
	os.Setenv("AWS_PROFILE", "astrocode")
}

type A struct {
	Id    string
	Value string
	B     B   `dynamo:"-"`
	Cs    []C `dynamo:"-"`
}

type B struct {
	Id    string
	Value string
}

type C struct {
	Id    string
	Value string
}

func main() {
	var (
		sess = connectAWS("us-east-1")
		db   = dynamodb.New(sess)
		c, _ = dynagraph.New(db, "astrocode")
		ctx  = context.Background()
	)

	// setup nodes
	nodeA := A{
		Id:    "a:id",
		Value: "some field value",
	}
	nodeB := B{
		Id:    "b:id",
		Value: "some field value",
	}
	nodeC := C{
		Id:    "c:id",
		Value: "some field value",
	}
	nodeA.B = nodeB
	nodeA.Cs = append(nodeA.Cs, nodeC)

	// create source node A
	c.Node(ctx, nodeA.Id).Put(nodeA)

	// create child node B (one to one)
	c.Node(ctx, nodeA.Id).Prop(nodeB)

	// create nodeA <-> nodeC edge (one/many to many)
	c.Node(ctx, nodeA.Id).Edge(nodeC)

	// update node A
	c.Node(ctx, nodeA.Id).Update(A{Value: "updated value"})

	// get nodeA
	var resultNodeA A
	c.Query(ctx).One(utils.ID{Id: nodeA.Id, Type: "node:a"}, &resultNodeA)

	// get nodeA property B
	c.Query(ctx).One(utils.ID{Id: nodeA.Id, Type: "prop:b"}, &resultNodeA.B)

	// get nodeA edges
	c.Query(ctx).All(utils.Query{ID: utils.ID{Id: nodeA.Id, Type: "edge"}}, &resultNodeA.Cs)

	// delete nodes
	nodeId := utils.ID{Id: nodeA.Id, Type: "node:a"}
	propId := utils.ID{Id: nodeA.Id, Type: "prop:b"}
	edgeId := utils.ID{Id: nodeA.Id, Type: "edge:c:id"}
	c.Node(ctx).Delete(nodeId)
	c.Node(ctx).Delete(propId)
	c.Node(ctx).Delete(edgeId)
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
