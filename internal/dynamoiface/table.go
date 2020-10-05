package dynamoiface

import (
	"github.com/guregu/dynamo"
)

type (
	// Table : DynamoDB table.
	Table interface {
		// Name : Returns this table's name.
		Name() string
		// Batch : Creates a new batch with the given hash key name, and range key name if provided.
		Batch(hashAndRangeKeyName ...string) Batch
		// Check : Creates a new ConditionCheck, which represents a condition for a write transaction to succeed.
		Check(hashKey string, value interface{}) ConditionCheck
		// Delete : Creates a new request to delete an item.
		Delete(name string, value interface{}) Delete
		// Describe : Begins a new request to describe this table.
		Describe() DescribeTable
		// Put : Creates a new request to create or replace an item.
		Put(item interface{}) Put
		// Get : Creates a new request to get an item.
		Get(name string, value interface{}) Query
		// Scan : Creates a new request to scan this table.
		Scan() Scan
		// DeleteTable : Begins a new request to delete this table.
		DeleteTable() DeleteTable
		// UpdateTTL : Begins a new request to enable or disable this table's time to live.
		UpdateTTL(attribute string, enabled bool) UpdateTTL
		// DescribeTTL : Begins a new request to obtain details about this table's time to live configuration.
		DescribeTTL() DescribeTTL
		// Update : Creates a new request to modify an existing item.
		Update(hashKey string, value interface{}) Update
		// UpdateTable : Makes changes to this table's settings.
		UpdateTable() UpdateTable
	}

	tableWrap struct {
		table *dynamo.Table
	}
)

// Name : Returns this table's name.
func (tw *tableWrap) Name() string {
	return tw.table.Name()
}

// Batch : Creates a new batch with the given hash key name, and range key name if provided.
func (tw *tableWrap) Batch(hashAndRangeKeyName ...string) Batch {
	batch := tw.table.Batch(hashAndRangeKeyName...)

	return &batchWrap{
		batch: &batch,
	}
}

// Check : Creates a new ConditionCheck, which represents a condition for a write transaction to succeed.
func (tw *tableWrap) Check(hashKey string, value interface{}) ConditionCheck {
	return &conditionCheckWrap{
		conditionCheck: tw.table.Check(hashKey, value),
	}
}

// Delete : Creates a new request to delete an item.
func (tw *tableWrap) Delete(name string, value interface{}) Delete {
	return &deleteWrap{
		delete: tw.table.Delete(name, value),
	}
}

// Describe : Begins a new request to describe this table.
func (tw *tableWrap) Describe() DescribeTable {
	return &describeTableWrap{
		describeTable: tw.table.Describe(),
	}
}

// Put : Creates a new request to create or replace an item.
func (tw *tableWrap) Put(item interface{}) Put {
	return &putWrap{
		put: tw.table.Put(item),
	}
}

// Get : Creates a new request to get an item.
func (tw *tableWrap) Get(name string, value interface{}) Query {
	return &queryWrap{
		query: tw.table.Get(name, value),
	}
}

// Scan : Creates a new request to scan this table.
func (tw *tableWrap) Scan() Scan {
	return &scanWrap{
		scan: tw.table.Scan(),
	}
}

// DeleteTable : Begins a new request to delete this table.
func (tw *tableWrap) DeleteTable() DeleteTable {
	return &deleteTableWrap{
		deleteTable: tw.table.DeleteTable(),
	}
}

// UpdateTTL : Begins a new request to enable or disable this table's time to live.
func (tw *tableWrap) UpdateTTL(attribute string, enabled bool) UpdateTTL {
	return &updateTTLWrap{
		updateTTL: tw.table.UpdateTTL(attribute, enabled),
	}
}

// DescribeTTL : Begins a new request to obtain details about this table's time to live configuration.
func (tw *tableWrap) DescribeTTL() DescribeTTL {
	return &describeTTLWrap{
		describeTTL: tw.table.DescribeTTL(),
	}
}

// Update : Creates a new request to modify an existing item.
func (tw *tableWrap) Update(hashKey string, value interface{}) Update {
	return &updateWrap{
		update: tw.table.Update(hashKey, value),
	}
}

// UpdateTable : Makes changes to this table's settings.
func (tw *tableWrap) UpdateTable() UpdateTable {
	return &updateTableWrap{
		updateTable: tw.table.UpdateTable(),
	}
}
