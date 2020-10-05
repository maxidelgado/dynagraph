package node

import (
	"context"
	"errors"
	"github.com/aws/aws-sdk-go/aws"
	"reflect"
	"testing"

	"github.com/maxidelgado/dynagraph/internal/dynamoiface"
	"github.com/maxidelgado/dynagraph/utils"
)

type tMock struct {
	dynamoiface.Table
	delete dMock
	put    pMock
	update uMock
}

func (m tMock) Delete(name string, value interface{}) dynamoiface.Delete {
	return m.delete
}

func (m tMock) Put(item interface{}) dynamoiface.Put {
	return m.put
}

func (m tMock) Update(hashKey string, value interface{}) dynamoiface.Update {
	return m.update
}

type dMock struct {
	dynamoiface.Delete
	err error
}

func (m dMock) Range(name string, value interface{}) dynamoiface.Delete {
	return m
}

func (m dMock) RunWithContext(ctx aws.Context) error {
	return m.err
}

type pMock struct {
	dynamoiface.Put
	err error
}

func (m pMock) RunWithContext(ctx aws.Context) error {
	return m.err
}

type uMock struct {
	dynamoiface.Update
	err error
}

func (m uMock) Range(name string, value interface{}) dynamoiface.Update {
	return m
}

func (m uMock) Set(path string, value interface{}) dynamoiface.Update {
	return m
}

func (m uMock) RunWithContext(ctx aws.Context) error {
	return m.err
}

func TestNew(t *testing.T) {
	type args struct {
		ctx context.Context
		id  string
		t   dynamoiface.Table
	}
	tests := []struct {
		name string
		args args
		want Node
	}{
		{
			name: "success: node client created",
			args: args{
				ctx: context.Background(),
				id:  "id",
				t:   tMock{},
			},
			want: node{id: "id", t: tMock{}, ctx: context.Background()},
		},
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
		t     dynamoiface.Table
		ctx   context.Context
	}
	type args struct {
		filter utils.ID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "success: node deleted",
			fields: fields{
				t:   tMock{delete: dMock{err: nil}},
				ctx: context.Background(),
			},
			args:    args{utils.ID{}},
			wantErr: false,
		},
		{
			name: "fail: delete node error",
			fields: fields{
				t:   tMock{delete: dMock{err: errors.New("error")}},
				ctx: context.Background(),
			},
			args:    args{utils.ID{}},
			wantErr: true,
		},
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
		t     dynamoiface.Table
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
		{
			name: "success: added edge",
			fields: fields{
				id:    "from",
				ntype: "type",
				t:     tMock{put: pMock{err: nil}},
				ctx:   nil,
			},
			args: args{
				value: struct {
					Id  string
					Msg string
				}{"id", "msg"},
			},
			wantErr: false,
		},
		{
			name: "fail: empty id",
			fields: fields{
				id:    "from",
				ntype: "type",
				t:     tMock{put: pMock{err: nil}},
				ctx:   nil,
			},
			args: args{
				value: struct {
					Id  string
					Msg string
				}{"", "msg"},
			},
			wantErr: true,
		},
		{
			name: "fail: put error",
			fields: fields{
				id:    "from",
				ntype: "type",
				t:     tMock{put: pMock{err: errors.New("error")}},
				ctx:   nil,
			},
			args: args{
				value: struct {
					Id  string
					Msg string
				}{"id", "msg"},
			},
			wantErr: true,
		},
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
		t     dynamoiface.Table
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
		{
			name: "success: added property",
			fields: fields{
				id:    "id",
				ntype: "type",
				t:     tMock{put: pMock{err: nil}},
				ctx:   context.Background(),
			},
			args: args{
				value: struct {
					Id  string
					Msg string
				}{"id", "msg"},
			},
			wantErr: false,
		},
		{
			name: "fail: empty id",
			fields: fields{
				t:   tMock{put: pMock{err: nil}},
				ctx: context.Background(),
			},
			args: args{
				value: struct {
					Id  string
					Msg string
				}{"", "msg"},
			},
			wantErr: true,
		},
		{
			name: "fail: put error",
			fields: fields{
				t:   tMock{put: pMock{err: errors.New("error")}},
				ctx: context.Background(),
			},
			args: args{
				value: struct {
					Id  string
					Msg string
				}{"", "msg"},
			},
			wantErr: true,
		},
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
		t     dynamoiface.Table
		ctx   context.Context
	}
	type args struct {
		value interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		newId   func(value interface{}) string
		want    string
		wantErr bool
	}{
		{
			name: "success: create node",
			fields: fields{
				t:   tMock{put: pMock{err: nil}},
				ctx: context.Background(),
			},
			args: args{
				value: struct {
					Id  string
					Msg string
				}{"id", "msg"},
			},
			newId: func(value interface{}) string {
				return "id"
			},
			want:    "id",
			wantErr: false,
		},
		{
			name: "success: create node with id",
			fields: fields{
				id:    "id",
				ntype: "type",
				t:     tMock{put: pMock{err: nil}},
				ctx:   context.Background(),
			},
			args: args{
				value: struct {
					Id  string
					Msg string
				}{"id", "msg"},
			},
			want:    "id",
			wantErr: false,
		},
		{
			name: "fail: put error",
			fields: fields{
				id:    "id",
				ntype: "type",
				t:     tMock{put: pMock{err: errors.New("error")}},
				ctx:   context.Background(),
			},
			args: args{
				value: struct {
					Id  string
					Msg string
				}{"id", "msg"},
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := node{
				id:    tt.fields.id,
				ntype: tt.fields.ntype,
				t:     tt.fields.t,
				ctx:   tt.fields.ctx,
			}
			if tt.newId != nil {
				newId = tt.newId
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
		t     dynamoiface.Table
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
		{
			name: "success: update node",
			fields: fields{
				id:    "id",
				ntype: "type",
				t:     tMock{put: pMock{err: nil}},
				ctx:   nil,
			},
			args:    args{refId: "refId"},
			wantErr: false,
		},
		{
			name: "fail: empty id",
			fields: fields{
				id:    "",
				ntype: "type",
				t:     tMock{put: pMock{err: nil}},
				ctx:   nil,
			},
			args:    args{refId: "refId"},
			wantErr: true,
		},
		{
			name: "fail: put error",
			fields: fields{
				id:    "id",
				ntype: "type",
				t:     tMock{put: pMock{err: errors.New("error")}},
				ctx:   nil,
			},
			args:    args{refId: "refId"},
			wantErr: true,
		},
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
		t     dynamoiface.Table
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
		{
			name: "success: update node",
			fields: fields{
				id:    "id",
				ntype: "type",
				t:     tMock{update: uMock{err: nil}},
				ctx:   context.Background(),
			},
			args: args{
				value: struct {
					Id  string
					Msg string
				}{"id", "msg"},
			},
			wantErr: false,
		},
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
