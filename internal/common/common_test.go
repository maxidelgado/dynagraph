package common

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"reflect"
	"testing"
)

func TestAddEdge(t *testing.T) {
	type args struct {
		id    string
		value interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]*dynamodb.AttributeValue
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Edge(tt.args.id, tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("Edge() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Edge() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddProp(t *testing.T) {
	type args struct {
		id    string
		value interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]*dynamodb.AttributeValue
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Prop(tt.args.id, tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("Prop() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Prop() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddRef(t *testing.T) {
	type args struct {
		sourceId string
		targetId string
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]*dynamodb.AttributeValue
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Ref(tt.args.sourceId, tt.args.targetId)
			if (err != nil) != tt.wantErr {
				t.Errorf("Ref() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Ref() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBuildPut(t *testing.T) {
	type args struct {
		id    string
		ntype string
		value interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]*dynamodb.AttributeValue
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Put(tt.args.id, tt.args.ntype, tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("Put() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Put() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBuildUpdate(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]*dynamodb.AttributeValue
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Update(tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Update() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetNodeType(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Type(tt.args.value); got != tt.want {
				t.Errorf("Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewId(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Id(tt.args.value); got != tt.want {
				t.Errorf("Id() = %v, want %v", got, tt.want)
			}
		})
	}
}
