package client

import (
	"context"
	"reflect"
	"testing"

	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/guregu/dynamo"
	"github.com/maxidelgado/dynagraph/internal/batch"
	"github.com/maxidelgado/dynagraph/internal/node"
	"github.com/maxidelgado/dynagraph/internal/query"
	"github.com/maxidelgado/dynagraph/internal/transaction"
	"github.com/maxidelgado/dynagraph/utils"
)

func TestNew(t *testing.T) {
	type args struct {
		d     dynamodbiface.DynamoDBAPI
		table string
	}
	tests := []struct {
		name    string
		args    args
		want    Client
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.d, tt.args.table)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_client_Batch(t *testing.T) {
	type fields struct {
		db *dynamo.DB
		t  dynamo.Table
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   batch.Batch
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := client{
				db: tt.fields.db,
				t:  tt.fields.t,
			}
			if got := c.Batch(tt.args.ctx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Batch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_client_Node(t *testing.T) {
	type fields struct {
		db *dynamo.DB
		t  dynamo.Table
	}
	type args struct {
		ctx context.Context
		id  []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   node.Node
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := client{
				db: tt.fields.db,
				t:  tt.fields.t,
			}
			if got := c.Node(tt.args.ctx, tt.args.id...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Node() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_client_Query(t *testing.T) {
	type fields struct {
		db *dynamo.DB
		t  dynamo.Table
	}
	type args struct {
		ctx    context.Context
		filter utils.Filter
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   query.Query
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := client{
				db: tt.fields.db,
				t:  tt.fields.t,
			}
			if got := c.Query(tt.args.ctx, tt.args.filter); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Query() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_client_Transaction(t *testing.T) {
	type fields struct {
		db *dynamo.DB
		t  dynamo.Table
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   transaction.Transaction
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := client{
				db: tt.fields.db,
				t:  tt.fields.t,
			}
			if got := c.Transaction(tt.args.ctx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Transaction() = %v, want %v", got, tt.want)
			}
		})
	}
}
