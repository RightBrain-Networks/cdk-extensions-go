package glue

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsglue"
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents the configuration for a job that will be triggered as part of a workflow.
type WorkflowJobAction interface {
	WorkflowActionBase
	ITriggerAction
	// The Glue job to be triggered as part of the workflow.
	Job() IJob
	NotifyDelayAfter() awscdk.Duration
	SecurityConfiguration() *string
	Timeout() awscdk.Duration
	AddArgument(key *string, value *string)
	// Associates this action with a resource that is configuring a Glue trigger.
	//
	// Returns: The configuration that can be used to configure the underlying
	// trigger resource.
	Bind(scope constructs.IConstruct) *awsglue.CfnTrigger_ActionProperty
	BindOptions(_scope constructs.IConstruct) interface{}
}

// The jsii proxy struct for WorkflowJobAction
type jsiiProxy_WorkflowJobAction struct {
	jsiiProxy_WorkflowActionBase
	jsiiProxy_ITriggerAction
}

func (j *jsiiProxy_WorkflowJobAction) Job() IJob {
	var returns IJob
	_jsii_.Get(
		j,
		"job",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkflowJobAction) NotifyDelayAfter() awscdk.Duration {
	var returns awscdk.Duration
	_jsii_.Get(
		j,
		"notifyDelayAfter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkflowJobAction) SecurityConfiguration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkflowJobAction) Timeout() awscdk.Duration {
	var returns awscdk.Duration
	_jsii_.Get(
		j,
		"timeout",
		&returns,
	)
	return returns
}


// Creates a new instance of the WorkflowJobAction class.
func NewWorkflowJobAction(job IJob, options *WorkflowJobActionOptions) WorkflowJobAction {
	_init_.Initialize()

	if err := validateNewWorkflowJobActionParameters(job, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_WorkflowJobAction{}

	_jsii_.Create(
		"cdk-extensions.glue.WorkflowJobAction",
		[]interface{}{job, options},
		&j,
	)

	return &j
}

// Creates a new instance of the WorkflowJobAction class.
func NewWorkflowJobAction_Override(w WorkflowJobAction, job IJob, options *WorkflowJobActionOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.glue.WorkflowJobAction",
		[]interface{}{job, options},
		w,
	)
}

func (w *jsiiProxy_WorkflowJobAction) AddArgument(key *string, value *string) {
	if err := w.validateAddArgumentParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"addArgument",
		[]interface{}{key, value},
	)
}

func (w *jsiiProxy_WorkflowJobAction) Bind(scope constructs.IConstruct) *awsglue.CfnTrigger_ActionProperty {
	if err := w.validateBindParameters(scope); err != nil {
		panic(err)
	}
	var returns *awsglue.CfnTrigger_ActionProperty

	_jsii_.Invoke(
		w,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkflowJobAction) BindOptions(_scope constructs.IConstruct) interface{} {
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

