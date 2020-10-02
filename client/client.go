package client

import (
	"context"
	"errors"

	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/guregu/dynamo"
	"github.com/maxidelgado/dynagraph/internal/batch"
	"github.com/maxidelgado/dynagraph/internal/node"
	"github.com/maxidelgado/dynagraph/internal/query"
	"github.com/maxidelgado/dynagraph/internal/transaction"
	"github.com/maxidelgado/dynagraph/utils"
)

var (
	newNode  = node.New
	newQuery = query.New
	newBatch = batch.New
	newTx    = transaction.New
)

type Client interface {
	Node(context.Context, ...string) node.Node
	Query(context.Context, utils.Filter) query.Query
	Batch(context.Context) batch.Batch
	Transaction(context.Context) transaction.Transaction
}

func New(d dynamodbiface.DynamoDBAPI, table string) (Client, error) {
	if table == "" {
		return nil, errors.New("t name is required")
	}

	db := dynamo.NewFromIface(d)

	return client{
		t: db.Table(table),
	}, nil
}

type client struct {
	db *dynamo.DB
	t  dynamo.Table
}

// If the Node id is not set when calling the Node() method, then a random id will be configured on it.
// You can check this value by accessing the Node Id() method.
func (c client) Node(ctx context.Context, id ...string) node.Node {
	if len(id) == 0 {
		return newNode(ctx, "", c.t)
	}
	return newNode(ctx, id[0], c.t)
}

func (c client) Query(ctx context.Context, filter utils.Filter) query.Query {
	return newQuery(ctx, c.t, filter)
}

func (c client) Batch(ctx context.Context) batch.Batch {
	return newBatch(ctx, c.t)
}

func (c client) Transaction(ctx context.Context) transaction.Transaction {
	return newTx(ctx, c.db, c.t)
}
