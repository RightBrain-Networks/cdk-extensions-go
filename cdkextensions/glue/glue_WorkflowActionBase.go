package glue

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
)

// Base class providing common functionality for workflow trigger actions.
type WorkflowActionBase interface {
	// After a job run starts, the number of minutes to wait before sending a job run delay notification.
	// See: [Trigger Actions.NotificationProperty.NotifyDelayAfter](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-trigger-notificationproperty.html#cfn-glue-trigger-notificationproperty-notifydelayafter)
	//
	NotifyDelayAfter() awscdk.Duration
	// The name of the SecurityConfiguration structure to be used with this action.
	// See: [Trigger Actions.SecurityConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-trigger-action.html#cfn-glue-trigger-action-securityconfiguration)
	//
	SecurityConfiguration() *string
	// The `JobRun` timeout in minutes.
	//
	// This is the maximum time that a job run
	// can consume resources before it is terminated and enters TIMEOUT status.
	// The default is 48 hours. This overrides the timeout value set in the
	// parent job.
	// See: [Trigger Actions.Timeout](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-trigger-action.html#cfn-glue-trigger-action-timeout)
	//
	Timeout() awscdk.Duration
	// Adds an argument that will be passed to the specified action when triggered as part of a workflow.
	// See: [AWS Glue job parameters](https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html)
	//
	AddArgument(key *string, value *string)
	// Associates the action with a construct that is configuring a trigger for a Glue workflow.
	//
	// Returns: A configuration object that can be used to configure a triggered
	// workflow action.
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


// Creates a new instance of the WorkflowActionBase class.
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

// Creates a new instance of the WorkflowActionBase class.
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

