package utils

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/guregu/dynamo"
	"reflect"
	"testing"
)

func TestID_HashKey(t *testing.T) {
	type fields struct {
		Id   string
		Type string
	}
	tests := []struct {
		name   string
		fields fields
		want   interface{}
	}{
		{
			name: "success: get hash key",
			fields: fields{
				Id:   "hash",
				Type: "range",
			},
			want: "hash",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := ID{
				Id:   tt.fields.Id,
				Type: tt.fields.Type,
			}
			if got := i.HashKey(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HashKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestID_RangeKey(t *testing.T) {
	type fields struct {
		Id   string
		Type string
	}
	tests := []struct {
		name   string
		fields fields
		want   interface{}
	}{
		{
			name: "success: get range key",
			fields: fields{
				Id:   "hash",
				Type: "range",
			},
			want: "range",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := ID{
				Id:   tt.fields.Id,
				Type: tt.fields.Type,
			}
			if got := i.RangeKey(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RangeKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIDs_AppendKeys(t *testing.T) {
	type args struct {
		keys dynamo.Keyed
	}
	tests := []struct {
		name string
		ids  IDs
		args args
		want IDs
	}{
		{
			name: "success: append key",
			ids:  IDs{ID{"hash", "range"}},
			args: args{
				keys: ID{"another hash", "another range"},
			},
			want: IDs{ID{"hash", "range"}, ID{"another hash", "another range"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ids.AppendKeys(tt.args.keys); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AppendKeys() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOperations_AppendEdge(t *testing.T) {
	type args struct {
		id    string
		value interface{}
	}
	tests := []struct {
		name       string
		operations Operations
		args       args
		want       Operations
		wantErr    bool
	}{
		{
			name:       "success: add edge",
			operations: Operations{},
			args: args{
				id: "id",
				value: struct {
					Id    string
					Value string
				}{"id", "value"},
			},
			want: Operations{
				map[string]*dynamodb.AttributeValue{
					"Id":    {S: aws.String("id")},
					"Type":  {S: aws.String("edge:id")},
					"Value": {S: aws.String("value")},
				},
			},
			wantErr: false,
		},
		{
			name:       "fail: not valid value",
			operations: Operations{},
			args: args{
				id:    "id",
				value: "not valid value",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.operations.AppendEdge(tt.args.id, tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("AppendEdge() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AppendEdge() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOperations_AppendNode(t *testing.T) {
	type TestType struct {
		Id    string
		Value string
	}
	type args struct {
		id    string
		value interface{}
	}
	tests := []struct {
		name       string
		operations Operations
		args       args
		want       Operations
		wantErr    bool
	}{
		{
			name:       "success: add node",
			operations: Operations{},
			args: args{
				id:    "id",
				value: TestType{"id", "value"},
			},
			want: Operations{
				map[string]*dynamodb.AttributeValue{
					"Id":    {S: aws.String("id")},
					"Type":  {S: aws.String("node:testtype")},
					"Value": {S: aws.String("value")},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.operations.AppendNode(tt.args.id, tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("AppendNode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AppendNode() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOperations_AppendProp(t *testing.T) {
	type TestType struct {
		Id    string
		Value string
	}
	type args struct {
		id    string
		value interface{}
	}
	tests := []struct {
		name       string
		operations Operations
		args       args
		want       Operations
		wantErr    bool
	}{
		{
			name:       "success: add property",
			operations: Operations{},
			args: args{
				id:    "id",
				value: TestType{"id", "value"},
			},
			want: Operations{
				map[string]*dynamodb.AttributeValue{
					"Id":    {S: aws.String("id")},
					"Type":  {S: aws.String("prop:testtype")},
					"Value": {S: aws.String("value")},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.operations.AppendProp(tt.args.id, tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("AppendProp() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AppendProp() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOperations_AppendRef(t *testing.T) {
	type args struct {
		sourceId string
		targetId string
	}
	tests := []struct {
		name       string
		operations Operations
		args       args
		want       Operations
		wantErr    bool
	}{
		{
			name:       "success: add reference",
			operations: Operations{},
			args: args{
				sourceId: "source",
				targetId: "target",
			},
			want: Operations{
				map[string]*dynamodb.AttributeValue{
					"Id":   {S: aws.String("source")},
					"Type": {S: aws.String("ref:target")},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.operations.AppendRef(tt.args.sourceId, tt.args.targetId)
			if (err != nil) != tt.wantErr {
				t.Errorf("AppendRef() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AppendRef() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQuery_GetHashSchema(t *testing.T) {
	type fields struct {
		ID       ID
		Operator Operator
		Index    Index
	}
	tests := []struct {
		name   string
		fields fields
		want   string
		want1  string
	}{
		{
			name: "success: get hash schema",
			fields: fields{
				ID: ID{"hash", "range"},
			},
			want:  "Id",
			want1: "hash",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := Query{
				ID:       tt.fields.ID,
				Operator: tt.fields.Operator,
				Index:    tt.fields.Index,
			}
			got, got1 := q.GetHashSchema()
			if got != tt.want {
				t.Errorf("GetHashSchema() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetHashSchema() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestQuery_GetRangeSchema(t *testing.T) {
	type fields struct {
		ID       ID
		Operator Operator
		Index    Index
	}
	tests := []struct {
		name   string
		fields fields
		want   string
		want1  dynamo.Operator
		want2  string
	}{
		{
			name: "success: get range schema",
			fields: fields{
				ID: ID{"hash", "range"},
			},
			want:  "Type",
			want1: "BEGINS_WITH",
			want2: "range",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := Query{
				ID:       tt.fields.ID,
				Operator: tt.fields.Operator,
				Index:    tt.fields.Index,
			}
			got, got1, got2 := q.GetRangeSchema()
			if got != tt.want {
				t.Errorf("GetRangeSchema() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetRangeSchema() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("GetRangeSchema() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func TestQuery_IndexKey(t *testing.T) {
	type fields struct {
		ID       ID
		Operator Operator
		Index    Index
	}
	tests := []struct {
		name   string
		fields fields
		want   Index
	}{
		{
			name: "success: get index key",
			fields: fields{
				ID: ID{"hash", "range"},
			},
			want: "default",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := Query{
				ID:       tt.fields.ID,
				Operator: tt.fields.Operator,
				Index:    tt.fields.Index,
			}
			if got := q.IndexKey(); got != tt.want {
				t.Errorf("IndexKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQuery_OperatorKey(t *testing.T) {
	type fields struct {
		ID       ID
		Operator Operator
		Index    Index
	}
	tests := []struct {
		name   string
		fields fields
		want   Operator
	}{
		{
			name: "success: get operator key",
			fields: fields{
				ID:       ID{"hash", "range"},
				Operator: Noop,
			},
			want: "Noop",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := Query{
				ID:       tt.fields.ID,
				Operator: tt.fields.Operator,
				Index:    tt.fields.Index,
			}
			if got := q.OperatorKey(); got != tt.want {
				t.Errorf("OperatorKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
