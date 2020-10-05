package dynamoiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/guregu/dynamo"
)

type (
	// DescribeTTL : Request to obtain details about a table's time to live configuration.
	DescribeTTL interface {
		// Run : Executes this request and returns details about time to live, or an error.
		Run() (dynamo.TTLDescription, error)
		// RunWithContext : Executes this request and returns details about time to live, or an error.
		RunWithContext(ctx aws.Context) (dynamo.TTLDescription, error)
	}

	describeTTLWrap struct {
		describeTTL *dynamo.DescribeTTL
	}
)

// Run : Executes this request and returns details about time to live, or an error.
func (dtw *describeTTLWrap) Run() (dynamo.TTLDescription, error) {
	return dtw.describeTTL.Run()
}

// RunWithContext : Executes this request and returns details about time to live, or an error.
func (dtw *describeTTLWrap) RunWithContext(ctx aws.Context) (dynamo.TTLDescription, error) {
	return dtw.describeTTL.RunWithContext(ctx)
}
