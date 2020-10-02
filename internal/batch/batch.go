package batch

import (
	"context"

	"github.com/guregu/dynamo"
	"github.com/maxidelgado/dynagraph/utils"
)

type Batch interface {
	Get(out interface{}, keys utils.KeysInput) error
	Put(inputs utils.WriteItemsInput) (int, error)
	Delete(keys utils.KeysInput) (int, error)
}

func New(ctx context.Context, t dynamo.Table) Batch {
	return batch{t: t, ctx: ctx}
}

type batch struct {
	t   dynamo.Table
	ctx context.Context
}

func (b batch) Put(inputs utils.WriteItemsInput) (int, error) {
	return b.t.
		Batch().
		Write().
		Put(inputs...).
		RunWithContext(b.ctx)
}

func (b batch) Delete(keys utils.KeysInput) (int, error) {
	return b.t.
		Batch().
		Write().
		Delete(keys...).
		RunWithContext(b.ctx)
}

func (b batch) Get(out interface{}, keys utils.KeysInput) error {
	err := b.t.
		Batch("Id", "Type").
		Get(keys...).
		AllWithContext(b.ctx, out)
	if err != nil && err != dynamo.ErrNotFound {
		return err
	}

	return nil
}
