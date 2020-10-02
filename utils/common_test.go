package utils

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
			got, err := AddEdge(tt.args.id, tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddEdge() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddEdge() got = %v, want %v", got, tt.want)
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
			got, err := AddProp(tt.args.id, tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddProp() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddProp() got = %v, want %v", got, tt.want)
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
			got, err := AddRef(tt.args.sourceId, tt.args.targetId)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddRef() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddRef() got = %v, want %v", got, tt.want)
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
			got, err := BuildPut(tt.args.id, tt.args.ntype, tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("BuildPut() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BuildPut() got = %v, want %v", got, tt.want)
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
			got, err := BuildUpdate(tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("BuildUpdate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BuildUpdate() got = %v, want %v", got, tt.want)
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
			if got := GetNodeType(tt.args.value); got != tt.want {
				t.Errorf("GetNodeType() = %v, want %v", got, tt.want)
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
			if got := NewId(tt.args.value); got != tt.want {
				t.Errorf("NewId() = %v, want %v", got, tt.want)
			}
		})
	}
}
