package dynamoiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/guregu/dynamo"
)

type (
	// BatchGet : BatchGetItem operation.
	BatchGet interface {
		// And : Adds more keys to be gotten.
		And(keys ...dynamo.Keyed) BatchGet
		// Consistent : Set the read consistency to strong or not.
		Consistent(on bool) BatchGet
		// ConsumedCapacity : Measures the throughput capacity consumed by this operation and add it to cc.
		ConsumedCapacity(cc *dynamo.ConsumedCapacity) BatchGet
		// All : Executes this request and unmarshals all results to out, which must be a pointer to a slice.
		All(out interface{}) error
		// AllWithContext : Executes this request and unmarshals all results to out, which must be a pointer to a slice.
		AllWithContext(ctx aws.Context, out interface{}) error
		// Iter : Returns a results iterator for this batch.
		Iter() dynamo.Iter
	}

	batchGetWrap struct {
		batchGet *dynamo.BatchGet
	}
)

// And : Adds more keys to be gotten.
func (bgw *batchGetWrap) And(keys ...dynamo.Keyed) BatchGet {
	return &batchGetWrap{
		batchGet: bgw.batchGet.And(keys...),
	}
}

// Consistent : Set the read consistency to strong or not.
func (bgw *batchGetWrap) Consistent(on bool) BatchGet {
	return &batchGetWrap{
		batchGet: bgw.batchGet.Consistent(on),
	}
}

// ConsumedCapacity : Measures the throughput capacity consumed by this operation and add it to cc.
func (bgw *batchGetWrap) ConsumedCapacity(cc *dynamo.ConsumedCapacity) BatchGet {
	return &batchGetWrap{
		batchGet: bgw.batchGet.ConsumedCapacity(cc),
	}
}

// All : Executes this request and unmarshals all results to out, which must be a pointer to a slice.
func (bgw *batchGetWrap) All(out interface{}) error {
	return bgw.batchGet.All(out)
}

// AllWithContext : Executes this request and unmarshals all results to out, which must be a pointer to a slice.
func (bgw *batchGetWrap) AllWithContext(ctx aws.Context, out interface{}) error {
	return bgw.batchGet.AllWithContext(ctx, out)
}

// Iter : Returns a results iterator for this batch.
func (bgw *batchGetWrap) Iter() dynamo.Iter {
	return bgw.batchGet.Iter()
}
