package glue

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
)

type WorkflowActionBase interface {
	NotifyDelayAfter() awscdk.Duration
	SecurityConfiguration() *string
	Timeout() awscdk.Duration
	AddArgument(key *string, value *string)
	BindOptions(_scope constructs.IConstruct) interface{}
}

// The jsii proxy struct for WorkflowActionBase
type jsiiProxy_WorkflowActionBase struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowActionBase) NotifyDelayAfter() awscdk.Duration {
	var returns awscdk.Duration
	_jsii_.Get(
		j,
		"notifyDelayAfter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkflowActionBase) SecurityConfiguration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkflowActionBase) Timeout() awscdk.Duration {
	var returns awscdk.Duration
	_jsii_.Get(
		j,
		"timeout",
		&returns,
	)
	return returns
}


func NewWorkflowActionBase(options *WorkflowActionOptions) WorkflowActionBase {
	_init_.Initialize()

	if err := validateNewWorkflowActionBaseParameters(options); err != nil {
		panic(err)
	}
	j := jsiiProxy_WorkflowActionBase{}

	_jsii_.Create(
		"cdk-extensions.glue.WorkflowActionBase",
		[]interface{}{options},
		&j,
	)

	return &j
}

func NewWorkflowActionBase_Override(w WorkflowActionBase, options *WorkflowActionOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.glue.WorkflowActionBase",
		[]interface{}{options},
		w,
	)
}

func (w *jsiiProxy_WorkflowActionBase) AddArgument(key *string, value *string) {
	if err := w.validateAddArgumentParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"addArgument",
		[]interface{}{key, value},
	)
}

func (w *jsiiProxy_WorkflowActionBase) BindOptions(_scope constructs.IConstruct) interface{} {
	if err := w.validateBindOptionsParameters(_scope); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		w,
		"bindOptions",
		[]interface{}{_scope},
		&returns,
	)

	return returns
}

