package glue

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsglue"
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents a condition that is predicated on a Glue job completion.
//
// The condition can be used to create a trigger that controls the execution of
// downstream tasks in a workflow.
type WorkflowJobPredicate interface {
	WorkflowPredicateBase
	ITriggerPredicate
	// The job which must complete in order to meet the requirements to trigger the next stage of the workflow.
	// See: [Trigger Predicate.Conditions.JobName](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-trigger-condition.html#cfn-glue-trigger-condition-jobname)
	//
	Job() IJob
	// The logical operator which should be applied in determining whether a job meets the requested conditions.
	//
	// At the moment, the only supported operator is `EQUALS`.
	// See: [Trigger Predicate.Conditions.LogicalOperator](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-trigger-condition.html#cfn-glue-trigger-condition-logicaloperator)
	//
	LogicalOperator() PredicateLogicalOperator
	// The state that the job must be in in order to meet the criteria to trigger the next stage of the workflow.
	// See: [Trigger Predicate.Conditions.State](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-trigger-condition.html#cfn-glue-trigger-condition-state)
	//
	State() JobState
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

// The jsii proxy struct for WorkflowJobPredicate
type jsiiProxy_WorkflowJobPredicate struct {
	jsiiProxy_WorkflowPredicateBase
	jsiiProxy_ITriggerPredicate
}

func (j *jsiiProxy_WorkflowJobPredicate) Job() IJob {
	var returns IJob
	_jsii_.Get(
		j,
		"job",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkflowJobPredicate) LogicalOperator() PredicateLogicalOperator {
	var returns PredicateLogicalOperator
	_jsii_.Get(
		j,
		"logicalOperator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkflowJobPredicate) State() JobState {
	var returns JobState
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}


// Creates a new instance of the WorkflowJobPredicate class.
func NewWorkflowJobPredicate(job IJob, options *WorkflowJobPredicateOptions) WorkflowJobPredicate {
	_init_.Initialize()

	if err := validateNewWorkflowJobPredicateParameters(job, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_WorkflowJobPredicate{}

	_jsii_.Create(
		"cdk-extensions.glue.WorkflowJobPredicate",
		[]interface{}{job, options},
		&j,
	)

	return &j
}

// Creates a new instance of the WorkflowJobPredicate class.
func NewWorkflowJobPredicate_Override(w WorkflowJobPredicate, job IJob, options *WorkflowJobPredicateOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.glue.WorkflowJobPredicate",
		[]interface{}{job, options},
		w,
	)
}

func (w *jsiiProxy_WorkflowJobPredicate) Bind(scope constructs.IConstruct) *awsglue.CfnTrigger_ConditionProperty {
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

func (w *jsiiProxy_WorkflowJobPredicate) BindOptions(_scope constructs.IConstruct) interface{} {
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

