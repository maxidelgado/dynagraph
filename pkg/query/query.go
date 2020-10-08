package query

import (
	"context"
	"github.com/maxidelgado/dynagraph/internal/common"
	"github.com/maxidelgado/dynagraph/internal/dynamoiface"

	"github.com/guregu/dynamo"
	"github.com/maxidelgado/dynagraph/pkg/utils"
)

func New(ctx context.Context, t dynamoiface.Table) Query {
	return query{
		t:   t,
		ctx: ctx,
	}
}

type Query interface {
	One(id utils.ID, out interface{}) error
	All(query utils.Query, out interface{}) error
}

type query struct {
	t   dynamoiface.Table
	ctx context.Context
}

func (f query) One(id utils.ID, out interface{}) error {
	err := f.t.
		Get(common.NodeId, id.HashKey()).
		Range(common.NodeType, dynamo.Equal, id.RangeKey()).
		OneWithContext(f.ctx, out)
	if err != nil {
		return err
	}

	return nil
}

func (f query) All(query utils.Query, out interface{}) error {
	var err error

	get := f.t.Get(query.GetHashSchema())

	if query.Index == utils.ByType {
		get = get.Index(string(query.Index))
	}

	if query.OperatorKey() == utils.Noop {
		err = get.
			AllWithContext(f.ctx, out)
	} else {
		err = get.
			Range(query.GetRangeSchema()).
			AllWithContext(f.ctx, out)
	}

	return err
}
