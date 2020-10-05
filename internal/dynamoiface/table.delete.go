package dynamoiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/guregu/dynamo"
)

type (
	// DeleteTable : Request to delete a table.
	DeleteTable interface {
		// Run : Executes this request and deletes the table.
		Run() error
		// RunWithContext : Executes this request and deletes the table.
		RunWithContext(ctx aws.Context) error
	}

	deleteTableWrap struct {
		deleteTable *dynamo.DeleteTable
	}
)

// Run : Executes this request and deletes the table.
func (dtw *deleteTableWrap) Run() error {
	return dtw.deleteTable.Run()
}

// RunWithContext : Executes this request and deletes the table.
func (dtw *deleteTableWrap) RunWithContext(ctx aws.Context) error {
	return dtw.deleteTable.RunWithContext(ctx)
}
