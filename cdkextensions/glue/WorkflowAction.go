package glue

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"
)

// Actions to be started by a Glue workflow trigger when it is activated.
type WorkflowAction interface {
}

// The jsii proxy struct for WorkflowAction
type jsiiProxy_WorkflowAction struct {
	_ byte // padding
}

func NewWorkflowAction() WorkflowAction {
	_init_.Initialize()

	j := jsiiProxy_WorkflowAction{}

	_jsii_.Create(
		"cdk-extensions.glue.WorkflowAction",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewWorkflowAction_Override(w WorkflowAction) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.glue.WorkflowAction",
		nil, // no parameters
		w,
	)
}

// An action that runs a crawler as part of a Glue workflow.
//
// Returns: A workflow action that runs the crawler with the given options.
func WorkflowAction_Crawler(crawler ICrawler, options *WorkflowCrawlerActionOptions) WorkflowCrawlerAction {
	_init_.Initialize()

	if err := validateWorkflowAction_CrawlerParameters(crawler, options); err != nil {
		panic(err)
	}
	var returns WorkflowCrawlerAction

	_jsii_.StaticInvoke(
		"cdk-extensions.glue.WorkflowAction",
		"crawler",
		[]interface{}{crawler, options},
		&returns,
	)

	return returns
}

// An action that runs a Glue job as part of a workflow.
//
// Returns: A workflow action that runs the job with the given options.
func WorkflowAction_Job(job IJob, options *WorkflowJobActionOptions) WorkflowJobAction {
	_init_.Initialize()

	if err := validateWorkflowAction_JobParameters(job, options); err != nil {
		panic(err)
	}
	var returns WorkflowJobAction

	_jsii_.StaticInvoke(
		"cdk-extensions.glue.WorkflowAction",
		"job",
		[]interface{}{job, options},
		&returns,
	)

	return returns
}

