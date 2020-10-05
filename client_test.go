package dynagraph

import (
	"context"
	"reflect"
	"testing"

	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/maxidelgado/dynagraph/internal/dynamoiface"
	"github.com/maxidelgado/dynagraph/pkg/batch"
	"github.com/maxidelgado/dynagraph/pkg/node"
	"github.com/maxidelgado/dynagraph/pkg/query"
	"github.com/maxidelgado/dynagraph/pkg/transaction"
	"github.com/maxidelgado/dynagraph/pkg/utils"
)

type dbMock struct {
	dynamodbiface.DynamoDBAPI
}

type batchMock struct {
	batch.Batch
}

type nodeMock struct {
	node.Node
}

type queryMock struct {
	query.Query
}

type txMock struct {
	transaction.Transaction
}

func TestNew(t *testing.T) {
	type args struct {
		db    dynamodbiface.DynamoDBAPI
		table string
	}
	tests := []struct {
		name    string
		args    args
		want    Client
		wantErr bool
	}{
		{
			name: "success: client created",
			args: args{
				db:    dbMock{},
				table: "mock_table",
			},
			wantErr: false,
		},
		{
			name: "fail: table name is empty",
			args: args{
				db:    dbMock{},
				table: "",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := New(tt.args.db, tt.args.table)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_client_Batch(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		newFunc func(ctx context.Context, t dynamoiface.Table) batch.Batch
		want    batch.Batch
	}{
		{
			name: "success: batch client created",
			newFunc: func(ctx context.Context, t dynamoiface.Table) batch.Batch {
				return batchMock{}
			},
			args: args{
				ctx: context.Background(),
			},
			want: batchMock{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := client{}
			newBatch = tt.newFunc
			if got := c.Batch(tt.args.ctx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Batch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_client_Node(t *testing.T) {
	type args struct {
		ctx context.Context
		id  []string
	}
	tests := []struct {
		name    string
		args    args
		newFunc func(ctx context.Context, id string, t dynamoiface.Table) node.Node
		want    node.Node
	}{
		{
			name: "success: node client created without id",
			args: args{ctx: context.Background()},
			newFunc: func(ctx context.Context, id string, t dynamoiface.Table) node.Node {
				return nodeMock{}
			},
			want: nodeMock{},
		},
		{
			name: "success: node client created with id",
			args: args{ctx: context.Background(), id: []string{"id"}},
			newFunc: func(ctx context.Context, id string, t dynamoiface.Table) node.Node {
				return nodeMock{}
			},
			want: nodeMock{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := client{}
			newNode = tt.newFunc
			if got := c.Node(tt.args.ctx, tt.args.id...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Node() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_client_Query(t *testing.T) {
	type args struct {
		ctx    context.Context
		filter utils.Query
	}
	tests := []struct {
		name    string
		args    args
		newFunc func(ctx context.Context, t dynamoiface.Table) query.Query
		want    query.Query
	}{
		{
			name: "success: query client created",
			args: args{ctx: context.Background()},
			newFunc: func(ctx context.Context, t dynamoiface.Table) query.Query {
				return queryMock{}
			},
			want: queryMock{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := client{}
			newQuery = tt.newFunc
			if got := c.Query(tt.args.ctx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Query() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_client_Transaction(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		newFunc func(ctx context.Context, db dynamoiface.DB, t dynamoiface.Table) transaction.Transaction
		want    transaction.Transaction
	}{
		{
			name: "success: tx client created",
			args: args{},
			newFunc: func(ctx context.Context, db dynamoiface.DB, t dynamoiface.Table) transaction.Transaction {
				return txMock{}
			},
			want: txMock{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := client{}
			newTx = tt.newFunc
			if got := c.Transaction(tt.args.ctx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Transaction() = %v, want %v", got, tt.want)
			}
		})
	}
}
