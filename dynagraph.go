package dynagraph

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

type Config struct {
	TableName string
}

func New(sess *session.Session, config Config) Graph {
	db := dynamo.New(sess)
	table := db.Table(config.TableName)
	return graph{
		db:    db,
		table: table,
	}
}

type graph struct {
	db    *dynamo.DB
	table dynamo.Table
}

// If the Node id is not set when calling the Node() method, then a random id will be configured on it.
// You can check this value by accessing the Node Id() method.
func (g graph) Node(id ...string) Node {
	n := node{
		table: g.table,
	}
	switch len(id) {
	case 0:
	case 1:
		n.id = id[0]
	default:
		n.err = fmt.Errorf("dynagraph: too many keys")
	}

	return n
}

func (g graph) Query(filter Filter) Query {
	return query{
		table:  g.table,
		filter: filter,
	}
}

func (g graph) Batch() Batch {
	return batch{table: g.table}
}

func (g graph) Transaction() Transaction {
	return &transaction{db: g.db, table: g.table, wtx: g.db.WriteTx()}
}
