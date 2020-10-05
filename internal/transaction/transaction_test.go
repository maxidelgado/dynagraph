package transaction

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/maxidelgado/dynagraph/internal/dynamoiface"
	"github.com/maxidelgado/dynagraph/utils"
	"reflect"
	"testing"
)

type db struct {
	dynamoiface.DB
	tx dynamoiface.WriteTx
}

func (m db) WriteTx() dynamoiface.WriteTx {
	return m.tx
}

type table struct {
	dynamoiface.Table
	delete delete
	put    put
	update update
}

func (m table) Delete(name string, value interface{}) dynamoiface.Delete {
	return m.delete
}

func (m table) Put(item interface{}) dynamoiface.Put {
	return m.put
}

func (m table) Update(hashKey string, value interface{}) dynamoiface.Update {
	return m.update
}

type writeTx struct {
	dynamoiface.WriteTx
	err error
}

func (m writeTx) Delete(d dynamoiface.Delete) dynamoiface.WriteTx {
	return m
}

func (m writeTx) Put(p dynamoiface.Put) dynamoiface.WriteTx {
	return m
}

func (m writeTx) Update(u dynamoiface.Update) dynamoiface.WriteTx {
	return m
}

func (m writeTx) RunWithContext(ctx aws.Context) error {
	return m.err
}

type delete struct {
	dynamoiface.Delete
	err error
}

func (m delete) Range(name string, value interface{}) dynamoiface.Delete {
	return m
}

func (m delete) RunWithContext(ctx aws.Context) error {
	return m.err
}

type put struct {
	dynamoiface.Put
	err error
}

func (m put) RunWithContext(ctx aws.Context) error {
	return m.err
}

type update struct {
	dynamoiface.Update
	err error
}

func (m update) Range(name string, value interface{}) dynamoiface.Update {
	return m
}

func (m update) Set(path string, value interface{}) dynamoiface.Update {
	return m
}

func (m update) RunWithContext(ctx aws.Context) error {
	return m.err
}

func TestNew(t *testing.T) {
	type args struct {
		ctx context.Context
		db  dynamoiface.DB
		t   dynamoiface.Table
	}
	tests := []struct {
		name string
		args args
		want Transaction
	}{
		{
			name: "success: create tx",
			args: args{
				ctx: context.Background(),
				db:  db{},
				t:   table{},
			},
			want: transaction{db: db{}, t: table{}, ctx: context.Background()},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.ctx, tt.args.db, tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_transaction_Delete(t1 *testing.T) {
	type fields struct {
		db  dynamoiface.DB
		t   dynamoiface.Table
		ctx context.Context
	}
	type args struct {
		keys utils.IDs
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Transaction
	}{
		{
			name: "success: delete node tx",
			fields: fields{
				db:  db{tx: writeTx{err: nil}},
				t:   table{delete: delete{err: nil}},
				ctx: context.Background(),
			},
			args: args{
				keys: utils.IDs{utils.ID{
					Id:   "id",
					Type: "type",
				}},
			},
			want: transaction{db: db{tx: writeTx{err: nil}}, t: table{delete: delete{err: nil}}, ctx: context.Background()},
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := transaction{
				db:  tt.fields.db,
				t:   tt.fields.t,
				ctx: tt.fields.ctx,
			}
			if got := t.Delete(tt.args.keys); !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_transaction_Put(t1 *testing.T) {
	ops := utils.Operations{}
	ops, _ = ops.AppendNode("id", struct {
		Id  string
		Msg string
	}{"id", "msg"})

	type fields struct {
		db  dynamoiface.DB
		t   dynamoiface.Table
		ctx context.Context
	}
	type args struct {
		inputs utils.Operations
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Transaction
	}{
		{
			name: "success: put node tx",
			fields: fields{
				db:  db{tx: writeTx{err: nil}},
				t:   table{put: put{err: nil}},
				ctx: context.Background(),
			},
			args: args{
				inputs: ops,
			},
			want: transaction{
				db:  db{tx: writeTx{err: nil}},
				t:   table{put: put{err: nil}},
				ctx: context.Background(),
			},
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := transaction{
				db:  tt.fields.db,
				t:   tt.fields.t,
				ctx: tt.fields.ctx,
			}
			if got := t.Put(tt.args.inputs); !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("Put() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_transaction_Run(t1 *testing.T) {
	type fields struct {
		db  dynamoiface.DB
		t   dynamoiface.Table
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "success: run tx",
			fields: fields{
				db:  db{tx: writeTx{err: nil}},
				ctx: context.Background(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := transaction{
				db:  tt.fields.db,
				t:   tt.fields.t,
				ctx: tt.fields.ctx,
			}
			if err := t.Run(); (err != nil) != tt.wantErr {
				t1.Errorf("Run() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_transaction_Update(t1 *testing.T) {
	ops := utils.Operations{}
	ops, _ = ops.AppendNode("id", struct {
		Id  string
		Msg string
	}{"id", "msg"})
	type fields struct {
		db  dynamoiface.DB
		t   dynamoiface.Table
		ctx context.Context
	}
	type args struct {
		inputs utils.Operations
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Transaction
	}{
		{
			name: "success: update node tx",
			fields: fields{
				db:  db{tx: writeTx{err: nil}},
				t:   table{update: update{err: nil}},
				ctx: context.Background(),
			},
			args: args{
				inputs: ops,
			},
			want: transaction{
				db:  db{tx: writeTx{err: nil}},
				t:   table{update: update{err: nil}},
				ctx: context.Background(),
			},
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := transaction{
				db:  tt.fields.db,
				t:   tt.fields.t,
				ctx: tt.fields.ctx,
			}
			if got := t.Update(tt.args.inputs); !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("Update() = %v, want %v", got, tt.want)
			}
		})
	}
}
