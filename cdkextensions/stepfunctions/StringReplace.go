package stepfunctions

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/vibe-io/cdk-extensions-go/cdkextensions/stepfunctions/internal"
)

type StringReplace interface {
	constructs.Construct
	awsstepfunctions.IChainable
	awsstepfunctions.INextable
	// The chainable end state(s) of this chainable.
	EndStates() *[]awsstepfunctions.INextable
	// Descriptive identifier for this chainable.
	Id() *string
	// The tree node.
	Node() constructs.Node
	// The start state of this chainable.
	StartState() awsstepfunctions.State
	// Go to the indicated state after this state.
	Next(state awsstepfunctions.IChainable) awsstepfunctions.Chain
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for StringReplace
type jsiiProxy_StringReplace struct {
	internal.Type__constructsConstruct
	internal.Type__awsstepfunctionsIChainable
	internal.Type__awsstepfunctionsINextable
}

func (j *jsiiProxy_StringReplace) EndStates() *[]awsstepfunctions.INextable {
	var returns *[]awsstepfunctions.INextable
	_jsii_.Get(
		j,
		"endStates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StringReplace) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StringReplace) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StringReplace) StartState() awsstepfunctions.State {
	var returns awsstepfunctions.State
	_jsii_.Get(
		j,
		"startState",
		&returns,
	)
	return returns
}


func NewStringReplace(scope constructs.IConstruct, id *string, props *StringReplaceProps) StringReplace {
	_init_.Initialize()

	if err := validateNewStringReplaceParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_StringReplace{}

	_jsii_.Create(
		"cdk-extensions.stepfunctions.StringReplace",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewStringReplace_Override(s StringReplace, scope constructs.IConstruct, id *string, props *StringReplaceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.stepfunctions.StringReplace",
		[]interface{}{scope, id, props},
		s,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func StringReplace_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateStringReplace_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-extensions.stepfunctions.StringReplace",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StringReplace) Next(state awsstepfunctions.IChainable) awsstepfunctions.Chain {
	if err := s.validateNextParameters(state); err != nil {
		panic(err)
	}
	var returns awsstepfunctions.Chain

	_jsii_.Invoke(
		s,
		"next",
		[]interface{}{state},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StringReplace) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

