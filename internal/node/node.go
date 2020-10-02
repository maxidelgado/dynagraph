package node

import (
	"context"

	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/guregu/dynamo"
	"github.com/maxidelgado/dynagraph/utils"
)

func New(ctx context.Context, id string, t dynamo.Table) Node {
	return node{
		id:  id,
		ctx: ctx,
		t:   t,
	}
}

type Node interface {
	Put(value interface{}) (string, error)
	Update(value interface{}) error
	Edge(value interface{}) error
	Prop(value interface{}) error
	Ref(id string) error
	Delete(filter utils.Filter) error
}

type node struct {
	id    string
	ntype string
	t     dynamo.Table
	ctx   context.Context
}

func (n node) Put(value interface{}) (string, error) {
	n.ntype = utils.GetNodeType(value)
	if n.id == "" {
		n.id = utils.NewId(value)
	}

	nodeMap, err := utils.BuildPut(n.id, n.ntype, value)
	if err != nil {
		return "", err
	}

	if err = n.t.Put(nodeMap).RunWithContext(n.ctx); err != nil {
		return "", err
	}

	return n.id, nil
}

func (n node) Update(value interface{}) error {
	n.ntype = utils.GetNodeType(value)
	updateMap, err := utils.BuildUpdate(value)
	if err != nil {
		return err
	}

	update := n.t.
		Update(utils.NodeId, n.id).
		Range(utils.NodeType, utils.NodeName+utils.Separator+n.ntype)

	for key, val := range updateMap {
		if key == utils.NodeId {
			continue
		}

		var out interface{}
		err := dynamodbattribute.Unmarshal(val, &out)
		if err != nil {
			return err
		}

		if out == nil {
			continue
		}

		update.Set(key, out)
	}

	return update.RunWithContext(n.ctx)
}

func (n node) Edge(value interface{}) error {
	edgeMap, err := utils.AddEdge(n.id, value)
	if err != nil {
		return err
	}

	return n.t.Put(edgeMap).RunWithContext(n.ctx)
}

func (n node) Prop(value interface{}) error {
	propMap, err := utils.AddProp(n.id, value)
	if err != nil {
		return err
	}

	return n.t.Put(propMap).RunWithContext(n.ctx)
}

func (n node) Ref(refId string) error {
	refMap, err := utils.AddRef(n.id, refId)
	if err != nil {
		return err
	}

	return n.t.Put(refMap).RunWithContext(n.ctx)
}

func (n node) Delete(filter utils.Filter) error {
	return n.t.
		Delete(utils.NodeId, filter.HashKey()).
		Range(utils.NodeType, filter.RangeKey()).
		RunWithContext(n.ctx)
}
