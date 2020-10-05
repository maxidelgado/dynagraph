package dynamoiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/guregu/dynamo"
)

type (
	// DB : DynamoDB client.
	DB interface {
		// Client : Returns this DB's internal client used to make API requests.
		Client() dynamodbiface.DynamoDBAPI
		// ListTables : Begins a new request to list all tables.
		ListTables() ListTables
		// Table : Returns a Table handle specified by name.
		Table(name string) Table
		// CreateTable : Begins a new operation to create a table with the given name.
		CreateTable(name string, from interface{}) CreateTable
		// GetTx : Begins a new get transaction.
		GetTx() GetTx
		// WriteTx : Begins a new write transaction.
		WriteTx() WriteTx
	}

	dbWrap struct {
		db *dynamo.DB
	}
)

// New : Creates a new client with the given configuration.
func New(p client.ConfigProvider, cfgs ...*aws.Config) DB {
	return &dbWrap{
		db: dynamo.New(p, cfgs...),
	}
}

func NewFromIface(client dynamodbiface.DynamoDBAPI) DB {
	return &dbWrap{db: dynamo.NewFromIface(client)}
}

// Client : Returns this DB's internal client used to make API requests.
func (dbw *dbWrap) Client() dynamodbiface.DynamoDBAPI {
	return dbw.db.Client()
}

// ListTables : Begins a new request to list all tables.
func (dbw *dbWrap) ListTables() ListTables {
	return &listTablesWrap{
		listTables: dbw.db.ListTables(),
	}
}

// Table : Returns a Table handle specified by name.
func (dbw *dbWrap) Table(name string) Table {
	table := dbw.db.Table(name)

	return &tableWrap{
		table: &table,
	}
}

// CreateTable : Begins a new operation to create a table with the given name.
func (dbw *dbWrap) CreateTable(name string, from interface{}) CreateTable {
	return &createTableWrap{
		createTable: dbw.db.CreateTable(name, from),
	}
}

// GetTx : Begins a new get transaction.
func (dbw *dbWrap) GetTx() GetTx {
	return &getTxWrap{
		getTx: dbw.db.GetTx(),
	}
}

// WriteTx : Begins a new write transaction.
func (dbw *dbWrap) WriteTx() WriteTx {
	return &writeTxWrap{
		writeTx: dbw.db.WriteTx(),
	}
}
