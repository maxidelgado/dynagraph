package dynamoiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/guregu/dynamo"
)

type (
	// UpdateTable : Request to change a table's settings.
	UpdateTable interface {
		// OnDemand : Sets this table to use on-demand (pay per request) billing mode if enabled is true.
		OnDemand(enabled bool) UpdateTable
		// Provision : Sets this table's read and write throughput capacity.
		Provision(read, write int64) UpdateTable
		// ProvisionIndex : Updates a global secondary index's read and write throughput capacity.
		ProvisionIndex(name string, read, write int64) UpdateTable
		// CreateIndex : Adds a new secondary global index.
		CreateIndex(index dynamo.Index) UpdateTable
		// DeleteIndex : Deletes the specified index.
		DeleteIndex(name string) UpdateTable
		// Stream : Enables streaming and sets the stream view type.
		Stream(view dynamo.StreamView) UpdateTable
		// DisableStream : Disables this table's stream.
		DisableStream() UpdateTable
		// Run : Executes this request and describes the table.
		Run() (dynamo.Description, error)
		// RunWithContext : Executes this request and describes the table.
		RunWithContext(ctx aws.Context) (dynamo.Description, error)
	}

	updateTableWrap struct {
		updateTable *dynamo.UpdateTable
	}
)

// OnDemand : Sets this table to use on-demand (pay per request) billing mode if enabled is true.
func (utw *updateTableWrap) OnDemand(enabled bool) UpdateTable {
	return &updateTableWrap{
		updateTable: utw.updateTable.OnDemand(enabled),
	}
}

// Provision : Sets this table's read and write throughput capacity.
func (utw *updateTableWrap) Provision(read, write int64) UpdateTable {
	return &updateTableWrap{
		updateTable: utw.updateTable.Provision(read, write),
	}
}

// ProvisionIndex : Updates a global secondary index's read and write throughput capacity.
func (utw *updateTableWrap) ProvisionIndex(name string, read, write int64) UpdateTable {
	return &updateTableWrap{
		updateTable: utw.updateTable.ProvisionIndex(name, read, write),
	}
}

// CreateIndex : Adds a new secondary global index.
func (utw *updateTableWrap) CreateIndex(index dynamo.Index) UpdateTable {
	return &updateTableWrap{
		updateTable: utw.updateTable.CreateIndex(index),
	}
}

// DeleteIndex : Deletes the specified index.
func (utw *updateTableWrap) DeleteIndex(name string) UpdateTable {
	return &updateTableWrap{
		updateTable: utw.updateTable.DeleteIndex(name),
	}
}

// Stream : Enables streaming and sets the stream view type.
func (utw *updateTableWrap) Stream(view dynamo.StreamView) UpdateTable {
	return &updateTableWrap{
		updateTable: utw.updateTable.Stream(view),
	}
}

// DisableStream : Disables this table's stream.
func (utw *updateTableWrap) DisableStream() UpdateTable {
	return &updateTableWrap{
		updateTable: utw.updateTable.DisableStream(),
	}
}

// Run : Executes this request and describes the table.
func (utw *updateTableWrap) Run() (dynamo.Description, error) {
	return utw.updateTable.Run()
}

// RunWithContext : Executes this request and describes the table.
func (utw *updateTableWrap) RunWithContext(ctx aws.Context) (dynamo.Description, error) {
	return utw.updateTable.RunWithContext(ctx)
}
