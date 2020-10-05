package dynamoiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/guregu/dynamo"
)

type (
	// Scan : Request to scan all the data in a table.
	Scan interface {
		// StartFrom : Makes this scan continue from a previous one.
		StartFrom(key dynamo.PagingKey) Scan
		// Index : Specifies the name of the index that Scan will operate on.
		Index(name string) Scan
		// Project : Limits the result attributes to the given paths.
		Project(paths ...string) Scan
		// Filter : Takes an expression that all results will be evaluated against.
		Filter(expr string, args ...interface{}) Scan
		// Consistent : Set the read consistency to strong or not.
		Consistent(on bool) Scan
		// Limit : Specifies the maximum amount of results to return.
		Limit(limit int64) Scan
		// SearchLimit : Specifies a maximum amount of results to evaluate.
		SearchLimit(limit int64) Scan
		// ConsumedCapacity : Measure the throughput capacity consumed by this operation and add it to cc.
		ConsumedCapacity(cc *dynamo.ConsumedCapacity) Scan
		// Iter : Returns a results iterator for this request.
		Iter() dynamo.PagingIter
		// All : Executes this request and unmarshals all results to out, which must be a pointer to a slice.
		All(out interface{}) error
		// AllWithContext : Executes this request and unmarshals all results to out, which must be a pointer to a slice.
		AllWithContext(ctx aws.Context, out interface{}) error
		// AllWithLastEvaluatedKey : Executes this request and unmarshals all results to out, which must be a pointer to a slice.
		AllWithLastEvaluatedKey(out interface{}) (dynamo.PagingKey, error)
		// AllWithLastEvaluatedKeyContext : Executes this request and unmarshals all results to out, which must be a pointer to a slice.
		AllWithLastEvaluatedKeyContext(ctx aws.Context, out interface{}) (dynamo.PagingKey, error)
		// Count : Executes this request and returns the number of items matching the scan.
		Count() (int64, error)
		// CountWithContext : Executes this request and returns the number of items matching the scan.
		CountWithContext(ctx aws.Context) (int64, error)
	}

	scanWrap struct {
		scan *dynamo.Scan
	}
)

// StartFrom : Makes this scan continue from a previous one.
func (sw *scanWrap) StartFrom(key dynamo.PagingKey) Scan {
	return &scanWrap{
		scan: sw.scan.StartFrom(key),
	}
}

// Index : Specifies the name of the index that Scan will operate on.
func (sw *scanWrap) Index(name string) Scan {
	return &scanWrap{
		scan: sw.scan.Index(name),
	}
}

// Project : Limits the result attributes to the given paths.
func (sw *scanWrap) Project(paths ...string) Scan {
	return &scanWrap{
		scan: sw.scan.Project(paths...),
	}
}

// Filter : Takes an expression that all results will be evaluated against.
func (sw *scanWrap) Filter(expr string, args ...interface{}) Scan {
	return &scanWrap{
		scan: sw.scan.Filter(expr, args...),
	}
}

// Consistent : Set the read consistency to strong or not.
func (sw *scanWrap) Consistent(on bool) Scan {
	return &scanWrap{
		scan: sw.scan.Consistent(on),
	}
}

// Limit : Specifies the maximum amount of results to return.
func (sw *scanWrap) Limit(limit int64) Scan {
	return &scanWrap{
		scan: sw.scan.Limit(limit),
	}
}

// SearchLimit : Specifies a maximum amount of results to evaluate.
func (sw *scanWrap) SearchLimit(limit int64) Scan {
	return &scanWrap{
		scan: sw.scan.SearchLimit(limit),
	}
}

// ConsumedCapacity : Measure the throughput capacity consumed by this operation and add it to cc.
func (sw *scanWrap) ConsumedCapacity(cc *dynamo.ConsumedCapacity) Scan {
	return &scanWrap{
		scan: sw.scan.ConsumedCapacity(cc),
	}
}

// Iter : Returns a results iterator for this request.
func (sw *scanWrap) Iter() dynamo.PagingIter {
	return sw.scan.Iter()
}

// All : Executes this request and unmarshals all results to out, which must be a pointer to a slice.
func (sw *scanWrap) All(out interface{}) error {
	return sw.scan.All(out)
}

// AllWithContext : Executes this request and unmarshals all results to out, which must be a pointer to a slice.
func (sw *scanWrap) AllWithContext(ctx aws.Context, out interface{}) error {
	return sw.scan.AllWithContext(ctx, out)
}

// AllWithLastEvaluatedKey : Executes this request and unmarshals all results to out, which must be a pointer to a slice.
func (sw *scanWrap) AllWithLastEvaluatedKey(out interface{}) (dynamo.PagingKey, error) {
	return sw.scan.AllWithLastEvaluatedKey(out)
}

// AllWithLastEvaluatedKeyContext : Executes this request and unmarshals all results to out, which must be a pointer to a slice.
func (sw *scanWrap) AllWithLastEvaluatedKeyContext(ctx aws.Context, out interface{}) (dynamo.PagingKey, error) {
	return sw.scan.AllWithLastEvaluatedKeyContext(ctx, out)
}

// Count : Executes this request and returns the number of items matching the scan.
func (sw *scanWrap) Count() (int64, error) {
	return sw.scan.Count()
}

// CountWithContext : Executes this request and returns the number of items matching the scan.
func (sw *scanWrap) CountWithContext(ctx aws.Context) (int64, error) {
	return sw.scan.CountWithContext(ctx)
}
