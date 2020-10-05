package batch

import (
	"context"

	"github.com/maxidelgado/dynagraph/internal/dynamoiface"
	"github.com/maxidelgado/dynagraph/pkg/utils"
)

type Batch interface {
	Get(out interface{}, ids utils.IDs) error
	Put(operations utils.Operations) (int, error)
	Delete(ids utils.IDs) (int, error)
}

func New(ctx context.Context, t dynamoiface.Table) Batch {
	return batch{t: t, ctx: ctx}
}

type batch struct {
	t   dynamoiface.Table
	ctx context.Context
}

func (b batch) Put(inputs utils.Operations) (int, error) {
	return b.t.
		Batch().
		Write().
		Put(inputs...).
		RunWithContext(b.ctx)
}

func (b batch) Delete(keys utils.IDs) (int, error) {
	return b.t.
		Batch().
		Write().
		Delete(keys...).
		RunWithContext(b.ctx)
}

func (b batch) Get(out interface{}, keys utils.IDs) error {
	err := b.t.
		Batch("Id", "Type").
		Get(keys...).
		AllWithContext(b.ctx, out)
	if err != nil {
		return err
	}

	return nil
}
