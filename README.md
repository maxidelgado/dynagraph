# DynaGraph

_This is an experimental DynamoDB helper that allows you to create relations between data and query it in a simple way.
This is based on this article from AWS:_

[https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/bp-adjacency-graphs.html][]

## Starting 🚀

_The only thing you need to proceed is a DynamoDB table with the following schema:_

```
aws dynamodb create-table \
    --table-name YOUR_TABLE_NAME \
    --attribute-definitions \
        AttributeName=Id,AttributeType=S \
        AttributeName=Type,AttributeType=S \
    --key-schema \
        AttributeName=Id,KeyType=HASH \
        AttributeName=Type,KeyType=RANGE 
```
        
_And you need to create a GSI as follows:_
```
aws dynamodb update-table \
     --table-name YOUR_TABLE_NAME \
     --global-secondary-index-updates \
     "[{\"Create\":{\"IndexName\": \"ByType\",\"KeySchema\":[{\"AttributeName\":\"Type\",\"KeyType\":\"HASH\"}, {\"AttributeName\":\"Id\",\"KeyType\":\"RANGE\"}], \
    \"Projection\":{\"ProjectionType\":\"ALL\"}}}]"
```
### Install

_Just need to import this package into your project and vendor_

```
import "github.com/maxidelgado/dynagraph"

go mod vendor
```

### How to use 🔧

_Look into the examples folder_