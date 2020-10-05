package dynamoiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/guregu/dynamo"
)

type (
	// Query : Request to get one or more items in a table.
	Query interface {
		// Range : Specifies the range key (a.k.a. sort key) or keys to get.
		Range(name string, op dynamo.Operator, values ...interface{}) Query
		// StartFrom : Makes this query continue from a previous one.
		StartFrom(key dynamo.PagingKey) Query
		// Index : Specifies the name of the index that this query will operate on.
		Index(name string) Query
		// Project : Limits the result attributes to the given paths.
		Project(paths ...string) Query
		// ProjectExpr : Limits the result attributes to the given expression.
		ProjectExpr(expr string, args ...interface{}) Query
		// Filter : Takes an expression that all results will be evaluated against.
		Filter(expr string, args ...interface{}) Query
		// Consistent : Set the read consistency to strong or not.
		Consistent(on bool) Query
		// Limit : Specifies the maximum amount of results to return.
		Limit(limit int64) Query
		// SearchLimit : Specifies the maximum amount of results to examine.
		SearchLimit(limit int64) Query
		// Order : Specifies the desired result order.
		Order(order dynamo.Order) Query
		// ConsumedCapacity : Measures the throughput capacity consumed by this operation and add it to cc.
		ConsumedCapacity(cc *dynamo.ConsumedCapacity) Query
		// One : Executes this query and retrieves a single result.
		One(out interface{}) error
		// OneWithContext : Executes this query and retrieves a single result.
		OneWithContext(ctx aws.Context, out interface{}) error
		// Count : Executes this request, returning the number of results.
		Count() (int64, error)
		// CountWithContext : Executes this request, returning the number of results.
		CountWithContext(ctx aws.Context) (int64, error)
		// All : Executes this request and unmarshals all results to out, which must be a pointer to a slice.
		All(out interface{}) error
		// AllWithContext : Executes this request and unmarshals all results to out, which must be a pointer to a slice.
		AllWithContext(ctx aws.Context, out interface{}) error
		// AllWithLastEvaluatedKey : Executes this request and unmarshals all results to out, which must be a pointer to a slice.
		AllWithLastEvaluatedKey(out interface{}) (dynamo.PagingKey, error)
		// AllWithLastEvaluatedKeyContext : Executes this request and unmarshals all results to out, which must be a pointer to a slice.
		AllWithLastEvaluatedKeyContext(ctx aws.Context, out interface{}) (dynamo.PagingKey, error)
		// Iter : Returns a results iterator for this request.
		Iter() dynamo.PagingIter
	}

	queryWrap struct {
		query *dynamo.Query
	}
)

// Range : Specifies the range key (a.k.a. sort key) or keys to get.
func (qw *queryWrap) Range(name string, op dynamo.Operator, values ...interface{}) Query {
	return &queryWrap{
		query: qw.query.Range(name, op, values...),
	}
}

// StartFrom : Makes this query continue from a previous one.
func (qw *queryWrap) StartFrom(key dynamo.PagingKey) Query {
	return &queryWrap{
		query: qw.query.StartFrom(key),
	}
}

// Index : Specifies the name of the index that this query will operate on.
func (qw *queryWrap) Index(name string) Query {
	return &queryWrap{
		query: qw.query.Index(name),
	}
}

// Project : Limits the result attributes to the given paths.
func (qw *queryWrap) Project(paths ...string) Query {
	return &queryWrap{
		query: qw.query.Project(paths...),
	}
}

// ProjectExpr : Limits the result attributes to the given expression.
func (qw *queryWrap) ProjectExpr(expr string, args ...interface{}) Query {
	return &queryWrap{
		query: qw.query.ProjectExpr(expr, args...),
	}
}

// Filter : Takes an expression that all results will be evaluated against.
func (qw *queryWrap) Filter(expr string, args ...interface{}) Query {
	return &queryWrap{
		query: qw.query.Filter(expr, args...),
	}
}

// Consistent : Set the read consistency to strong or not.
func (qw *queryWrap) Consistent(on bool) Query {
	return &queryWrap{
		query: qw.query.Consistent(on),
	}
}

// Limit : Specifies the maximum amount of results to return.
func (qw *queryWrap) Limit(limit int64) Query {
	return &queryWrap{
		query: qw.query.Limit(limit),
	}
}

// SearchLimit : Specifies the maximum amount of results to examine.
func (qw *queryWrap) SearchLimit(limit int64) Query {
	return &queryWrap{
		query: qw.query.SearchLimit(limit),
	}
}

// Order : Specifies the desired result order.
func (qw *queryWrap) Order(order dynamo.Order) Query {
	return &queryWrap{
		query: qw.query.Order(order),
	}
}

// ConsumedCapacity : Measures the throughput capacity consumed by this operation and add it to cc.
func (qw *queryWrap) ConsumedCapacity(cc *dynamo.ConsumedCapacity) Query {
	return &queryWrap{
		query: qw.query.ConsumedCapacity(cc),
	}
}

// One : Executes this query and retrieves a single result.
func (qw *queryWrap) One(out interface{}) error {
	return qw.query.One(out)
}

// OneWithContext : Executes this query and retrieves a single result.
func (qw *queryWrap) OneWithContext(ctx aws.Context, out interface{}) error {
	return qw.query.OneWithContext(ctx, out)
}

// Count : Executes this request, returning the number of results.
func (qw *queryWrap) Count() (int64, error) {
	return qw.query.Count()
}

// CountWithContext : Executes this request, returning the number of results.
func (qw *queryWrap) CountWithContext(ctx aws.Context) (int64, error) {
	return qw.query.CountWithContext(ctx)
}

// All : Executes this request and unmarshals all results to out, which must be a pointer to a slice.
func (qw *queryWrap) All(out interface{}) error {
	return qw.query.All(out)
}

// AllWithContext : Executes this request and unmarshals all results to out, which must be a pointer to a slice.
func (qw *queryWrap) AllWithContext(ctx aws.Context, out interface{}) error {
	return qw.query.AllWithContext(ctx, out)
}

// AllWithLastEvaluatedKey : Executes this request and unmarshals all results to out, which must be a pointer to a slice.
func (qw *queryWrap) AllWithLastEvaluatedKey(out interface{}) (dynamo.PagingKey, error) {
	return qw.query.AllWithLastEvaluatedKey(out)
}

// AllWithLastEvaluatedKeyContext : Executes this request and unmarshals all results to out, which must be a pointer to a slice.
func (qw *queryWrap) AllWithLastEvaluatedKeyContext(ctx aws.Context, out interface{}) (dynamo.PagingKey, error) {
	return qw.query.AllWithLastEvaluatedKeyContext(ctx, out)
}

// Iter : Returns a results iterator for this request.
func (qw *queryWrap) Iter() dynamo.PagingIter {
	return qw.query.Iter()
}
