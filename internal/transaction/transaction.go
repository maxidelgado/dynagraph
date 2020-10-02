package transaction

import (
	"context"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/guregu/dynamo"
	"github.com/maxidelgado/dynagraph/utils"
)

type Transaction interface {
	Put(inputs utils.WriteItemsInput) Transaction
	Update(inputs utils.WriteItemsInput) Transaction
	Delete(keys utils.KeysInput) Transaction
	Run() error
}

func New(ctx context.Context, db *dynamo.DB, t dynamo.Table) Transaction {
	return transaction{
		db:  db,
		t:   t,
		ctx: ctx,
	}
}

type transaction struct {
	db  *dynamo.DB
	t   dynamo.Table
	ctx context.Context
}

func (t transaction) Put(inputs utils.WriteItemsInput) Transaction {
	for _, input := range inputs {
		t.db.WriteTx().Put(t.t.Put(input))
	}

	return t
}

func (t transaction) Delete(keys utils.KeysInput) Transaction {
	for _, key := range keys {
		t.db.WriteTx().Delete(t.t.Delete(utils.NodeId, key.HashKey()).Range(utils.NodeType, key.RangeKey()))
	}

	return t
}

func (t transaction) Update(inputs utils.WriteItemsInput) Transaction {
	for _, input := range inputs {
		in := input.(map[string]*dynamodb.AttributeValue)
		update := t.t.
			Update(utils.NodeId, in[utils.NodeId].S).
			Range(utils.NodeType, in[utils.NodeType].S)

		for key, val := range in {
			if key == utils.NodeId || key == utils.NodeType {
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
