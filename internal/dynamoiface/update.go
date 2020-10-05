package dynamoiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/guregu/dynamo"
)

type (
	// Update : Represents changes to an existing item.
	Update interface {
		// Range : Specifies the range key (sort key) for the item to update.
		Range(name string, value interface{}) Update
		// Set : Changes path to the given value.
		Set(path string, value interface{}) Update
		// SetNullable : Changes path to the given value, allowing empty and nil values.
		SetNullable(path string, value interface{}) Update
		// SetSet : Changes a set at the given path to the given value.
		SetSet(path string, value interface{}) Update
		// SetIfNotExists : Changes path to the given value, if it does not already exist.
		SetIfNotExists(path string, value interface{}) Update
		// SetExpr : Performs a custom set expression, substituting the args into expr as in filter expressions.
		SetExpr(expr string, args ...interface{}) Update
		// Append : Appends value to the end of the list specified by path.
		Append(path string, value interface{}) Update
		// Prepend : Inserts value to the beginning of the list specified by path.
		Prepend(path string, value interface{}) Update
		// Add : Adds value to path.
		Add(path string, value interface{}) Update
		// AddStringsToSet : Adds the given values to the string set specified by path.
		AddStringsToSet(path string, values ...string) Update
		// AddIntsToSet : Adds the given values to the number set specified by path.
		AddIntsToSet(path string, values ...int) Update
		// AddFloatsToSet : Adds the given values to the number set specified by path.
		AddFloatsToSet(path string, values ...float64) Update
		// DeleteStringsFromSet : Deletes the given values from the string set specified by path.
		DeleteStringsFromSet(path string, values ...string) Update
		// DeleteIntsFromSet : Deletes the given values from the number set specified by path.
		DeleteIntsFromSet(path string, values ...int) Update
		// DeleteFloatsFromSet : Deletes the given values from the number set specified by path.
		DeleteFloatsFromSet(path string, values ...float64) Update
		// Remove : Removes the paths from this item, deleting the specified attributes.
		Remove(paths ...string) Update
		// RemoveExpr : Performs a custom remove expression, substituting the args into expr as in filter expressions.
		RemoveExpr(expr string, args ...interface{}) Update
		// If : Specifies a conditional expression for this update to succeed.
		If(expr string, args ...interface{}) Update
		// ConsumedCapacity : Measure the throughput capacity consumed by this operation and add it to cc.
		ConsumedCapacity(cc *dynamo.ConsumedCapacity) Update
		// Run : Executes this update.
		Run() error
		// RunWithContext : Executes this update.
		RunWithContext(ctx aws.Context) error
		// Value : Executes this update, encoding out with the new value.
		Value(out interface{}) error
		// ValueWithContext : Executes this update, encoding out with the new value.
		ValueWithContext(ctx aws.Context, out interface{}) error
		// OldValue : Executes this update, encoding out with the previous value.
		OldValue(out interface{}) error
		// OldValueWithContext : Executes this update, encoding out with the previous value.
		OldValueWithContext(ctx aws.Context, out interface{}) error
	}

	updateWrap struct {
		update *dynamo.Update
	}
)

// Range : Specifies the range key (sort key) for the item to update.
func (uw *updateWrap) Range(name string, value interface{}) Update {
	return &updateWrap{
		update: uw.update.Range(name, value),
	}
}

// Set : Changes path to the given value.
func (uw *updateWrap) Set(path string, value interface{}) Update {
	return &updateWrap{
		update: uw.update.Set(path, value),
	}
}

// SetNullable : Changes path to the given value, allowing empty and nil values.
func (uw *updateWrap) SetNullable(path string, value interface{}) Update {
	return &updateWrap{
		update: uw.update.SetNullable(path, value),
	}
}

// SetSet : Changes a set at the given path to the given value.
func (uw *updateWrap) SetSet(path string, value interface{}) Update {
	return &updateWrap{
		update: uw.update.SetSet(path, value),
	}
}

// SetIfNotExists : Changes path to the given value, if it does not already exist.
func (uw *updateWrap) SetIfNotExists(path string, value interface{}) Update {
	return &updateWrap{
		update: uw.update.SetIfNotExists(path, value),
	}
}

