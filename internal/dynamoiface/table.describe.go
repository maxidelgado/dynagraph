package dynamoiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/guregu/dynamo"
)

type (
	// DescribeTable : Request for information about a table and its indexes.
	DescribeTable interface {
		// Run : Executes this request and describe the table.
		Run() (dynamo.Description, error)
		// RunWithContext : Executes this request and describe the table.
		RunWithContext(ctx aws.Context) (dynamo.Description, error)
	}

	describeTableWrap struct {
		describeTable *dynamo.DescribeTable
	}
)

// Run : Executes this request and describe the table.
func (dtw *describeTableWrap) Run() (dynamo.Description, error) {
	return dtw.describeTable.Run()
}

// RunWithContext : Executes this request and describe the table.
func (dtw *describeTableWrap) RunWithContext(ctx aws.Context) (dynamo.Description, error) {
	return dtw.describeTable.RunWithContext(ctx)
}
