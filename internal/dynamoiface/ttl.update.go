package dynamoiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/guregu/dynamo"
)

type (
	// UpdateTTL : Request to enable or disable a table's time to live functionality.
	UpdateTTL interface {
		// Run : Executes this request.
		Run() error
		// RunWithContext : Executes this request.
		RunWithContext(ctx aws.Context) error
	}

	updateTTLWrap struct {
		updateTTL *dynamo.UpdateTTL
	}
)

// Run : Executes this request.
func (utw *updateTTLWrap) Run() error {
	return utw.updateTTL.Run()
}

// RunWithContext : Executes this request.
func (utw *updateTTLWrap) RunWithContext(ctx aws.Context) error {
	return utw.updateTTL.RunWithContext(ctx)
}
