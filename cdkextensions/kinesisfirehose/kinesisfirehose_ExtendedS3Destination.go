package kinesisfirehose

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskinesisfirehose"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/constructs-go/constructs/v10"
)

type ExtendedS3Destination interface {
	S3Destination
	BackupConfiguration() BackupConfiguration
	Bucket() awss3.IBucket
	Buffering() BufferingConfiguration
	CloudwatchLoggingConfiguration() CloudWatchLoggingConfiguration
	CompressionFormat() S3CompressionFormat
	DataFormatConversion() DataFormatConversion
	DynamicPartitioning() DynamicPartitioning
	EncryptionEnabled() *bool
	EncryptionKey() awskms.IKey
	ErrorOutputPrefix() *string
	KeyPrefix() *string
	ProcessingEnabled() *bool
	ProcessorConfiguration() ProcessorConfiguration
	Processors() *[]DeliveryStreamProcessor
	Role() awsiam.IRole
	AddProcessor(processor DeliveryStreamProcessor) ExtendedS3Destination
	Bind(scope constructs.IConstruct) *DeliveryStreamDestinationConfiguration
	BuildConfiguration(scope constructs.IConstruct) *awskinesisfirehose.CfnDeliveryStream_S3DestinationConfigurationProperty
	RenderBackupConfiguration(scope constructs.IConstruct, enabled *bool) *BackupConfigurationResult
	RenderProcessorConfiguration(scope constructs.IConstruct) *awskinesisfirehose.CfnDeliveryStream_ProcessingConfigurationProperty
}

// The jsii proxy struct for ExtendedS3Destination
type jsiiProxy_ExtendedS3Destination struct {
	jsiiProxy_S3Destination
}

func (j *jsiiProxy_ExtendedS3Destination) BackupConfiguration() BackupConfiguration {
	var returns BackupConfiguration
	_jsii_.Get(
		j,
		"backupConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExtendedS3Destination) Bucket() awss3.IBucket {
	var returns awss3.IBucket
	_jsii_.Get(
		j,
		"bucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExtendedS3Destination) Buffering() BufferingConfiguration {
	var returns BufferingConfiguration
	_jsii_.Get(
		j,
		"buffering",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExtendedS3Destination) CloudwatchLoggingConfiguration() CloudWatchLoggingConfiguration {
	var returns CloudWatchLoggingConfiguration
	_jsii_.Get(
		j,
		"cloudwatchLoggingConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExtendedS3Destination) CompressionFormat() S3CompressionFormat {
	var returns S3CompressionFormat
	_jsii_.Get(
		j,
		"compressionFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExtendedS3Destination) DataFormatConversion() DataFormatConversion {
	var returns DataFormatConversion
	_jsii_.Get(
		j,
		"dataFormatConversion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExtendedS3Destination) DynamicPartitioning() DynamicPartitioning {
	var returns DynamicPartitioning
	_jsii_.Get(
		j,
		"dynamicPartitioning",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExtendedS3Destination) EncryptionEnabled() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"encryptionEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExtendedS3Destination) EncryptionKey() awskms.IKey {
	var returns awskms.IKey
	_jsii_.Get(
		j,
		"encryptionKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExtendedS3Destination) ErrorOutputPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"errorOutputPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExtendedS3Destination) KeyPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExtendedS3Destination) ProcessingEnabled() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"processingEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExtendedS3Destination) ProcessorConfiguration() ProcessorConfiguration {
	var returns ProcessorConfiguration
	_jsii_.Get(
		j,
		"processorConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExtendedS3Destination) Processors() *[]DeliveryStreamProcessor {
	var returns *[]DeliveryStreamProcessor
	_jsii_.Get(
		j,
		"processors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExtendedS3Destination) Role() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}


func NewExtendedS3Destination(bucket awss3.IBucket, options *ExtendedS3DestinationOptions) ExtendedS3Destination {
	_init_.Initialize()

	if err := validateNewExtendedS3DestinationParameters(bucket, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_ExtendedS3Destination{}

	_jsii_.Create(
		"cdk-extensions.kinesis_firehose.ExtendedS3Destination",
		[]interface{}{bucket, options},
		&j,
	)

	return &j
}

func NewExtendedS3Destination_Override(e ExtendedS3Destination, bucket awss3.IBucket, options *ExtendedS3DestinationOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.kinesis_firehose.ExtendedS3Destination",
		[]interface{}{bucket, options},
		e,
	)
}

func (e *jsiiProxy_ExtendedS3Destination) AddProcessor(processor DeliveryStreamProcessor) ExtendedS3Destination {
	if err := e.validateAddProcessorParameters(processor); err != nil {
		panic(err)
	}
	var returns ExtendedS3Destination

	_jsii_.Invoke(
		e,
		"addProcessor",
		[]interface{}{processor},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExtendedS3Destination) Bind(scope constructs.IConstruct) *DeliveryStreamDestinationConfiguration {
	if err := e.validateBindParameters(scope); err != nil {
		panic(err)
	}
	var returns *DeliveryStreamDestinationConfiguration

	_jsii_.Invoke(
		e,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExtendedS3Destination) BuildConfiguration(scope constructs.IConstruct) *awskinesisfirehose.CfnDeliveryStream_S3DestinationConfigurationProperty {
	if err := e.validateBuildConfigurationParameters(scope); err != nil {
		panic(err)
	}
	var returns *awskinesisfirehose.CfnDeliveryStream_S3DestinationConfigurationProperty

	_jsii_.Invoke(
		e,
		"buildConfiguration",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExtendedS3Destination) RenderBackupConfiguration(scope constructs.IConstruct, enabled *bool) *BackupConfigurationResult {
	if err := e.validateRenderBackupConfigurationParameters(scope); err != nil {
		panic(err)
	}
	var returns *BackupConfigurationResult

	_jsii_.Invoke(
		e,
		"renderBackupConfiguration",
		[]interface{}{scope, enabled},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExtendedS3Destination) RenderProcessorConfiguration(scope constructs.IConstruct) *awskinesisfirehose.CfnDeliveryStream_ProcessingConfigurationProperty {
	if err := e.validateRenderProcessorConfigurationParameters(scope); err != nil {
		panic(err)
	}
	var returns *awskinesisfirehose.CfnDeliveryStream_ProcessingConfigurationProperty

	_jsii_.Invoke(
		e,
		"renderProcessorConfiguration",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

