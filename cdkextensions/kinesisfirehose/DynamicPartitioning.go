package kinesisfirehose

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
)

type DynamicPartitioning interface {
	Enabled() *bool
	RetryInterval() awscdk.Duration
	Bind(_scope constructs.IConstruct) *DynamicPartitioningConfiguration
}

// The jsii proxy struct for DynamicPartitioning
type jsiiProxy_DynamicPartitioning struct {
	_ byte // padding
}

func (j *jsiiProxy_DynamicPartitioning) Enabled() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamicPartitioning) RetryInterval() awscdk.Duration {
	var returns awscdk.Duration
	_jsii_.Get(
		j,
		"retryInterval",
		&returns,
	)
	return returns
}


func NewDynamicPartitioning(options *CommonPartitioningOptions) DynamicPartitioning {
	_init_.Initialize()

	if err := validateNewDynamicPartitioningParameters(options); err != nil {
		panic(err)
	}
	j := jsiiProxy_DynamicPartitioning{}

	_jsii_.Create(
		"cdk-extensions.kinesis_firehose.DynamicPartitioning",
		[]interface{}{options},
		&j,
	)

	return &j
}

func NewDynamicPartitioning_Override(d DynamicPartitioning, options *CommonPartitioningOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.kinesis_firehose.DynamicPartitioning",
		[]interface{}{options},
		d,
	)
}

func DynamicPartitioning_FromJson(options *JsonPartitioningOptions) JsonPartitioningSource {
	_init_.Initialize()

	if err := validateDynamicPartitioning_FromJsonParameters(options); err != nil {
		panic(err)
	}
	var returns JsonPartitioningSource

	_jsii_.StaticInvoke(
		"cdk-extensions.kinesis_firehose.DynamicPartitioning",
		"fromJson",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func DynamicPartitioning_FromLambda(options *LambdaPartitioningOptions) LambdaPartitioningSource {
	_init_.Initialize()

	if err := validateDynamicPartitioning_FromLambdaParameters(options); err != nil {
		panic(err)
	}
	var returns LambdaPartitioningSource

	_jsii_.StaticInvoke(
		"cdk-extensions.kinesis_firehose.DynamicPartitioning",
		"fromLambda",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamicPartitioning) Bind(_scope constructs.IConstruct) *DynamicPartitioningConfiguration {
	if err := d.validateBindParameters(_scope); err != nil {
		panic(err)
	}
	var returns *DynamicPartitioningConfiguration

	_jsii_.Invoke(
		d,
		"bind",
		[]interface{}{_scope},
		&returns,
	)

	return returns
}

