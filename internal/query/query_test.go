package query

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
		t   dynamo.Table
		f   utils.Filter
	}
	tests := []struct {
		name string
		args args
		want Query
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.ctx, tt.args.t, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_query_All(t *testing.T) {
	type fields struct {
		f   utils.Filter
		t   dynamo.Table
		ctx context.Context
	}
	type args struct {
		out interface{}
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
			f := query{
				f:   tt.fields.f,
				t:   tt.fields.t,
				ctx: tt.fields.ctx,
			}
			if err := f.All(tt.args.out); (err != nil) != tt.wantErr {
				t.Errorf("All() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_query_One(t *testing.T) {
	type fields struct {
		f   utils.Filter
		t   dynamo.Table
		ctx context.Context
	}
	type args struct {
		out interface{}
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
			f := query{
				f:   tt.fields.f,
				t:   tt.fields.t,
				ctx: tt.fields.ctx,
			}
			if err := f.One(tt.args.out); (err != nil) != tt.wantErr {
				t.Errorf("One() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
