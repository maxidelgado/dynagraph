package common

import (
	"errors"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/fatih/structs"
	"github.com/google/uuid"
	"github.com/guregu/dynamo"
)

const (
	NodeId    = "Id"
	NodeType  = "Type"
	NodeName  = "node"
	EdgeName  = "edge"
	PropName  = "prop"
	RefName   = "ref"
	Separator = ":"
)

func Put(id, ntype string, value interface{}) (map[string]*dynamodb.AttributeValue, error) {
	if id == "" {
		return nil, errors.New("id is mandatory")
	}

	nodeMap, err := dynamo.MarshalItem(value)
	if err != nil {
		return nil, err
	}

	nodeMap[NodeId] = &dynamodb.AttributeValue{S: aws.String(id)}
	nodeMap[NodeType] = &dynamodb.AttributeValue{S: aws.String(NodeName + Separator + ntype)}

	return nodeMap, nil
}

func Update(value interface{}) (map[string]*dynamodb.AttributeValue, error) {
	avs, err := dynamo.MarshalItem(value)
	if err != nil {
		return nil, err
	}

	return avs, nil
}

func Edge(id string, value interface{}) (map[string]*dynamodb.AttributeValue, error) {
	if id == "" {
		return nil, errors.New("source node id is mandatory")
	}

	edgeMap, err := dynamo.MarshalItem(value)
	if err != nil {
		return nil, err
	}
	targetUID, ok := edgeMap[NodeId]
	if !ok || *targetUID.S == "" {
		return nil, errors.New("target's id is mandatory")
	}

	edgeMap[NodeType] = &dynamodb.AttributeValue{S: aws.String(EdgeName + Separator + *targetUID.S)}
	edgeMap[NodeId] = &dynamodb.AttributeValue{S: aws.String(id)}

	return edgeMap, nil
}

func Prop(id string, value interface{}) (map[string]*dynamodb.AttributeValue, error) {
	if id == "" {
		return nil, errors.New("source node id is mandatory")
	}

	propMap, err := dynamo.MarshalItem(value)
	if err != nil {
		return nil, err
	}

	pName := Type(value)

	propMap[NodeId] = &dynamodb.AttributeValue{S: aws.String(id)}
	propMap[NodeType] = &dynamodb.AttributeValue{S: aws.String(PropName + Separator + pName)}

	return propMap, nil
}

func Ref(sourceId, targetId string) (map[string]*dynamodb.AttributeValue, error) {
	if sourceId == "" || targetId == "" {
		return nil, errors.New("sourceId and targetId are mandatory")
	}

	return map[string]*dynamodb.AttributeValue{
		NodeId:   {S: aws.String(sourceId)},
		NodeType: {S: aws.String(RefName + Separator + targetId)},
	}, nil
}

func Type(value interface{}) string {
	return strings.ToLower(structs.Name(value))
}

func Id(value interface{}) string {
	nodeType := Type(value)
	return fmt.Sprintf("%s:%s", nodeType, uuid.New().String())
}
