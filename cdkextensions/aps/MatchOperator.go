package aps


// The logical operator an alert manager matcher should use when evaluating filters for labels.
type MatchOperator string

const (
	// Evaluate an alert manager filter on the basis that the label matches the string it is being compared against.
	MatchOperator_EQUALS MatchOperator = "EQUALS"
	// Evaluate an alert manager filter on the basis that the label does not match the string it is being compared against.
	MatchOperator_NOT_EQUALS MatchOperator = "NOT_EQUALS"
	// Evaluate an alert manager filter on the basis that the label matches the regular expression it is being compared against.
	MatchOperator_RE_EQUALS MatchOperator = "RE_EQUALS"
	// Evaluate an alert manager filter on the basis that the label does not match the regular expression it is being compared against.
	MatchOperator_RE_NOT_EQUALS MatchOperator = "RE_NOT_EQUALS"
)

