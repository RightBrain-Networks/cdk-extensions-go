package kinesisfirehose

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskinesisfirehose"
	"github.com/aws/constructs-go/constructs/v10"
)

type HttpEndpointDestination interface {
	DeliveryStreamDestination
	AccessKey() awscdk.SecretValue
	BackupConfiguration() BackupConfiguration
	Buffering() BufferingConfiguration
	CloudwatchLoggingConfiguration() CloudWatchLoggingConfiguration
	CommonAttributes() *map[string]*string
	ContentEncoding() ContentEncoding
	EndpointName() *string
	EndpointUrl() *string
	ProcessingEnabled() *bool
	ProcessorConfiguration() ProcessorConfiguration
	RetryDuration() awscdk.Duration
	Role() awsiam.IRole
	AddCommonAttribute(name *string, value *string) HttpEndpointDestination
	AddProcessor(processor DeliveryStreamProcessor) HttpEndpointDestination
	Bind(scope constructs.IConstruct) *DeliveryStreamDestinationConfiguration
	BuildBackupConfiguration(scope constructs.IConstruct) BackupConfiguration
	GetOrCreateRole(scope constructs.IConstruct) awsiam.IRole
	RenderProcessorConfiguration(scope constructs.IConstruct) *awskinesisfirehose.CfnDeliveryStream_ProcessingConfigurationProperty
}

// The jsii proxy struct for HttpEndpointDestination
type jsiiProxy_HttpEndpointDestination struct {
	jsiiProxy_DeliveryStreamDestination
}

func (j *jsiiProxy_HttpEndpointDestination) AccessKey() awscdk.SecretValue {
	var returns awscdk.SecretValue
	_jsii_.Get(
		j,
		"accessKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpEndpointDestination) BackupConfiguration() BackupConfiguration {
	var returns BackupConfiguration
	_jsii_.Get(
		j,
		"backupConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpEndpointDestination) Buffering() BufferingConfiguration {
	var returns BufferingConfiguration
	_jsii_.Get(
		j,
		"buffering",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpEndpointDestination) CloudwatchLoggingConfiguration() CloudWatchLoggingConfiguration {
	var returns CloudWatchLoggingConfiguration
	_jsii_.Get(
		j,
		"cloudwatchLoggingConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpEndpointDestination) CommonAttributes() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"commonAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpEndpointDestination) ContentEncoding() ContentEncoding {
	var returns ContentEncoding
	_jsii_.Get(
		j,
		"contentEncoding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpEndpointDestination) EndpointName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpEndpointDestination) EndpointUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpEndpointDestination) ProcessingEnabled() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"processingEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpEndpointDestination) ProcessorConfiguration() ProcessorConfiguration {
	var returns ProcessorConfiguration
	_jsii_.Get(
		j,
		"processorConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpEndpointDestination) RetryDuration() awscdk.Duration {
	var returns awscdk.Duration
	_jsii_.Get(
		j,
		"retryDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpEndpointDestination) Role() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}


func NewHttpEndpointDestination(url *string, options *HttpEndpointDestinationOptions) HttpEndpointDestination {
	_init_.Initialize()

	if err := validateNewHttpEndpointDestinationParameters(url, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_HttpEndpointDestination{}

	_jsii_.Create(
		"cdk-extensions.kinesis_firehose.HttpEndpointDestination",
		[]interface{}{url, options},
		&j,
	)

	return &j
}

func NewHttpEndpointDestination_Override(h HttpEndpointDestination, url *string, options *HttpEndpointDestinationOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.kinesis_firehose.HttpEndpointDestination",
		[]interface{}{url, options},
		h,
	)
}

func (h *jsiiProxy_HttpEndpointDestination) AddCommonAttribute(name *string, value *string) HttpEndpointDestination {
	if err := h.validateAddCommonAttributeParameters(name, value); err != nil {
		panic(err)
	}
	var returns HttpEndpointDestination

	_jsii_.Invoke(
		h,
		"addCommonAttribute",
		[]interface{}{name, value},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HttpEndpointDestination) AddProcessor(processor DeliveryStreamProcessor) HttpEndpointDestination {
	if err := h.validateAddProcessorParameters(processor); err != nil {
		panic(err)
	}
	var returns HttpEndpointDestination

	_jsii_.Invoke(
		h,
		"addProcessor",
		[]interface{}{processor},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HttpEndpointDestination) Bind(scope constructs.IConstruct) *DeliveryStreamDestinationConfiguration {
	if err := h.validateBindParameters(scope); err != nil {
		panic(err)
	}
	var returns *DeliveryStreamDestinationConfiguration

	_jsii_.Invoke(
		h,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HttpEndpointDestination) BuildBackupConfiguration(scope constructs.IConstruct) BackupConfiguration {
	if err := h.validateBuildBackupConfigurationParameters(scope); err != nil {
		panic(err)
	}
	var returns BackupConfiguration

	_jsii_.Invoke(
		h,
		"buildBackupConfiguration",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HttpEndpointDestination) GetOrCreateRole(scope constructs.IConstruct) awsiam.IRole {
	if err := h.validateGetOrCreateRoleParameters(scope); err != nil {
		panic(err)
	}
	var returns awsiam.IRole

	_jsii_.Invoke(
		h,
		"getOrCreateRole",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HttpEndpointDestination) RenderProcessorConfiguration(scope constructs.IConstruct) *awskinesisfirehose.CfnDeliveryStream_ProcessingConfigurationProperty {
	if err := h.validateRenderProcessorConfigurationParameters(scope); err != nil {
		panic(err)
	}
	var returns *awskinesisfirehose.CfnDeliveryStream_ProcessingConfigurationProperty

	_jsii_.Invoke(
		h,
		"renderProcessorConfiguration",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

