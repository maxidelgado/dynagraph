package dynagraph

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
	nodeId    = "Id"
	nodeType  = "Type"
	nodeName  = "node"
	edgeName  = "edge"
	propName  = "prop"
	refName   = "ref"
	separator = ":"
)

func buildPut(id, ntype string, value interface{}) (map[string]*dynamodb.AttributeValue, error) {
	if id == "" {
		return nil, errors.New("id is mandatory")
	}

	nodeMap, err := dynamo.MarshalItem(value)
	if err != nil {
		return nil, err
	}

	nodeMap[nodeId] = &dynamodb.AttributeValue{S: aws.String(id)}
	nodeMap[nodeType] = &dynamodb.AttributeValue{S: aws.String(nodeName + separator + ntype)}

	return nodeMap, nil
}

func buildUpdate(value interface{}) (map[string]*dynamodb.AttributeValue, error) {
	avs, err := dynamo.MarshalItem(value)
	if err != nil {
		return nil, err
	}

	return avs, nil
}

func addEdge(id string, value interface{}) (map[string]*dynamodb.AttributeValue, error) {
	if id == "" {
		return nil, errors.New("source node id is mandatory")
	}

	edgeMap, err := dynamo.MarshalItem(value)
	if err != nil {
		return nil, err
	}
	targetUID, ok := edgeMap[nodeId]
	if !ok || *targetUID.S == "" {
		return nil, errors.New("target's id is mandatory")
	}

	edgeMap[nodeType] = &dynamodb.AttributeValue{S: aws.String(edgeName + separator + *targetUID.S)}
	edgeMap[nodeId] = &dynamodb.AttributeValue{S: aws.String(id)}

	return edgeMap, nil
}

func addProp(id string, value interface{}) (map[string]*dynamodb.AttributeValue, error) {
	if id == "" {
		return nil, errors.New("source node id is mandatory")
	}

	propMap, err := dynamo.MarshalItem(value)
	if err != nil {
		return nil, err
	}

	pName := getNodeType(value)

	propMap[nodeId] = &dynamodb.AttributeValue{S: aws.String(id)}
	propMap[nodeType] = &dynamodb.AttributeValue{S: aws.String(propName + separator + pName)}

	return propMap, nil
}

func addRef(sourceId, targetId string) (map[string]*dynamodb.AttributeValue, error) {
	if sourceId == "" || targetId == "" {
		return nil, errors.New("sourceId and targetId are mandatory")
	}

	return map[string]*dynamodb.AttributeValue{
		nodeId:   {S: aws.String(sourceId)},
		nodeType: {S: aws.String(refName + separator + targetId)},
	}, nil
}

func getNodeType(value interface{}) string {
	return strings.ToLower(structs.Name(value))
}

func NewId(value interface{}) string {
	nodeType := getNodeType(value)
	return fmt.Sprintf("%s:%s", nodeType, uuid.New().String())
}
