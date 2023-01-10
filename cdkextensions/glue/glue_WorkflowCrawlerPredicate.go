package glue

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsglue"
	"github.com/aws/constructs-go/constructs/v10"
)

type WorkflowCrawlerPredicate interface {
	WorkflowPredicateBase
	ITriggerPredicate
	Crawler() ICrawler
	LogicalOperator() PredicateLogicalOperator
	State() CrawlerState
	Bind(scope constructs.IConstruct) *awsglue.CfnTrigger_ConditionProperty
	BindOptions(_scope constructs.IConstruct) interface{}
}

// The jsii proxy struct for WorkflowCrawlerPredicate
type jsiiProxy_WorkflowCrawlerPredicate struct {
	jsiiProxy_WorkflowPredicateBase
	jsiiProxy_ITriggerPredicate
}

func (j *jsiiProxy_WorkflowCrawlerPredicate) Crawler() ICrawler {
	var returns ICrawler
	_jsii_.Get(
		j,
		"crawler",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkflowCrawlerPredicate) LogicalOperator() PredicateLogicalOperator {
	var returns PredicateLogicalOperator
	_jsii_.Get(
		j,
		"logicalOperator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkflowCrawlerPredicate) State() CrawlerState {
	var returns CrawlerState
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}


func NewWorkflowCrawlerPredicate(crawler ICrawler, options *WorkflowCrawlerPredicateOptions) WorkflowCrawlerPredicate {
	_init_.Initialize()

	if err := validateNewWorkflowCrawlerPredicateParameters(crawler, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_WorkflowCrawlerPredicate{}

	_jsii_.Create(
		"cdk-extensions.glue.WorkflowCrawlerPredicate",
		[]interface{}{crawler, options},
		&j,
	)

	return &j
}

func NewWorkflowCrawlerPredicate_Override(w WorkflowCrawlerPredicate, crawler ICrawler, options *WorkflowCrawlerPredicateOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.glue.WorkflowCrawlerPredicate",
		[]interface{}{crawler, options},
		w,
	)
}

func (w *jsiiProxy_WorkflowCrawlerPredicate) Bind(scope constructs.IConstruct) *awsglue.CfnTrigger_ConditionProperty {
	if err := w.validateBindParameters(scope); err != nil {
		panic(err)
	}
	var returns *awsglue.CfnTrigger_ConditionProperty

	_jsii_.Invoke(
		w,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkflowCrawlerPredicate) BindOptions(_scope constructs.IConstruct) interface{} {
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

