package glue

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"
)

// Predicate conditions for controlling trigger activation in a Glue workflow.
type WorkflowPredicate interface {
}

// The jsii proxy struct for WorkflowPredicate
type jsiiProxy_WorkflowPredicate struct {
	_ byte // padding
}

func NewWorkflowPredicate() WorkflowPredicate {
	_init_.Initialize()

	j := jsiiProxy_WorkflowPredicate{}

	_jsii_.Create(
		"cdk-extensions.glue.WorkflowPredicate",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewWorkflowPredicate_Override(w WorkflowPredicate) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.glue.WorkflowPredicate",
		nil, // no parameters
		w,
	)
}

// A predicate condition dependent on the completion of a Glue crawler.
//
// Returns: A workflow condition that is predicated on the completion of the
// specified Glue crawler.
func WorkflowPredicate_Crawler(crawler ICrawler, options *WorkflowCrawlerPredicateOptions) WorkflowCrawlerPredicate {
	_init_.Initialize()

	if err := validateWorkflowPredicate_CrawlerParameters(crawler, options); err != nil {
		panic(err)
	}
	var returns WorkflowCrawlerPredicate

	_jsii_.StaticInvoke(
		"cdk-extensions.glue.WorkflowPredicate",
		"crawler",
		[]interface{}{crawler, options},
		&returns,
	)

	return returns
}

// A predicate condition dependent on the completion of a Glue job.
//
// Returns: A workflow condition that is predicated on the completion of the
// specified Glue crawler.
func WorkflowPredicate_Job(job IJob, options *WorkflowJobPredicateOptions) WorkflowJobPredicate {
	_init_.Initialize()

	if err := validateWorkflowPredicate_JobParameters(job, options); err != nil {
		panic(err)
	}
	var returns WorkflowJobPredicate

	_jsii_.StaticInvoke(
		"cdk-extensions.glue.WorkflowPredicate",
		"job",
		[]interface{}{job, options},
		&returns,
	)

	return returns
}

