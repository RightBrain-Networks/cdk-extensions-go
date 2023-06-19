package alerting

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/vibe-io/cdk-extensions-go/cdkextensions/alerting/internal"
)

type IssueTrigger interface {
	constructs.Construct
	EventPattern() *awsevents.EventPattern
	// The tree node.
	Node() constructs.Node
	Overrides() *[]IssueHandlerOverride
	Parser() IIssueParser
	AddOverride(handlerOverrides IssueHandlerOverride)
	Bind(stateMachine awsstepfunctions.StateMachine) awsevents.Rule
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for IssueTrigger
type jsiiProxy_IssueTrigger struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_IssueTrigger) EventPattern() *awsevents.EventPattern {
	var returns *awsevents.EventPattern
	_jsii_.Get(
		j,
		"eventPattern",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IssueTrigger) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IssueTrigger) Overrides() *[]IssueHandlerOverride {
	var returns *[]IssueHandlerOverride
	_jsii_.Get(
		j,
		"overrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IssueTrigger) Parser() IIssueParser {
	var returns IIssueParser
	_jsii_.Get(
		j,
		"parser",
		&returns,
	)
	return returns
}


func NewIssueTrigger(scope constructs.IConstruct, id *string, props *IssueTriggerProps) IssueTrigger {
	_init_.Initialize()

	if err := validateNewIssueTriggerParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_IssueTrigger{}

	_jsii_.Create(
		"cdk-extensions.alerting.IssueTrigger",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewIssueTrigger_Override(i IssueTrigger, scope constructs.IConstruct, id *string, props *IssueTriggerProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.alerting.IssueTrigger",
		[]interface{}{scope, id, props},
		i,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func IssueTrigger_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIssueTrigger_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-extensions.alerting.IssueTrigger",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IssueTrigger) AddOverride(handlerOverrides IssueHandlerOverride) {
	if err := i.validateAddOverrideParameters(handlerOverrides); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addOverride",
		[]interface{}{handlerOverrides},
	)
}

func (i *jsiiProxy_IssueTrigger) Bind(stateMachine awsstepfunctions.StateMachine) awsevents.Rule {
	if err := i.validateBindParameters(stateMachine); err != nil {
		panic(err)
	}
	var returns awsevents.Rule

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{stateMachine},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IssueTrigger) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

