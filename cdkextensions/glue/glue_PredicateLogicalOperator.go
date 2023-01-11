package glue


// Logical operator that specifies how the conditions of a predicate should be evaluated.
type PredicateLogicalOperator string

const (
	// State equals specified value.
	PredicateLogicalOperator_EQUALS PredicateLogicalOperator = "EQUALS"
)

