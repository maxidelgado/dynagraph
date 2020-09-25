package dynagraph

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/guregu/dynamo"
)

type transaction struct {
	wtx   *dynamo.WriteTx
	db    *dynamo.DB
	table dynamo.Table
}

func (t *transaction) Put(inputs WriteItemsInput) Transaction {
	for _, input := range inputs {
		t.wtx.Put(t.table.Put(input))
	}

	return t
}

func (t *transaction) Delete(keys KeysInput) Transaction {
	for _, key := range keys {
		t.wtx.Delete(t.table.Delete(nodeId, key.HashKey()).Range(nodeType, key.RangeKey()))
	}

	return t
}

func (t *transaction) Update(inputs WriteItemsInput) Transaction {
	for _, input := range inputs {
		in := input.(map[string]*dynamodb.AttributeValue)
		update := t.table.
			Update(nodeId, in[nodeId].S).
			Range(nodeType, in[nodeType].S)

		for key, val := range in {
			if key == nodeId || key == nodeType {
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

		t.wtx.Update(update)
	}

	return t
}

func (t *transaction) Run() error {
	return t.wtx.Run()
}
