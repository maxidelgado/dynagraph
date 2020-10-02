package utils

import (
	"github.com/guregu/dynamo"
	"reflect"
	"testing"
)

func TestFilter_GetHashValue(t *testing.T) {
	type fields struct {
		Id       string
		Type     string
		Operator Operator
		Index    Index
	}
	tests := []struct {
		name   string
		fields fields
		want   string
		want1  string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := Filter{
				Id:       tt.fields.Id,
				Type:     tt.fields.Type,
				Operator: tt.fields.Operator,
				Index:    tt.fields.Index,
			}
			got, got1 := k.GetHashValue()
			if got != tt.want {
				t.Errorf("GetHashValue() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetHashValue() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestFilter_GetRangeValues(t *testing.T) {
	type fields struct {
		Id       string
		Type     string
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := Filter{
				Id:       tt.fields.Id,
				Type:     tt.fields.Type,
				Operator: tt.fields.Operator,
				Index:    tt.fields.Index,
			}
			got, got1, got2 := k.GetRangeValues()
			if got != tt.want {
				t.Errorf("GetRangeValues() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetRangeValues() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("GetRangeValues() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func TestFilter_HashKey(t *testing.T) {
	type fields struct {
		Id       string
		Type     string
		Operator Operator
		Index    Index
	}
	tests := []struct {
		name   string
		fields fields
		want   interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := Filter{
				Id:       tt.fields.Id,
				Type:     tt.fields.Type,
				Operator: tt.fields.Operator,
				Index:    tt.fields.Index,
			}
			if got := k.HashKey(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HashKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilter_IndexKey(t *testing.T) {
	type fields struct {
		Id       string
		Type     string
		Operator Operator
		Index    Index
	}
	tests := []struct {
		name   string
		fields fields
		want   Index
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := Filter{
				Id:       tt.fields.Id,
				Type:     tt.fields.Type,
				Operator: tt.fields.Operator,
				Index:    tt.fields.Index,
			}
			if got := k.IndexKey(); got != tt.want {
				t.Errorf("IndexKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilter_OperatorKey(t *testing.T) {
	type fields struct {
		Id       string
		Type     string
		Operator Operator
		Index    Index
	}
	tests := []struct {
		name   string
		fields fields
		want   Operator
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := Filter{
				Id:       tt.fields.Id,
				Type:     tt.fields.Type,
				Operator: tt.fields.Operator,
				Index:    tt.fields.Index,
			}
			if got := k.OperatorKey(); got != tt.want {
				t.Errorf("OperatorKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilter_RangeKey(t *testing.T) {
	type fields struct {
		Id       string
		Type     string
		Operator Operator
		Index    Index
	}
	tests := []struct {
		name   string
		fields fields
		want   interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := Filter{
				Id:       tt.fields.Id,
				Type:     tt.fields.Type,
				Operator: tt.fields.Operator,
				Index:    tt.fields.Index,
			}
			if got := k.RangeKey(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RangeKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKeysInput_AppendKeys(t *testing.T) {
	type args struct {
		keys Filter
	}
	tests := []struct {
		name string
		i    KeysInput
		args args
		want KeysInput
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.AppendKeys(tt.args.keys); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AppendKeys() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWriteItemsInput_AppendEdge(t *testing.T) {
	type args struct {
		id    string
		value interface{}
	}
	tests := []struct {
		name    string
		i       WriteItemsInput
		args    args
		want    WriteItemsInput
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.i.AppendEdge(tt.args.id, tt.args.value)
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

func TestWriteItemsInput_AppendNode(t *testing.T) {
	type args struct {
		id    string
		value interface{}
	}
	tests := []struct {
		name    string
		i       WriteItemsInput
		args    args
		want    WriteItemsInput
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.i.AppendNode(tt.args.id, tt.args.value)
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

func TestWriteItemsInput_AppendProp(t *testing.T) {
	type args struct {
		id    string
		value interface{}
	}
	tests := []struct {
		name    string
		i       WriteItemsInput
		args    args
		want    WriteItemsInput
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.i.AppendProp(tt.args.id, tt.args.value)
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

func TestWriteItemsInput_AppendRef(t *testing.T) {
	type args struct {
		sourceId string
		targetId string
	}
	tests := []struct {
		name    string
		i       WriteItemsInput
		args    args
		want    WriteItemsInput
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.i.AppendRef(tt.args.sourceId, tt.args.targetId)
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
