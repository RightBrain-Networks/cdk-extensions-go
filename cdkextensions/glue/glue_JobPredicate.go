package glue

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsglue"
	"github.com/aws/constructs-go/constructs/v10"
)

type JobPredicate interface {
	PredicateBase
	ITriggerPredicate
	Job() IJob
	LogicalOperator() PredicateLogicalOperator
	State() JobState
	Bind(scope constructs.IConstruct) *awsglue.CfnTrigger_ConditionProperty
	BindOptions(_scope constructs.IConstruct) interface{}
}

// The jsii proxy struct for JobPredicate
type jsiiProxy_JobPredicate struct {
	jsiiProxy_PredicateBase
	jsiiProxy_ITriggerPredicate
}

func (j *jsiiProxy_JobPredicate) Job() IJob {
	var returns IJob
	_jsii_.Get(
		j,
		"job",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobPredicate) LogicalOperator() PredicateLogicalOperator {
	var returns PredicateLogicalOperator
	_jsii_.Get(
		j,
		"logicalOperator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobPredicate) State() JobState {
	var returns JobState
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}


func NewJobPredicate(job IJob, options *JobPredicateOptions) JobPredicate {
	_init_.Initialize()

	if err := validateNewJobPredicateParameters(job, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_JobPredicate{}

	_jsii_.Create(
		"cdk-extensions.glue.JobPredicate",
		[]interface{}{job, options},
		&j,
	)

	return &j
}

func NewJobPredicate_Override(j JobPredicate, job IJob, options *JobPredicateOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.glue.JobPredicate",
		[]interface{}{job, options},
		j,
	)
}

func (j *jsiiProxy_JobPredicate) Bind(scope constructs.IConstruct) *awsglue.CfnTrigger_ConditionProperty {
	if err := j.validateBindParameters(scope); err != nil {
		panic(err)
	}
	var returns *awsglue.CfnTrigger_ConditionProperty

	_jsii_.Invoke(
		j,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobPredicate) BindOptions(_scope constructs.IConstruct) interface{} {
	if err := j.validateBindOptionsParameters(_scope); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		j,
		"bindOptions",
		[]interface{}{_scope},
		&returns,
	)

	return returns
}

