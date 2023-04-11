package k8saws

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"
)

// Represents a filter that can be applied to Filter and Output plugins that scopes down what records the given filter should apply to.
type FluentBitMatch interface {
	// The pattern matching syntax to use when evaluating incoming tags.
	Evaluator() FluentBitMatchEvaluator
	// The pattern to compare against the tags of incoming records.
	Pattern() *string
	// Creates a record object that can be used to represent the match in Fluent Bit configuration files.
	//
	// Returns: The object that can be used to represent this match object.
	ToObject() *map[string]*string
	// Creates a string representation of this match object that reflects how it will appear in a Fluent Bit configuration file.
	//
	// Returns: A string representation of this match.
	ToString() *string
}

// The jsii proxy struct for FluentBitMatch
type jsiiProxy_FluentBitMatch struct {
	_ byte // padding
}

func (j *jsiiProxy_FluentBitMatch) Evaluator() FluentBitMatchEvaluator {
	var returns FluentBitMatchEvaluator
	_jsii_.Get(
		j,
		"evaluator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitMatch) Pattern() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pattern",
		&returns,
	)
	return returns
}


// Creates a match pattern that supports basic wildcard matching using the star character (`*`).
//
// Returns: A match object representing the given pattern.
func FluentBitMatch_Glob(pattern *string) FluentBitMatch {
	_init_.Initialize()

	if err := validateFluentBitMatch_GlobParameters(pattern); err != nil {
		panic(err)
	}
	var returns FluentBitMatch

	_jsii_.StaticInvoke(
		"cdk-extensions.k8s_aws.FluentBitMatch",
		"glob",
		[]interface{}{pattern},
		&returns,
	)

	return returns
}

// Creates a match pattern that supports full regex matching.
//
// Returns: A match object representing the given pattern.
func FluentBitMatch_Regex(pattern *string) FluentBitMatch {
	_init_.Initialize()

	if err := validateFluentBitMatch_RegexParameters(pattern); err != nil {
		panic(err)
	}
	var returns FluentBitMatch

	_jsii_.StaticInvoke(
		"cdk-extensions.k8s_aws.FluentBitMatch",
		"regex",
		[]interface{}{pattern},
		&returns,
	)

	return returns
}

func FluentBitMatch_ALL() FluentBitMatch {
	_init_.Initialize()
	var returns FluentBitMatch
	_jsii_.StaticGet(
		"cdk-extensions.k8s_aws.FluentBitMatch",
		"ALL",
		&returns,
	)
	return returns
}

func (f *jsiiProxy_FluentBitMatch) ToObject() *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		f,
		"toObject",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FluentBitMatch) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