// SetExpr : Performs a custom set expression, substituting the args into expr as in filter expressions.
func (uw *updateWrap) SetExpr(expr string, args ...interface{}) Update {
	return &updateWrap{
		update: uw.update.SetExpr(expr, args...),
	}
}

// Append : Appends value to the end of the list specified by path.
func (uw *updateWrap) Append(path string, value interface{}) Update {
	return &updateWrap{
		update: uw.update.Append(path, value),
	}
}

// Prepend : Inserts value to the beginning of the list specified by path.
func (uw *updateWrap) Prepend(path string, value interface{}) Update {
	return &updateWrap{
		update: uw.update.Prepend(path, value),
	}
}

// Add : Adds value to path.
func (uw *updateWrap) Add(path string, value interface{}) Update {
	return &updateWrap{
		update: uw.update.Add(path, value),
	}
}

// AddStringsToSet : Adds the given values to the string set specified by path.
func (uw *updateWrap) AddStringsToSet(path string, values ...string) Update {
	return &updateWrap{
		update: uw.update.AddStringsToSet(path, values...),
	}
}

// AddIntsToSet : Adds the given values to the number set specified by path.
func (uw *updateWrap) AddIntsToSet(path string, values ...int) Update {
	return &updateWrap{
		update: uw.update.AddIntsToSet(path, values...),
	}
}

// AddFloatsToSet : Adds the given values to the number set specified by path.
func (uw *updateWrap) AddFloatsToSet(path string, values ...float64) Update {
	return &updateWrap{
		update: uw.update.AddFloatsToSet(path, values...),
	}
}

// DeleteStringsFromSet : Deletes the given values from the string set specified by path.
func (uw *updateWrap) DeleteStringsFromSet(path string, values ...string) Update {
	return &updateWrap{
		update: uw.update.DeleteStringsFromSet(path, values...),
	}
}

// DeleteIntsFromSet : Deletes the given values from the number set specified by path.
func (uw *updateWrap) DeleteIntsFromSet(path string, values ...int) Update {
	return &updateWrap{
		update: uw.update.DeleteIntsFromSet(path, values...),
	}
}

// DeleteFloatsFromSet : Deletes the given values from the number set specified by path.
func (uw *updateWrap) DeleteFloatsFromSet(path string, values ...float64) Update {
	return &updateWrap{
		update: uw.update.DeleteFloatsFromSet(path, values...),
	}
}

// Remove : Removes the paths from this item, deleting the specified attributes.
func (uw *updateWrap) Remove(paths ...string) Update {
	return &updateWrap{
		update: uw.update.Remove(paths...),
	}
}

// RemoveExpr : Performs a custom remove expression, substituting the args into expr as in filter expressions.
func (uw *updateWrap) RemoveExpr(expr string, args ...interface{}) Update {
	return &updateWrap{
		update: uw.update.RemoveExpr(expr, args...),
	}
}

// If : Specifies a conditional expression for this update to succeed.
func (uw *updateWrap) If(expr string, args ...interface{}) Update {
	return &updateWrap{
		update: uw.update.If(expr, args...),
	}
}

// ConsumedCapacity : Measure the throughput capacity consumed by this operation and add it to cc.
func (uw *updateWrap) ConsumedCapacity(cc *dynamo.ConsumedCapacity) Update {
	return &updateWrap{
		update: uw.update.ConsumedCapacity(cc),
	}
}

// Run : Executes this update.
func (uw *updateWrap) Run() error {
	return uw.update.Run()
}

// RunWithContext : Executes this update.
func (uw *updateWrap) RunWithContext(ctx aws.Context) error {
	return uw.update.RunWithContext(ctx)
}

// Value : Executes this update, encoding out with the new value.
func (uw *updateWrap) Value(out interface{}) error {
	return uw.update.Value(out)
}

// ValueWithContext : Executes this update, encoding out with the new value.
func (uw *updateWrap) ValueWithContext(ctx aws.Context, out interface{}) error {
	return uw.update.ValueWithContext(ctx, out)
}

// OldValue : Executes this update, encoding out with the previous value.
func (uw *updateWrap) OldValue(out interface{}) error {
	return uw.update.OldValue(out)
}

// OldValueWithContext : Executes this update, encoding out with the previous value.
func (uw *updateWrap) OldValueWithContext(ctx aws.Context, out interface{}) error {
	return uw.update.OldValueWithContext(ctx, out)
}
