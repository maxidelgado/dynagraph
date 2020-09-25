package dynagraph

import (
	"github.com/guregu/dynamo"
)

type batch struct {
	table dynamo.Table
}

func (b batch) Put(inputs WriteItemsInput) (int, error) {
	return b.table.
		Batch().
		Write().
		Put(inputs...).
		Run()
}

func (b batch) Delete(keys KeysInput) (int, error) {
	return b.table.
		Batch().
		Write().
		Delete(keys...).
		Run()
}

func (b batch) Get(out interface{}, keys KeysInput) error {
	err := b.table.
		Batch("Id", "Type").
		Get(keys...).
		All(out)
	if err != nil && err != dynamo.ErrNotFound {
		return err
	}

	return nil
}
