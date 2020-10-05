package dynamoiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/guregu/dynamo"
)

type (
	// Delete : Request to delete an item.
	Delete interface {
		// Range : Specifies the range key (a.k.a. sort key) to delete.
		Range(name string, value interface{}) Delete
		// If : Specifies a conditional expression for this delete to succeed.
		If(expr string, args ...interface{}) Delete
		// ConsumedCapacity : Measures the throughput capacity consumed by this operation and add it to cc.
		ConsumedCapacity(cc *dynamo.ConsumedCapacity) Delete
		// Run : Executes this delete request.
		Run() error
		// RunWithContext : Executes this delete request.
		RunWithContext(ctx aws.Context) error
		// OldValue : Executes this delete request, unmarshaling the previous value to out.
		OldValue(out interface{}) error
		// OldValueWithContext : Executes this delete request, unmarshaling the previous value to out.
		OldValueWithContext(ctx aws.Context, out interface{}) error
	}

	deleteWrap struct {
		delete *dynamo.Delete
	}
)

// Range : Specifies the range key (a.k.a. sort key) to delete.
func (dw *deleteWrap) Range(name string, value interface{}) Delete {
	return &deleteWrap{
		delete: dw.delete.Range(name, value),
	}
}

// If : Specifies a conditional expression for this delete to succeed.
func (dw *deleteWrap) If(expr string, args ...interface{}) Delete {
	return &deleteWrap{
		delete: dw.delete.If(expr, args...),
	}
}

// ConsumedCapacity : Measures the throughput capacity consumed by this operation and add it to cc.
func (dw *deleteWrap) ConsumedCapacity(cc *dynamo.ConsumedCapacity) Delete {
	return &deleteWrap{
		delete: dw.delete.ConsumedCapacity(cc),
	}
}

// Run : Executes this delete request.
func (dw *deleteWrap) Run() error {
	return dw.delete.Run()
}

// RunWithContext : Executes this delete request.
func (dw *deleteWrap) RunWithContext(ctx aws.Context) error {
	return dw.delete.RunWithContext(ctx)
}

// OldValue : Executes this delete request, unmarshaling the previous value to out.
func (dw *deleteWrap) OldValue(out interface{}) error {
	return dw.delete.OldValue(out)
}

// OldValueWithContext : Executes this delete request, unmarshaling the previous value to out.
func (dw *deleteWrap) OldValueWithContext(ctx aws.Context, out interface{}) error {
	return dw.delete.OldValueWithContext(ctx, out)
}
