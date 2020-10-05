package dynamoiface

import (
	"github.com/guregu/dynamo"
)

type (
	// Batch : Stores the names of the hash key and range key for creating new batches.
	Batch interface {
		// Get : Creates a new batch get item request with the given keys.
		Get(keys ...dynamo.Keyed) BatchGet
		// Write : Creates a new batch write request, to which puts and deletes can be added.
		Write() BatchWrite
	}

	batchWrap struct {
		batch *dynamo.Batch
	}
)

// Get : Creates a new batch get item request with the given keys.
func (bw *batchWrap) Get(keys ...dynamo.Keyed) BatchGet {
	return &batchGetWrap{
		batchGet: bw.batch.Get(keys...),
	}
}

// Write : Creates a new batch write request, to which puts and deletes can be added.
func (bw *batchWrap) Write() BatchWrite {
	return &batchWriteWrap{
		batchWrite: bw.batch.Write(),
	}
}
