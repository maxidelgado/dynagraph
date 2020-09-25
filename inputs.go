package dynagraph

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/guregu/dynamo"
)

// Operator is an operation to apply in filter comparisons.
type Operator string

// Operators used for comparing against the range filter in queries.
const (
	Equal          Operator = "EQ"
	NotEqual       Operator = "NE"
	Less           Operator = "LT"
	LessOrEqual    Operator = "LE"
	Greater        Operator = "GT"
	GreaterOrEqual Operator = "GE"
	BeginsWith     Operator = "BEGINS_WITH"
	Between        Operator = "BETWEEN"
	Noop           Operator = "Noop"
)

type Index string

const (
	Default Index = "default"
	ByType  Index = "ByType"
)

type WriteItemsInput []interface{}

func (i WriteItemsInput) AppendNode(id string, value interface{}) (WriteItemsInput, error) {
	nodeType := getNodeType(value)
	if id == "" {
		id = fmt.Sprintf("%s-%s", nodeType, uuid.New().String())
	}

	edgeMap, err := buildPut(id, nodeType, value)
	if err != nil {
		return nil, err
	}
	return append(i, edgeMap), nil
}

func (i WriteItemsInput) AppendEdge(id string, value interface{}) (WriteItemsInput, error) {
	edgeMap, err := addEdge(id, value)
	if err != nil {
		return nil, err
	}
	return append(i, edgeMap), nil
}

func (i WriteItemsInput) AppendProp(id string, value interface{}) (WriteItemsInput, error) {
	propMap, err := addProp(id, value)
	if err != nil {
		return nil, err
	}
	return append(i, propMap), nil
}

func (i WriteItemsInput) AppendRef(sourceId, targetId string) (WriteItemsInput, error) {
	refMap, err := addRef(sourceId, targetId)
	if err != nil {
		return nil, err
	}
	return append(i, refMap), nil
}

type KeysInput []dynamo.Keyed

func (i KeysInput) AppendKeys(keys Filter) KeysInput {
	return append(i, keys)
}

type Filter struct {
	Id       string
	Type     string
	Operator Operator
	Index    Index
}

func (k Filter) HashKey() interface{}  { return k.Id }
func (k Filter) RangeKey() interface{} { return k.Type }
func (k Filter) OperatorKey() Operator {
	if len(k.Operator) == 0 {
		return BeginsWith
	}
	return k.Operator
}
func (k Filter) IndexKey() Index {
	if len(k.Index) == 0 {
		return Default
	}
	return k.Index
}

func (k Filter) GetHashValue() (string, string) {
	if k.Index == ByType {
		return nodeType, k.Type
	}
	return nodeId, k.Id
}
func (k Filter) GetRangeValues() (string, dynamo.Operator, string) {
	if k.Index == ByType {
		return nodeId, dynamo.Operator(k.OperatorKey()), k.Id
	}
	return nodeType, dynamo.Operator(k.OperatorKey()), k.Type
}
