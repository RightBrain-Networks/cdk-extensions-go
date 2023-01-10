package glue

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsglue"
	"github.com/aws/constructs-go/constructs/v10"
)

type WorkflowJobPredicate interface {
	WorkflowPredicateBase
	ITriggerPredicate
	Job() IJob
	LogicalOperator() PredicateLogicalOperator
	State() JobState
	Bind(scope constructs.IConstruct) *awsglue.CfnTrigger_ConditionProperty
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

