package dynamoiface

import "github.com/guregu/dynamo"

type (
	// ConditionCheck : Represents a condition for a write transaction to succeed.
	ConditionCheck interface {
		// Range : Specifies the name and value of the range key for this item.
		Range(rangeKey string, value interface{}) ConditionCheck
		// If : Specifies a conditional expression for this coniditon check to succeed.
		If(expr string, args ...interface{}) ConditionCheck
		// IfExists : Sets this check to succeed if the item exists.
		IfExists() ConditionCheck
		// IfNotExists : Sets this check to succeed if the item does not exist.
		IfNotExists() ConditionCheck
	}

	conditionCheckWrap struct {
		conditionCheck *dynamo.ConditionCheck
	}
)

// Range : Specifies the name and value of the range key for this item.
func (ccw *conditionCheckWrap) Range(rangeKey string, value interface{}) ConditionCheck {
	return &conditionCheckWrap{
		conditionCheck: ccw.conditionCheck.Range(rangeKey, value),
	}
}

// If : Specifies a conditional expression for this coniditon check to succeed.
func (ccw *conditionCheckWrap) If(expr string, args ...interface{}) ConditionCheck {
	return &conditionCheckWrap{
		conditionCheck: ccw.conditionCheck.If(expr, args...),
	}
}

// IfExists : Sets this check to succeed if the item exists.
func (ccw *conditionCheckWrap) IfExists() ConditionCheck {
	return &conditionCheckWrap{
		conditionCheck: ccw.conditionCheck.IfExists(),
	}
}

// IfNotExists : Sets this check to succeed if the item does not exist.
func (ccw *conditionCheckWrap) IfNotExists() ConditionCheck {
	return &conditionCheckWrap{
		conditionCheck: ccw.conditionCheck.IfNotExists(),
	}
}
