package node

import (
	"context"
	"github.com/maxidelgado/dynagraph/internal/common"

	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/maxidelgado/dynagraph/internal/dynamoiface"
	"github.com/maxidelgado/dynagraph/utils"
)

var (
	newId = common.Id
)

func New(ctx context.Context, id string, t dynamoiface.Table) Node {
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
	Delete(filter utils.ID) error
}

type node struct {
	id    string
	ntype string
	t     dynamoiface.Table
	ctx   context.Context
}

func (n node) Put(value interface{}) (string, error) {
	n.ntype = common.Type(value)
	if n.id == "" {
		n.id = newId(value)
	}

	nodeMap, err := common.Put(n.id, n.ntype, value)
	if err != nil {
		return "", err
	}

	if err = n.t.Put(nodeMap).RunWithContext(n.ctx); err != nil {
		return "", err
	}

	return n.id, nil
}

func (n node) Update(value interface{}) error {
	n.ntype = common.Type(value)
	updateMap, err := common.Update(value)
	if err != nil {
		return err
	}

	update := n.t.
		Update(common.NodeId, n.id).
		Range(common.NodeType, common.NodeName+common.Separator+n.ntype)

	for key, val := range updateMap {
		if key == common.NodeId {
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
	edgeMap, err := common.Edge(n.id, value)
	if err != nil {
		return err
	}

	return n.t.Put(edgeMap).RunWithContext(n.ctx)
}

func (n node) Prop(value interface{}) error {
	propMap, err := common.Prop(n.id, value)
	if err != nil {
		return err
	}

	return n.t.Put(propMap).RunWithContext(n.ctx)
}

func (n node) Ref(refId string) error {
	refMap, err := common.Ref(n.id, refId)
	if err != nil {
		return err
	}

	return n.t.Put(refMap).RunWithContext(n.ctx)
}

func (n node) Delete(filter utils.ID) error {
	return n.t.
		Delete(common.NodeId, filter.HashKey()).
		Range(common.NodeType, filter.RangeKey()).
		RunWithContext(n.ctx)
}
