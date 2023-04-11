package aps

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"
)

// An expression that alert manager can use to evaluate incoming alerts to determine the actions it should take.
type AlertManagerMatcher interface {
	// The PromQL like expression to use for the matcher.
	Expression() *string
}

// The jsii proxy struct for AlertManagerMatcher
type jsiiProxy_AlertManagerMatcher struct {
	_ byte // padding
}

func (j *jsiiProxy_AlertManagerMatcher) Expression() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expression",
		&returns,
	)
	return returns
}


// Builds a simple matcher using the standard components supported by the matcher expression syntax.
//
// Handles formatting an escaping of values.
//
// Returns: A matcher object representing the given operation.
func AlertManagerMatcher_FromComponents(label *string, operator MatchOperator, value *string) AlertManagerMatcher {
	_init_.Initialize()

	if err := validateAlertManagerMatcher_FromComponentsParameters(label, operator, value); err != nil {
		panic(err)
	}
	var returns AlertManagerMatcher

	_jsii_.StaticInvoke(
		"cdk-extensions.aps.AlertManagerMatcher",
		"fromComponents",
		[]interface{}{label, operator, value},
		&returns,
	)

	return returns
}

// Creates a compund matcher expression by joining one or more other matcher objects.
//
// Returns: A compound matcher representing the joined expressions.
func AlertManagerMatcher_FromCompound(matchers ...AlertManagerMatcher) AlertManagerMatcher {
	_init_.Initialize()

	args := []interface{}{}
	for _, a := range matchers {
		args = append(args, a)
	}

	var returns AlertManagerMatcher

	_jsii_.StaticInvoke(
		"cdk-extensions.aps.AlertManagerMatcher",
		"fromCompound",
		args,
		&returns,
	)

	return returns
}

// Creates a matcher from a raw string expression.
//
// This allows for specifying
// arbitrary matching conditions that may be too complex to be supported by
// the other means of constructing matchers.
//
// Returns: A matcher object representing the given expression.
// See: [Matcher expression syntax](https://prometheus.io/docs/alerting/latest/configuration/#matcher)
//
func AlertManagerMatcher_FromString(expression *string) AlertManagerMatcher {
	_init_.Initialize()

	if err := validateAlertManagerMatcher_FromStringParameters(expression); err != nil {
		panic(err)
	}
	var returns AlertManagerMatcher

	_jsii_.StaticInvoke(
		"cdk-extensions.aps.AlertManagerMatcher",
		"fromString",
		[]interface{}{expression},
		&returns,
	)

	return returns
}

