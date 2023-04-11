package glue

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsglue"
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents a condition that is predicated on a Glue crawler completion.
//
// The condition can be used to create a trigger that controls the execution of
// downstream tasks in a workflow.
type WorkflowCrawlerPredicate interface {
	WorkflowPredicateBase
	ITriggerPredicate
	// The crawler which must complete in order to meet the requirements to trigger the next stage of the workflow.
	// See: [Trigger Predicate.Conditions.CrawlerName](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-trigger-condition.html#cfn-glue-trigger-condition-crawlername)
	//
	Crawler() ICrawler
	// The logical operator which should be applied in determining whether a crawler meets the requested conditions.
	//
	// At the moment, the only supported operator is `EQUALS`.
	// See: [Trigger Predicate.Conditions.LogicalOperator](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-trigger-condition.html#cfn-glue-trigger-condition-logicaloperator)
	//
	LogicalOperator() PredicateLogicalOperator
	// The state that the crawler must be in in order to meet the criteria to trigger the next stage of the workflow.
	// See: [Trigger Predicate.Conditions.CrawlState](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-trigger-condition.html#cfn-glue-trigger-condition-crawlstate)
	//
	State() CrawlerState
	// Associates the predicate with a construct that is configuring a trigger for a Glue workflow.
	//
	// Returns: A configuration object that can be used to configure a predicate
	// condition for the Glue trigger.
	Bind(scope constructs.IConstruct) *awsglue.CfnTrigger_ConditionProperty
	// Associates the predicate with a construct that is configuring a trigger for a Glue workflow.
	//
	// Returns: A configuration object that can be used to configure a predicate
	// condition for the Glue trigger.
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


// Creates a new instance of the WorkflowCrawlerPredicate class.
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

// Creates a new instance of the WorkflowCrawlerPredicate class.
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

