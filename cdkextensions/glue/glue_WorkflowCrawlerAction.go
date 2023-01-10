package glue

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsglue"
	"github.com/aws/constructs-go/constructs/v10"
)

// Configuration options for the WorkflowCrawlerAction class.
type WorkflowCrawlerAction interface {
	WorkflowActionBase
	ITriggerAction
	// The Glue crawler to be triggered as part of the workflow.
	Crawler() ICrawler
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

// The jsii proxy struct for WorkflowCrawlerAction
type jsiiProxy_WorkflowCrawlerAction struct {
	jsiiProxy_WorkflowActionBase
	jsiiProxy_ITriggerAction
}

func (j *jsiiProxy_WorkflowCrawlerAction) Crawler() ICrawler {
	var returns ICrawler
	_jsii_.Get(
		j,
		"crawler",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkflowCrawlerAction) NotifyDelayAfter() awscdk.Duration {
	var returns awscdk.Duration
	_jsii_.Get(
		j,
		"notifyDelayAfter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkflowCrawlerAction) SecurityConfiguration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkflowCrawlerAction) Timeout() awscdk.Duration {
	var returns awscdk.Duration
	_jsii_.Get(
		j,
		"timeout",
		&returns,
	)
	return returns
}


// Creates a new instance of the WorkflowCrawlerAction class.
func NewWorkflowCrawlerAction(crawler ICrawler, options *WorkflowActionOptions) WorkflowCrawlerAction {
	_init_.Initialize()

	if err := validateNewWorkflowCrawlerActionParameters(crawler, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_WorkflowCrawlerAction{}

	_jsii_.Create(
		"cdk-extensions.glue.WorkflowCrawlerAction",
		[]interface{}{crawler, options},
		&j,
	)

	return &j
}

// Creates a new instance of the WorkflowCrawlerAction class.
func NewWorkflowCrawlerAction_Override(w WorkflowCrawlerAction, crawler ICrawler, options *WorkflowActionOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.glue.WorkflowCrawlerAction",
		[]interface{}{crawler, options},
		w,
	)
}

func (w *jsiiProxy_WorkflowCrawlerAction) AddArgument(key *string, value *string) {
	if err := w.validateAddArgumentParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"addArgument",
		[]interface{}{key, value},
	)
}

func (w *jsiiProxy_WorkflowCrawlerAction) Bind(scope constructs.IConstruct) *awsglue.CfnTrigger_ActionProperty {
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

func (w *jsiiProxy_WorkflowCrawlerAction) BindOptions(_scope constructs.IConstruct) interface{} {
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

