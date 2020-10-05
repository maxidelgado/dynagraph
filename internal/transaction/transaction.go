package transaction

import (
	"context"
	"github.com/maxidelgado/dynagraph/internal/common"
	"github.com/maxidelgado/dynagraph/internal/dynamoiface"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/maxidelgado/dynagraph/utils"
)

type Transaction interface {
	Put(inputs utils.Operations) Transaction
	Update(inputs utils.Operations) Transaction
	Delete(keys utils.IDs) Transaction
	Run() error
}

func New(ctx context.Context, db dynamoiface.DB, t dynamoiface.Table) Transaction {
	return transaction{
		db:  db,
		t:   t,
		ctx: ctx,
	}
}

type transaction struct {
	db  dynamoiface.DB
	t   dynamoiface.Table
	ctx context.Context
}

func (t transaction) Put(inputs utils.Operations) Transaction {
	for _, input := range inputs {
		t.db.WriteTx().Put(t.t.Put(input))
	}

	return t
}

func (t transaction) Delete(keys utils.IDs) Transaction {
	for _, key := range keys {
		t.db.WriteTx().Delete(t.t.Delete(common.NodeId, key.HashKey()).Range(common.NodeType, key.RangeKey()))
	}

	return t
}

func (t transaction) Update(inputs utils.Operations) Transaction {
	for _, input := range inputs {
		in := input.(map[string]*dynamodb.AttributeValue)
		update := t.t.
			Update(common.NodeId, in[common.NodeId].S).
			Range(common.NodeType, in[common.NodeType].S)

		for key, val := range in {
			if key == common.NodeId || key == common.NodeType {
				continue
			}

			var out interface{}
			err := dynamodbattribute.Unmarshal(val, &out)
			if err != nil {
				return nil
			}

			if out == nil {
				continue
			}

			update.Set(key, out)
		}

		t.db.WriteTx().Update(update)
	}

	return t
}

func (t transaction) Run() error {
	return t.db.WriteTx().RunWithContext(t.ctx)
}
