package k8saws


// Matching patterns supported by Fluent Bit plugins for scoping down incoming records.
type FluentBitMatchEvaluator string

const (
	// A basic pattern match supporting the star (`*`) character as a wildcard.
	FluentBitMatchEvaluator_GLOB FluentBitMatchEvaluator = "GLOB"
	// Full pattern matching using regular expressions.
	FluentBitMatchEvaluator_REGEX FluentBitMatchEvaluator = "REGEX"
)

