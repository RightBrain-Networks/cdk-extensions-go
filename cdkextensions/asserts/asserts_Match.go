package asserts

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/assertions"
	"github.com/vibe-io/cdk-extensions-go/cdkextensions/asserts/internal"
)

type Match interface {
	assertions.Match
}

// The jsii proxy struct for Match
type jsiiProxy_Match struct {
	internal.Type__assertionsMatch
}

func NewMatch() Match {
	_init_.Initialize()

	j := jsiiProxy_Match{}

	_jsii_.Create(
		"cdk-extensions.asserts.Match",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewMatch_Override(m Match) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.asserts.Match",
		nil, // no parameters
		m,
	)
}

// Use this matcher in the place of a field's value, if the field must not be present.
func Match_Absent() assertions.Matcher {
	_init_.Initialize()

	var returns assertions.Matcher

	_jsii_.StaticInvoke(
		"cdk-extensions.asserts.Match",
		"absent",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Matches any non-null value at the target.
func Match_AnyValue() assertions.Matcher {
	_init_.Initialize()

	var returns assertions.Matcher

	_jsii_.StaticInvoke(
		"cdk-extensions.asserts.Match",
		"anyValue",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Matches the specified pattern with the array found in the same relative path of the target.
//
// The set of elements (or matchers) must match exactly and in order.
func Match_ArrayEquals(pattern *[]interface{}) assertions.Matcher {
	_init_.Initialize()

	if err := validateMatch_ArrayEqualsParameters(pattern); err != nil {
		panic(err)
	}
	var returns assertions.Matcher

	_jsii_.StaticInvoke(
		"cdk-extensions.asserts.Match",
		"arrayEquals",
		[]interface{}{pattern},
		&returns,
	)

	return returns
}

// Matches the specified pattern with the array found in the same relative path of the target.
//
// The set of elements (or matchers) must be in the same order as would be found.
func Match_ArrayWith(pattern *[]interface{}) assertions.Matcher {
	_init_.Initialize()

	if err := validateMatch_ArrayWithParameters(pattern); err != nil {
		panic(err)
	}
	var returns assertions.Matcher

	_jsii_.StaticInvoke(
		"cdk-extensions.asserts.Match",
		"arrayWith",
		[]interface{}{pattern},
		&returns,
	)

	return returns
}

// Deep exact matching of the specified pattern to the target.
func Match_Exact(pattern interface{}) assertions.Matcher {
	_init_.Initialize()

	if err := validateMatch_ExactParameters(pattern); err != nil {
		panic(err)
	}
	var returns assertions.Matcher

	_jsii_.StaticInvoke(
		"cdk-extensions.asserts.Match",
		"exact",
		[]interface{}{pattern},
		&returns,
	)

	return returns
}

func Match_JoinedJson(pattern interface{}) assertions.Matcher {
	_init_.Initialize()

	if err := validateMatch_JoinedJsonParameters(pattern); err != nil {
		panic(err)
	}
	var returns assertions.Matcher

	_jsii_.StaticInvoke(
		"cdk-extensions.asserts.Match",
		"joinedJson",
		[]interface{}{pattern},
		&returns,
	)

	return returns
}

// Matches any target which does NOT follow the specified pattern.
func Match_Not(pattern interface{}) assertions.Matcher {
	_init_.Initialize()

	if err := validateMatch_NotParameters(pattern); err != nil {
		panic(err)
	}
	var returns assertions.Matcher

	_jsii_.StaticInvoke(
		"cdk-extensions.asserts.Match",
		"not",
		[]interface{}{pattern},
		&returns,
	)

	return returns
}

// Matches the specified pattern to an object found in the same relative path of the target.
//
// The keys and their values (or matchers) must match exactly with the target.
func Match_ObjectEquals(pattern *map[string]interface{}) assertions.Matcher {
	_init_.Initialize()

	if err := validateMatch_ObjectEqualsParameters(pattern); err != nil {
		panic(err)
	}
	var returns assertions.Matcher

	_jsii_.StaticInvoke(
		"cdk-extensions.asserts.Match",
		"objectEquals",
		[]interface{}{pattern},
		&returns,
	)

	return returns
}

// Matches the specified pattern to an object found in the same relative path of the target.
//
// The keys and their values (or matchers) must be present in the target but the target can be a superset.
func Match_ObjectLike(pattern *map[string]interface{}) assertions.Matcher {
	_init_.Initialize()

	if err := validateMatch_ObjectLikeParameters(pattern); err != nil {
		panic(err)
	}
	var returns assertions.Matcher

	_jsii_.StaticInvoke(
		"cdk-extensions.asserts.Match",
		"objectLike",
		[]interface{}{pattern},
		&returns,
	)

	return returns
}

// Matches any string-encoded JSON and applies the specified pattern after parsing it.
func Match_SerializedJson(pattern interface{}) assertions.Matcher {
	_init_.Initialize()

	if err := validateMatch_SerializedJsonParameters(pattern); err != nil {
		panic(err)
	}
	var returns assertions.Matcher

	_jsii_.StaticInvoke(
		"cdk-extensions.asserts.Match",
		"serializedJson",
		[]interface{}{pattern},
		&returns,
	)

	return returns
}

// Matches targets according to a regular expression.
func Match_StringLikeRegexp(pattern *string) assertions.Matcher {
	_init_.Initialize()

	if err := validateMatch_StringLikeRegexpParameters(pattern); err != nil {
		panic(err)
	}
	var returns assertions.Matcher

	_jsii_.StaticInvoke(
		"cdk-extensions.asserts.Match",
		"stringLikeRegexp",
		[]interface{}{pattern},
		&returns,
	)

	return returns
}

