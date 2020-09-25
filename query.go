package dynagraph

import "github.com/guregu/dynamo"

type query struct {
	filter Filter
	table  dynamo.Table
	err    error
}

func (f query) One(out interface{}) error {
	if f.err != nil {
		return f.err
	}

	err := f.table.
		Get(f.filter.GetHashValue()).
		Range(nodeType, dynamo.Equal, f.filter.Type).
		One(out)
	if err != nil && err != dynamo.ErrNotFound {
		return err
	}

	return nil
}

func (f query) All(out interface{}) error {
	var err error
	if f.err != nil {
		return f.err
	}

	get := f.table.Get(f.filter.GetHashValue())

	if f.filter.Index == ByType {
		get = get.Index(string(f.filter.Index))
	}

	if f.filter.OperatorKey() == Noop {
		err = get.
			All(out)
	} else {
		err = get.
			Range(f.filter.GetRangeValues()).
			All(out)
	}

	if err != nil && err != dynamo.ErrNotFound {
		return err
	}

	return nil
}
