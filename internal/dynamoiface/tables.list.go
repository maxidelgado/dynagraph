package dynamoiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/guregu/dynamo"
)

type (
	// ListTables : Request to list tables.
	ListTables interface {
		// All : Returns every table or an error.
		All() ([]string, error)
		// AllWithContext : Returns every table or an error.
		AllWithContext(ctx aws.Context) ([]string, error)
		// Iter : Returns an iterator of table names.
		Iter() dynamo.Iter
	}

	listTablesWrap struct {
		listTables *dynamo.ListTables
	}
)

// All : Returns every table or an error.
func (ltw *listTablesWrap) All() ([]string, error) {
	return ltw.listTables.All()
}

// AllWithContext : Returns every table or an error.
func (ltw *listTablesWrap) AllWithContext(ctx aws.Context) ([]string, error) {
	return ltw.listTables.AllWithContext(ctx)
}

// Iter : Returns an iterator of table names.
func (ltw *listTablesWrap) Iter() dynamo.Iter {
	return ltw.listTables.Iter()
}
