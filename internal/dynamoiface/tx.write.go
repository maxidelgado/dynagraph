package dynamoiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/guregu/dynamo"
)

type (
	// WriteTx : Transaction to delete, put, update, and check items.
	WriteTx interface {
		// Delete : Adds a new delete operation to this transaction.
		Delete(d Delete) WriteTx
		// Put : Adds a put operation to this transaction.
		Put(p Put) WriteTx
		// Update : Adds an update operation to this transaction.
		Update(u Update) WriteTx
		// Check : Adds a conditional check to this transaction.
		Check(check *dynamo.ConditionCheck) WriteTx
		// Idempotent : Marks this transaction as idempotent when enabled is true.
		Idempotent(enabled bool) WriteTx
		// IdempotentWithToken : Marks this transaction as idempotent and explicitly specifies the token value.
		IdempotentWithToken(token string) WriteTx
		// ConsumedCapacity : Measure the throughput capacity consumed by this transaction and add it to cc.
		ConsumedCapacity(cc *dynamo.ConsumedCapacity) WriteTx
		// Run : Executes this transaction.
		Run() error
		// RunWithContext : Executes this transaction.
		RunWithContext(ctx aws.Context) error
	}

	writeTxWrap struct {
		writeTx *dynamo.WriteTx
	}
)

// Delete : Adds a new delete operation to this transaction.
func (wtw *writeTxWrap) Delete(d Delete) WriteTx {
	return &writeTxWrap{
		writeTx: wtw.writeTx.Delete(d.(*deleteWrap).delete),
	}
}

// Put : Adds a put operation to this transaction.
func (wtw *writeTxWrap) Put(p Put) WriteTx {
	return &writeTxWrap{
		writeTx: wtw.writeTx.Put(p.(*putWrap).put),
	}
}

// Update : Adds an update operation to this transaction.
func (wtw *writeTxWrap) Update(u Update) WriteTx {
	return &writeTxWrap{
		writeTx: wtw.writeTx.Update(u.(*updateWrap).update),
	}
}

// Check : Adds a conditional check to this transaction.
func (wtw *writeTxWrap) Check(check *dynamo.ConditionCheck) WriteTx {
	return &writeTxWrap{
		writeTx: wtw.writeTx.Check(check),
	}
}

// Idempotent : Marks this transaction as idempotent when enabled is true.
func (wtw *writeTxWrap) Idempotent(enabled bool) WriteTx {
	return &writeTxWrap{
		writeTx: wtw.writeTx.Idempotent(enabled),
	}
}

// IdempotentWithToken : Marks this transaction as idempotent and explicitly specifies the token value.
func (wtw *writeTxWrap) IdempotentWithToken(token string) WriteTx {
	return &writeTxWrap{
		writeTx: wtw.writeTx.IdempotentWithToken(token),
	}
}

// ConsumedCapacity : Measure the throughput capacity consumed by this transaction and add it to cc.
func (wtw *writeTxWrap) ConsumedCapacity(cc *dynamo.ConsumedCapacity) WriteTx {
	return &writeTxWrap{
		writeTx: wtw.writeTx.ConsumedCapacity(cc),
	}
}

// Run : Executes this transaction.
func (wtw *writeTxWrap) Run() error {
	return wtw.writeTx.Run()
}

// RunWithContext : Executes this transaction.
func (wtw *writeTxWrap) RunWithContext(ctx aws.Context) error {
	return wtw.writeTx.RunWithContext(ctx)
}
