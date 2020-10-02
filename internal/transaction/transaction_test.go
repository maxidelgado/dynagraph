package transaction

import (
	"context"
	"reflect"
	"testing"

	"github.com/guregu/dynamo"
	"github.com/maxidelgado/dynagraph/utils"
)

func TestNew(t *testing.T) {
	type args struct {
		ctx context.Context
		db  *dynamo.DB
		t   dynamo.Table
	}
	tests := []struct {
		name string
		args args
		want Transaction
	}{
		// TODO: Add test cases.
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
		db  *dynamo.DB
		t   dynamo.Table
		ctx context.Context
	}
	type args struct {
		keys utils.KeysInput
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Transaction
	}{
		// TODO: Add test cases.
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
	type fields struct {
		db  *dynamo.DB
		t   dynamo.Table
		ctx context.Context
	}
	type args struct {
		inputs utils.WriteItemsInput
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Transaction
	}{
		// TODO: Add test cases.
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
		db  *dynamo.DB
		t   dynamo.Table
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
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
	type fields struct {
		db  *dynamo.DB
		t   dynamo.Table
		ctx context.Context
	}
	type args struct {
		inputs utils.WriteItemsInput
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Transaction
	}{
		// TODO: Add test cases.
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
