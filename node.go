package dynagraph

import (
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/guregu/dynamo"
)

type node struct {
	id    string
	ntype string
	table dynamo.Table
	err   error
}

func (n node) Put(value interface{}) (string, error) {
	n.ntype = getNodeType(value)
	if n.id == "" {
		n.id = NewId(value)
	}

	nodeMap, err := buildPut(n.id, n.ntype, value)
	if err != nil {
		return "", err
	}

	if err = n.table.Put(nodeMap).Run(); err != nil {
		return "", err
	}

	return n.id, nil
}

func (n node) Update(value interface{}) error {
	n.ntype = getNodeType(value)
	updateMap, err := buildUpdate(value)
	if err != nil {
		return err
	}

	update := n.table.
		Update(nodeId, n.id).
		Range(nodeType, nodeName+separator+n.ntype)

	for key, val := range updateMap {
		if key == nodeId {
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

	return update.Run()
}

func (n node) Edge(value interface{}) error {
	edgeMap, err := addEdge(n.id, value)
	if err != nil {
		return err
	}

	return n.table.Put(edgeMap).Run()
}

func (n node) Prop(value interface{}) error {
	propMap, err := addProp(n.id, value)
	if err != nil {
		return err
	}

	return n.table.Put(propMap).Run()
}

func (n node) Ref(refId string) error {
	refMap, err := addRef(n.id, refId)
	if err != nil {
		return err
	}

	return n.table.Put(refMap).Run()
}

func (n node) Delete(filter Filter) error {
	return n.table.
		Delete(nodeId, filter.HashKey()).
		Range(nodeType, filter.RangeKey()).
		Run()
}
