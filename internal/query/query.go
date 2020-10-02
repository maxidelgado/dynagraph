package query

import (
	"context"

	"github.com/guregu/dynamo"
	"github.com/maxidelgado/dynagraph/utils"
)

func New(ctx context.Context, t dynamo.Table, f utils.Filter) Query {
	return query{
		f:   f,
		t:   t,
		ctx: ctx,
	}
}

type Query interface {
	One(out interface{}) error
	All(out interface{}) error
}

type query struct {
	f   utils.Filter
	t   dynamo.Table
	ctx context.Context
}

func (f query) One(out interface{}) error {
	err := f.t.
		Get(f.f.GetHashValue()).
		Range(utils.NodeType, dynamo.Equal, f.f.Type).
		OneWithContext(f.ctx, out)
	if err != nil {
		return err
	}

	return nil
}

func (f query) All(out interface{}) error {
	var err error

	get := f.t.Get(f.f.GetHashValue())

	if f.f.Index == utils.ByType {
		get = get.Index(string(f.f.Index))
	}

	if f.f.OperatorKey() == utils.Noop {
		err = get.
			AllWithContext(f.ctx, out)
	} else {
		err = get.
			Range(f.f.GetRangeValues()).
			AllWithContext(f.ctx, out)
	}

	if err != nil {
		return err
	}

	return nil
}
