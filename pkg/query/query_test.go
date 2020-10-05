package query

import (
	"context"
	"errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/guregu/dynamo"
	"github.com/maxidelgado/dynagraph/internal/dynamoiface"
	"github.com/maxidelgado/dynagraph/pkg/utils"
	"reflect"
	"testing"
)

type tMock struct {
	dynamoiface.Table
	get dynamoiface.Query
}

func (m tMock) Get(name string, value interface{}) dynamoiface.Query {
	return m.get
}

type getMock struct {
	dynamoiface.Query
	err error
}

func (m getMock) Range(name string, op dynamo.Operator, values ...interface{}) dynamoiface.Query {
	return m
}

func (m getMock) Index(name string) dynamoiface.Query {
	return m
}

func (m getMock) OneWithContext(ctx aws.Context, out interface{}) error {
	return m.err
}

func (m getMock) AllWithContext(ctx aws.Context, out interface{}) error {
	return m.err
}

func TestNew(t *testing.T) {
	type args struct {
		ctx context.Context
		t   dynamoiface.Table
	}
	tests := []struct {
		name string
		args args
		want Query
	}{
		{
			name: "success: query created",
			args: args{
				ctx: context.Background(),
				t:   tMock{get: getMock{}},
			},
			want: query{ctx: context.Background(), t: tMock{get: getMock{}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.ctx, tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_query_All(t *testing.T) {
	type fields struct {
		t   dynamoiface.Table
		ctx context.Context
	}
	type args struct {
		query utils.Query
		out   interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "success: filter nodes by type",
			fields: fields{
				t:   tMock{get: getMock{err: nil}},
				ctx: context.Background(),
			},
			args: args{
				query: utils.Query{ID: utils.ID{Type: "prefix"}, Operator: utils.BeginsWith, Index: utils.ByType},
				out:   []string{},
			},
			wantErr: false,
		},
		{
			name: "success: filter nodes by id",
			fields: fields{
				t:   tMock{get: getMock{err: nil}},
				ctx: context.Background(),
			},
			args: args{
				query: utils.Query{ID: utils.ID{Id: "id"}, Operator: utils.BeginsWith, Index: utils.Default},
				out:   []string{},
			},
			wantErr: false,
		},
		{
			name: "success: with no filter",
			fields: fields{
				t:   tMock{get: getMock{err: nil}},
				ctx: context.Background(),
			},
			args: args{
				query: utils.Query{ID: utils.ID{Id: "id"}, Operator: utils.Noop, Index: utils.Default},
				out:   []string{},
			},
			wantErr: false,
		},
		{
			name: "fail: get error",
			fields: fields{
				t:   tMock{get: getMock{err: errors.New("error")}},
				ctx: context.Background(),
			},
			args: args{
				query: utils.Query{ID: utils.ID{Id: "id"}, Operator: utils.Noop, Index: utils.Default},
				out:   []string{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := query{
				t:   tt.fields.t,
				ctx: tt.fields.ctx,
			}
			if err := f.All(tt.args.query, tt.args.out); (err != nil) != tt.wantErr {
				t.Errorf("All() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_query_One(t *testing.T) {
	type fields struct {
		t   dynamoiface.Table
		ctx context.Context
	}
	type args struct {
		id  utils.ID
		out interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "success: get single node",
			fields: fields{
				t:   tMock{get: getMock{}},
				ctx: context.Background(),
			},
			args: args{
				id: utils.ID{},
				out: &struct {
					Id string
				}{},
			},
			wantErr: false,
		},
		{
			name: "fail: get error",
			fields: fields{
				t:   tMock{get: getMock{err: errors.New("error")}},
				ctx: context.Background(),
			},
			args: args{
				id: utils.ID{},
				out: &struct {
					Id string
				}{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := query{
				t:   tt.fields.t,
				ctx: tt.fields.ctx,
			}
			if err := f.One(tt.args.id, tt.args.out); (err != nil) != tt.wantErr {
				t.Errorf("One() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
