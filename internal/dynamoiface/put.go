package dynamoiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/guregu/dynamo"
)

type (
	// Put : Request to create or replace an item.
	Put interface {
		// If : Specifies a conditional expression for this put to succeed.
		If(expr string, args ...interface{}) Put
		// ConsumedCapacity : Measures the throughput capacity consumed by this operation and add it to cc.
		ConsumedCapacity(cc *dynamo.ConsumedCapacity) Put
		// Run : Executes this put.
		Run() error
		// RunWithContext : Executes this put.
		RunWithContext(ctx aws.Context) error
		// OldValue : Executes this put, unmarshaling the previous value into out.
		OldValue(out interface{}) error
		// OldValueWithContext : Executes this put, unmarshaling the previous value into out.
		OldValueWithContext(ctx aws.Context, out interface{}) error
	}

	putWrap struct {
		put *dynamo.Put
	}
)

// If : Specifies a conditional expression for this put to succeed.
func (pw *putWrap) If(expr string, args ...interface{}) Put {
	return &putWrap{
		put: pw.put.If(expr, args...),
	}
}

// ConsumedCapacity : Measures the throughput capacity consumed by this operation and add it to cc.
func (pw *putWrap) ConsumedCapacity(cc *dynamo.ConsumedCapacity) Put {
	return &putWrap{
		put: pw.put.ConsumedCapacity(cc),
	}
}

// Run : Executes this put.
func (pw *putWrap) Run() error {
	return pw.put.Run()
}

// RunWithContext : Executes this put.
func (pw *putWrap) RunWithContext(ctx aws.Context) error {
	return pw.put.RunWithContext(ctx)
}

// OldValue : Executes this put, unmarshaling the previous value into out.
func (pw *putWrap) OldValue(out interface{}) error {
	return pw.put.OldValue(out)
}

// OldValueWithContext : Executes this put, unmarshaling the previous value into out.
func (pw *putWrap) OldValueWithContext(ctx aws.Context, out interface{}) error {
	return pw.put.OldValueWithContext(ctx, out)
}
