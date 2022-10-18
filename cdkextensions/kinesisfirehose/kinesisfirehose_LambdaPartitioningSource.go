package kinesisfirehose

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/constructs-go/constructs/v10"
)

type LambdaPartitioningSource interface {
	DynamicPartitioning
	Enabled() *bool
	LambdaFunction() awslambda.IFunction
	RetryInterval() awscdk.Duration
	Bind(scope constructs.IConstruct) *DynamicPartitioningConfiguration
}

// The jsii proxy struct for LambdaPartitioningSource
type jsiiProxy_LambdaPartitioningSource struct {
	jsiiProxy_DynamicPartitioning
}

func (j *jsiiProxy_LambdaPartitioningSource) Enabled() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaPartitioningSource) LambdaFunction() awslambda.IFunction {
	var returns awslambda.IFunction
	_jsii_.Get(
		j,
		"lambdaFunction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaPartitioningSource) RetryInterval() awscdk.Duration {
	var returns awscdk.Duration
	_jsii_.Get(
		j,
		"retryInterval",
		&returns,
	)
	return returns
}


func NewLambdaPartitioningSource(options *LambdaPartitioningOptions) LambdaPartitioningSource {
	_init_.Initialize()

	if err := validateNewLambdaPartitioningSourceParameters(options); err != nil {
		panic(err)
	}
	j := jsiiProxy_LambdaPartitioningSource{}

	_jsii_.Create(
		"cdk-extensions.kinesis_firehose.LambdaPartitioningSource",
		[]interface{}{options},
		&j,
	)

	return &j
}

func NewLambdaPartitioningSource_Override(l LambdaPartitioningSource, options *LambdaPartitioningOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.kinesis_firehose.LambdaPartitioningSource",
		[]interface{}{options},
		l,
	)
}

func LambdaPartitioningSource_FromJson(options *JsonPartitioningOptions) JsonPartitioningSource {
	_init_.Initialize()

	if err := validateLambdaPartitioningSource_FromJsonParameters(options); err != nil {
		panic(err)
	}
	var returns JsonPartitioningSource

	_jsii_.StaticInvoke(
		"cdk-extensions.kinesis_firehose.LambdaPartitioningSource",
		"fromJson",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func LambdaPartitioningSource_FromLambda(options *LambdaPartitioningOptions) LambdaPartitioningSource {
	_init_.Initialize()

	if err := validateLambdaPartitioningSource_FromLambdaParameters(options); err != nil {
		panic(err)
	}
	var returns LambdaPartitioningSource

	_jsii_.StaticInvoke(
		"cdk-extensions.kinesis_firehose.LambdaPartitioningSource",
		"fromLambda",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaPartitioningSource) Bind(scope constructs.IConstruct) *DynamicPartitioningConfiguration {
	if err := l.validateBindParameters(scope); err != nil {
		panic(err)
	}
	var returns *DynamicPartitioningConfiguration

	_jsii_.Invoke(
		l,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

