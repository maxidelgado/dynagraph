package node

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
		id  string
		t   dynamo.Table
	}
	tests := []struct {
		name string
		args args
		want Node
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.ctx, tt.args.id, tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_node_Delete(t *testing.T) {
	type fields struct {
		id    string
		ntype string
		t     dynamo.Table
		ctx   context.Context
	}
	type args struct {
		filter utils.Filter
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
			n := node{
				id:    tt.fields.id,
				ntype: tt.fields.ntype,
				t:     tt.fields.t,
				ctx:   tt.fields.ctx,
			}
			if err := n.Delete(tt.args.filter); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_node_Edge(t *testing.T) {
	type fields struct {
		id    string
		ntype string
		t     dynamo.Table
		ctx   context.Context
	}
	type args struct {
		value interface{}
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
			n := node{
				id:    tt.fields.id,
				ntype: tt.fields.ntype,
				t:     tt.fields.t,
				ctx:   tt.fields.ctx,
			}
			if err := n.Edge(tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("Edge() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_node_Prop(t *testing.T) {
	type fields struct {
		id    string
		ntype string
		t     dynamo.Table
		ctx   context.Context
	}
	type args struct {
		value interface{}
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
			n := node{
				id:    tt.fields.id,
				ntype: tt.fields.ntype,
				t:     tt.fields.t,
				ctx:   tt.fields.ctx,
			}
			if err := n.Prop(tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("Prop() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_node_Put(t *testing.T) {
	type fields struct {
		id    string
		ntype string
		t     dynamo.Table
		ctx   context.Context
	}
	type args struct {
		value interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := node{
				id:    tt.fields.id,
				ntype: tt.fields.ntype,
				t:     tt.fields.t,
				ctx:   tt.fields.ctx,
			}
			got, err := n.Put(tt.args.value)
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

func Test_node_Ref(t *testing.T) {
	type fields struct {
		id    string
		ntype string
		t     dynamo.Table
		ctx   context.Context
	}
	type args struct {
		refId string
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
			n := node{
				id:    tt.fields.id,
				ntype: tt.fields.ntype,
				t:     tt.fields.t,
				ctx:   tt.fields.ctx,
			}
			if err := n.Ref(tt.args.refId); (err != nil) != tt.wantErr {
				t.Errorf("Ref() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_node_Update(t *testing.T) {
	type fields struct {
		id    string
		ntype string
		t     dynamo.Table
		ctx   context.Context
	}
	type args struct {
		value interface{}
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
			n := node{
				id:    tt.fields.id,
				ntype: tt.fields.ntype,
				t:     tt.fields.t,
				ctx:   tt.fields.ctx,
			}
			if err := n.Update(tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
