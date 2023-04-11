package kinesisfirehose

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awskinesisfirehose"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/constructs-go/constructs/v10"
)

type CloudWatchLoggingConfiguration interface {
	Enabled() *bool
	LogGroup() awslogs.ILogGroup
	LogStream() awslogs.ILogStream
	Bind(scope constructs.IConstruct) *awskinesisfirehose.CfnDeliveryStream_CloudWatchLoggingOptionsProperty
}

// The jsii proxy struct for CloudWatchLoggingConfiguration
type jsiiProxy_CloudWatchLoggingConfiguration struct {
	_ byte // padding
}

func (j *jsiiProxy_CloudWatchLoggingConfiguration) Enabled() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudWatchLoggingConfiguration) LogGroup() awslogs.ILogGroup {
	var returns awslogs.ILogGroup
	_jsii_.Get(
		j,
		"logGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudWatchLoggingConfiguration) LogStream() awslogs.ILogStream {
	var returns awslogs.ILogStream
	_jsii_.Get(
		j,
		"logStream",
		&returns,
	)
	return returns
}


func NewCloudWatchLoggingConfiguration(options *CloudWatchLoggingConfigurationOptions) CloudWatchLoggingConfiguration {
	_init_.Initialize()

	if err := validateNewCloudWatchLoggingConfigurationParameters(options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CloudWatchLoggingConfiguration{}

	_jsii_.Create(
		"cdk-extensions.kinesis_firehose.CloudWatchLoggingConfiguration",
		[]interface{}{options},
		&j,
	)

	return &j
}

func NewCloudWatchLoggingConfiguration_Override(c CloudWatchLoggingConfiguration, options *CloudWatchLoggingConfigurationOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.kinesis_firehose.CloudWatchLoggingConfiguration",
		[]interface{}{options},
		c,
	)
}

func (c *jsiiProxy_CloudWatchLoggingConfiguration) Bind(scope constructs.IConstruct) *awskinesisfirehose.CfnDeliveryStream_CloudWatchLoggingOptionsProperty {
	if err := c.validateBindParameters(scope); err != nil {
		panic(err)
	}
	var returns *awskinesisfirehose.CfnDeliveryStream_CloudWatchLoggingOptionsProperty

	_jsii_.Invoke(
		c,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

