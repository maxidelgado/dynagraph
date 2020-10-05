package dynamoiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/guregu/dynamo"
)

type (
	// GetTx : Transaction to retrieve items.
	GetTx interface {
		// Get : Adds a get request to this transaction.
		Get(q Query) GetTx
		// GetOne : Adds a get request to this transaction, and specifies out to which the results are marshaled.
		GetOne(q Query, out interface{}) GetTx
		// ConsumedCapacity : Measure the throughput capacity consumed by this transaction and add it to cc.
		ConsumedCapacity(cc *dynamo.ConsumedCapacity) GetTx
		// Run : Executes this transaction and unmarshals everything specified by GetOne.
		Run() error
		// RunWithContext : Executes this transaction and unmarshals everything specified by GetOne.
		RunWithContext(ctx aws.Context) error
		// All : Executes this transaction and unmarshals every value to out, which must be a pointer to a slice.
		All(out interface{}) error
		// AllWithContext : Executes this transaction and unmarshals every value to out, which must be a pointer to a slice.
		AllWithContext(ctx aws.Context, out interface{}) error
	}

	getTxWrap struct {
		getTx *dynamo.GetTx
	}
)

// Get : Adds a get request to this transaction.
func (gtw *getTxWrap) Get(q Query) GetTx {
	return &getTxWrap{
		getTx: gtw.getTx.Get(q.(*queryWrap).query),
	}
}

// GetOne : Adds a get request to this transaction, and specifies out to which the results are marshaled.
func (gtw *getTxWrap) GetOne(q Query, out interface{}) GetTx {
	return &getTxWrap{
		getTx: gtw.getTx.GetOne(q.(*queryWrap).query, out),
	}
}

// ConsumedCapacity : Measure the throughput capacity consumed by this transaction and add it to cc.
func (gtw *getTxWrap) ConsumedCapacity(cc *dynamo.ConsumedCapacity) GetTx {
	return &getTxWrap{
		getTx: gtw.getTx.ConsumedCapacity(cc),
	}
}

// Run : Executes this transaction and unmarshals everything specified by GetOne.
func (gtw *getTxWrap) Run() error {
	return gtw.getTx.Run()
}

// RunWithContext : Executes this transaction and unmarshals everything specified by GetOne.
func (gtw *getTxWrap) RunWithContext(ctx aws.Context) error {
	return gtw.getTx.RunWithContext(ctx)
}

// All : Executes this transaction and unmarshals every value to out, which must be a pointer to a slice.
func (gtw *getTxWrap) All(out interface{}) error {
	return gtw.getTx.All(out)
}

// AllWithContext : Executes this transaction and unmarshals every value to out, which must be a pointer to a slice.
func (gtw *getTxWrap) AllWithContext(ctx aws.Context, out interface{}) error {
	return gtw.getTx.AllWithContext(ctx, out)
}
