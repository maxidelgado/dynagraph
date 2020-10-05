package dynamoiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/guregu/dynamo"
)

type (
	// CreateTable : Request to create a new table.
	CreateTable interface {
		// OnDemand : Specifies to create the table with on-demand (pay per request) billing mode.
		OnDemand(enabled bool) CreateTable
		// Provision : Specifies the provisioned read and write capacity for this table.
		Provision(readUnits, writeUnits int64) CreateTable
		// ProvisionIndex : pecifies the provisioned read and write capacity for the given global secondary index.
		ProvisionIndex(index string, readUnits, writeUnits int64) CreateTable
		// Stream : Enables DynamoDB Streams for this table which the specified type of view.
		Stream(view dynamo.StreamView) CreateTable
		// Project : Specifies the projection type for the given table.
		Project(index string, projection dynamo.IndexProjection, includeAttribs ...string) CreateTable
		// Index : Specifies an index to add to this table.
		Index(index dynamo.Index) CreateTable
		// Tag : Specifies a metadata tag for this table. Multiple tags may be specified.
		Tag(key, value string) CreateTable
		// Run : Creates this table or returns and error.
		Run() error
		// RunWithContext : Creates this table or returns and error.
		RunWithContext(ctx aws.Context) error
	}

	createTableWrap struct {
		createTable *dynamo.CreateTable
	}
)

// OnDemand : Specifies to create the table with on-demand (pay per request) billing mode.
func (ctw *createTableWrap) OnDemand(enabled bool) CreateTable {
	return &createTableWrap{
		createTable: ctw.createTable.OnDemand(enabled),
	}
}

// Provision : Specifies the provisioned read and write capacity for this table.
func (ctw *createTableWrap) Provision(readUnits, writeUnits int64) CreateTable {
	return &createTableWrap{
		createTable: ctw.createTable.Provision(readUnits, writeUnits),
	}
}

// ProvisionIndex : pecifies the provisioned read and write capacity for the given global secondary index.
// Local secondary indices share their capacity with the table.
func (ctw *createTableWrap) ProvisionIndex(index string, readUnits, writeUnits int64) CreateTable {
	return &createTableWrap{
		createTable: ctw.createTable.ProvisionIndex(index, readUnits, writeUnits),
	}
}

// Stream : Enables DynamoDB Streams for this table which the specified type of view.
func (ctw *createTableWrap) Stream(view dynamo.StreamView) CreateTable {
	return &createTableWrap{
		createTable: ctw.createTable.Stream(view),
	}
}

// Project : Specifies the projection type for the given table.
func (ctw *createTableWrap) Project(index string, projection dynamo.IndexProjection, includeAttribs ...string) CreateTable {
	return &createTableWrap{
		createTable: ctw.createTable.Project(index, projection, includeAttribs...),
	}
}

// Index : Specifies an index to add to this table.
func (ctw *createTableWrap) Index(index dynamo.Index) CreateTable {
	return &createTableWrap{
		createTable: ctw.createTable.Index(index),
	}
}

// Tag : Specifies a metadata tag for this table. Multiple tags may be specified.
func (ctw *createTableWrap) Tag(key, value string) CreateTable {
	return &createTableWrap{
		createTable: ctw.createTable.Tag(key, value),
	}
}

// Run : Creates this table or returns and error.
func (ctw *createTableWrap) Run() error {
	return ctw.createTable.Run()
}

// RunWithContext : Creates this table or returns and error.
func (ctw *createTableWrap) RunWithContext(ctx aws.Context) error {
	return ctw.createTable.RunWithContext(ctx)
}
