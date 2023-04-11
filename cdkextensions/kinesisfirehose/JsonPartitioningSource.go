package kinesisfirehose

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
)

type JsonPartitioningSource interface {
	DynamicPartitioning
	Enabled() *bool
	RetryInterval() awscdk.Duration
	AddPartition(name *string, query *string)
	Bind(scope constructs.IConstruct) *DynamicPartitioningConfiguration
}

// The jsii proxy struct for JsonPartitioningSource
type jsiiProxy_JsonPartitioningSource struct {
	jsiiProxy_DynamicPartitioning
}

func (j *jsiiProxy_JsonPartitioningSource) Enabled() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JsonPartitioningSource) RetryInterval() awscdk.Duration {
	var returns awscdk.Duration
	_jsii_.Get(
		j,
		"retryInterval",
		&returns,
	)
	return returns
}


func NewJsonPartitioningSource(options *JsonPartitioningOptions) JsonPartitioningSource {
	_init_.Initialize()

	if err := validateNewJsonPartitioningSourceParameters(options); err != nil {
		panic(err)
	}
	j := jsiiProxy_JsonPartitioningSource{}

	_jsii_.Create(
		"cdk-extensions.kinesis_firehose.JsonPartitioningSource",
		[]interface{}{options},
		&j,
	)

	return &j
}

func NewJsonPartitioningSource_Override(j JsonPartitioningSource, options *JsonPartitioningOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.kinesis_firehose.JsonPartitioningSource",
		[]interface{}{options},
		j,
	)
}

func JsonPartitioningSource_FromJson(options *JsonPartitioningOptions) JsonPartitioningSource {
	_init_.Initialize()

	if err := validateJsonPartitioningSource_FromJsonParameters(options); err != nil {
		panic(err)
	}
	var returns JsonPartitioningSource

	_jsii_.StaticInvoke(
		"cdk-extensions.kinesis_firehose.JsonPartitioningSource",
		"fromJson",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func JsonPartitioningSource_FromLambda(options *LambdaPartitioningOptions) LambdaPartitioningSource {
	_init_.Initialize()

	if err := validateJsonPartitioningSource_FromLambdaParameters(options); err != nil {
		panic(err)
	}
	var returns LambdaPartitioningSource

	_jsii_.StaticInvoke(
		"cdk-extensions.kinesis_firehose.JsonPartitioningSource",
		"fromLambda",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JsonPartitioningSource) AddPartition(name *string, query *string) {
	if err := j.validateAddPartitionParameters(name, query); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"addPartition",
		[]interface{}{name, query},
	)
}

func (j *jsiiProxy_JsonPartitioningSource) Bind(scope constructs.IConstruct) *DynamicPartitioningConfiguration {
	if err := j.validateBindParameters(scope); err != nil {
		panic(err)
	}
	var returns *DynamicPartitioningConfiguration

	_jsii_.Invoke(
		j,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

