package batch

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/guregu/dynamo"
	"github.com/maxidelgado/dynagraph/internal/dynamoiface"
	"github.com/maxidelgado/dynagraph/pkg/utils"
)

type tMock struct {
	dynamoiface.Table
	batch dynamoiface.Batch
}

func (m tMock) Batch(hashAndRangeKeyName ...string) dynamoiface.Batch {
	return m.batch
}

type bMock struct {
	dynamoiface.Batch
	batchWrite dynamoiface.BatchWrite
	batchGet   dynamoiface.BatchGet
}

func (m bMock) Write() dynamoiface.BatchWrite {
	return m.batchWrite
}

func (m bMock) Get(keys ...dynamo.Keyed) dynamoiface.BatchGet {
	return m.batchGet
}

type bwMock struct {
	dynamoiface.BatchWrite
	err   error
	wrote int
}

func (m bwMock) Put(items ...interface{}) dynamoiface.BatchWrite {
	return m
}

func (m bwMock) Delete(keys ...dynamo.Keyed) dynamoiface.BatchWrite {
	return m
}

func (m bwMock) RunWithContext(ctx aws.Context) (wrote int, err error) {
	return m.wrote, m.err
}

type bgMock struct {
	dynamoiface.BatchGet
	err error
}

func (m bgMock) AllWithContext(ctx aws.Context, out interface{}) error {
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
		want Batch
	}{
		{
			name: "success: batch client created",
			args: args{},
			want: batch{},
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

func Test_batch_Delete(t *testing.T) {

	type fields struct {
		t   dynamoiface.Table
		ctx context.Context
	}
	type args struct {
		keys utils.IDs
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "success: one node deleted",
			fields: fields{
				t:   tMock{batch: bMock{batchWrite: bwMock{err: nil, wrote: 1}}},
				ctx: context.Background(),
			},
			args: args{
				keys: utils.IDs{},
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "fail: delete error",
			fields: fields{
				t:   tMock{batch: bMock{batchWrite: bwMock{err: errors.New("error"), wrote: 0}}},
				ctx: context.Background(),
			},
			args: args{
				keys: utils.IDs{},
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := batch{
				t:   tt.fields.t,
				ctx: tt.fields.ctx,
			}
			got, err := b.Delete(tt.args.keys)
			if (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Delete() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_batch_Get(t *testing.T) {
	type fields struct {
		t   dynamoiface.Table
		ctx context.Context
	}
	type args struct {
		out  interface{}
		keys utils.IDs
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "success: get items",
			fields: fields{
				t:   tMock{batch: bMock{batchGet: bgMock{err: nil}}},
				ctx: context.Background(),
			},
			args: args{
				out:  &struct{ Id string }{},
				keys: utils.IDs{},
			},
			wantErr: false,
		},
		{
			name: "fail: get items error",
			fields: fields{
				t:   tMock{batch: bMock{batchGet: bgMock{err: errors.New("error")}}},
				ctx: context.Background(),
			},
			args: args{
				out:  &struct{ Id string }{},
				keys: utils.IDs{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := batch{
				t:   tt.fields.t,
				ctx: tt.fields.ctx,
			}
			if err := b.Get(tt.args.out, tt.args.keys); (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_batch_Put(t *testing.T) {
	type fields struct {
		t   dynamoiface.Table
		ctx context.Context
	}
	type args struct {
		inputs utils.Operations
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "success: items created",
			fields: fields{
				t:   tMock{batch: bMock{batchWrite: bwMock{err: nil, wrote: 1}}},
				ctx: context.Background(),
			},
			args: args{
				inputs: utils.Operations{},
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "fail: create items error",
			fields: fields{
				t:   tMock{batch: bMock{batchWrite: bwMock{err: errors.New("error"), wrote: 0}}},
				ctx: context.Background(),
			},
			args: args{
				inputs: utils.Operations{},
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := batch{
				t:   tt.fields.t,
				ctx: tt.fields.ctx,
			}
			got, err := b.Put(tt.args.inputs)
			if (err != nil) != tt.wantErr {
				t.Errorf("Put() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Put() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFechas(t *testing.T) {
	t.Run("ok", func(t *testing.T) {

	})
}
