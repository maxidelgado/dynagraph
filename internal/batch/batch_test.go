package batch

import (
	"context"
	"github.com/guregu/dynamo"
	"github.com/maxidelgado/dynagraph/utils"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		ctx context.Context
		t   dynamo.Table
	}
	tests := []struct {
		name string
		args args
		want Batch
	}{
		// TODO: Add test cases.
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
		t   dynamo.Table
		ctx context.Context
	}
	type args struct {
		keys utils.KeysInput
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
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
		t   dynamo.Table
		ctx context.Context
	}
	type args struct {
		out  interface{}
		keys utils.KeysInput
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
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
		t   dynamo.Table
		ctx context.Context
	}
	type args struct {
		inputs utils.WriteItemsInput
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
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
