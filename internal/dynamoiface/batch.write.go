package dynamoiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/guregu/dynamo"
)

type (
	// BatchWrite : BatchWriteItem operation.
	BatchWrite interface {
		// Put : Adds put operations for items to this batch.
		Put(items ...interface{}) BatchWrite
		// Delete : Adds delete operations for the given keys to this batch.
		Delete(keys ...dynamo.Keyed) BatchWrite
		// ConsumedCapacity : Measures the throughput capacity consumed by this operation and add it to cc.
		ConsumedCapacity(cc *dynamo.ConsumedCapacity) BatchWrite
		// Run : Executes this batch.
		Run() (wrote int, err error)
		// RunWithContext : Executes this batch.
		RunWithContext(ctx aws.Context) (wrote int, err error)
	}

	batchWriteWrap struct {
		batchWrite *dynamo.BatchWrite
	}
)

// Put : Adds put operations for items to this batch.
func (bww *batchWriteWrap) Put(items ...interface{}) BatchWrite {
	return &batchWriteWrap{
		batchWrite: bww.batchWrite.Put(items...),
	}
}

// Delete : Adds delete operations for the given keys to this batch.
func (bww *batchWriteWrap) Delete(keys ...dynamo.Keyed) BatchWrite {
	return &batchWriteWrap{
		batchWrite: bww.batchWrite.Delete(keys...),
	}
}

// ConsumedCapacity : Measures the throughput capacity consumed by this operation and add it to cc.
func (bww *batchWriteWrap) ConsumedCapacity(cc *dynamo.ConsumedCapacity) BatchWrite {
	return &batchWriteWrap{
		batchWrite: bww.batchWrite.ConsumedCapacity(cc),
	}
}

// Run : Executes this batch.
func (bww *batchWriteWrap) Run() (wrote int, err error) {
	return bww.batchWrite.Run()
}

// RunWithContext : Executes this batch.
func (bww *batchWriteWrap) RunWithContext(ctx aws.Context) (wrote int, err error) {
	return bww.batchWrite.RunWithContext(ctx)
}
