package utils

import (
	"github.com/maxidelgado/dynagraph/internal/common"

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

type Operations []interface{}

func (i Operations) AppendNode(id string, value interface{}) (Operations, error) {
	nodeType := common.Type(value)
	if id == "" {
		id = nodeType + "-" + uuid.New().String()
	}

	edgeMap, err := common.Put(id, nodeType, value)
	if err != nil {
		return nil, err
	}
	return append(i, edgeMap), nil
}

func (i Operations) AppendEdge(id string, value interface{}) (Operations, error) {
	edgeMap, err := common.Edge(id, value)
	if err != nil {
		return nil, err
	}
	return append(i, edgeMap), nil
}

func (i Operations) AppendProp(id string, value interface{}) (Operations, error) {
	propMap, err := common.Prop(id, value)
	if err != nil {
		return nil, err
	}
	return append(i, propMap), nil
}

func (i Operations) AppendRef(sourceId, targetId string) (Operations, error) {
	refMap, err := common.Ref(sourceId, targetId)
	if err != nil {
		return nil, err
	}
	return append(i, refMap), nil
}

type IDs []dynamo.Keyed

func (i IDs) AppendKeys(keys dynamo.Keyed) IDs {
	return append(i, keys)
}

type ID struct {
	Id   string
	Type string
}

func (i ID) HashKey() interface{}  { return i.Id }
func (i ID) RangeKey() interface{} { return i.Type }

type Query struct {
	ID
	Operator Operator
	Index    Index
}

func (q Query) OperatorKey() Operator {
	if len(q.Operator) == 0 {
		return BeginsWith
	}
	return q.Operator
}
func (q Query) IndexKey() Index {
	if len(q.Index) == 0 {
		return Default
	}
	return q.Index
}

func (q Query) GetHashSchema() (string, string) {
	if q.Index == ByType {
		return common.NodeType, q.Type
	}
	return common.NodeId, q.Id
}

func (q Query) GetRangeSchema() (string, dynamo.Operator, string) {
	if q.Index == ByType {
		return common.NodeId, dynamo.Operator(q.OperatorKey()), q.Id
	}
	return common.NodeType, dynamo.Operator(q.OperatorKey()), q.Type
}
