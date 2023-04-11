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

type S3Destination interface {
	DeliveryStreamDestination
	IDeliveryStreamBackupDestination
	Bucket() awss3.IBucket
	Buffering() BufferingConfiguration
	CloudwatchLoggingConfiguration() CloudWatchLoggingConfiguration
	CompressionFormat() S3CompressionFormat
	EncryptionEnabled() *bool
	EncryptionKey() awskms.IKey
	ErrorOutputPrefix() *string
	KeyPrefix() *string
	Role() awsiam.IRole
	Bind(scope constructs.IConstruct) *DeliveryStreamDestinationConfiguration
	BuildConfiguration(scope constructs.IConstruct) *awskinesisfirehose.CfnDeliveryStream_S3DestinationConfigurationProperty
	RenderBackupConfiguration(scope constructs.IConstruct, enabled *bool) *BackupConfigurationResult
}

// The jsii proxy struct for S3Destination
type jsiiProxy_S3Destination struct {
	jsiiProxy_DeliveryStreamDestination
	jsiiProxy_IDeliveryStreamBackupDestination
}

func (j *jsiiProxy_S3Destination) Bucket() awss3.IBucket {
	var returns awss3.IBucket
	_jsii_.Get(
		j,
		"bucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Destination) Buffering() BufferingConfiguration {
	var returns BufferingConfiguration
	_jsii_.Get(
		j,
		"buffering",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Destination) CloudwatchLoggingConfiguration() CloudWatchLoggingConfiguration {
	var returns CloudWatchLoggingConfiguration
	_jsii_.Get(
		j,
		"cloudwatchLoggingConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Destination) CompressionFormat() S3CompressionFormat {
	var returns S3CompressionFormat
	_jsii_.Get(
		j,
		"compressionFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Destination) EncryptionEnabled() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"encryptionEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Destination) EncryptionKey() awskms.IKey {
	var returns awskms.IKey
	_jsii_.Get(
		j,
		"encryptionKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Destination) ErrorOutputPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"errorOutputPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Destination) KeyPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Destination) Role() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}


func NewS3Destination(bucket awss3.IBucket, options *S3DestinationOptions) S3Destination {
	_init_.Initialize()

	if err := validateNewS3DestinationParameters(bucket, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_S3Destination{}

	_jsii_.Create(
		"cdk-extensions.kinesis_firehose.S3Destination",
		[]interface{}{bucket, options},
		&j,
	)

	return &j
}

func NewS3Destination_Override(s S3Destination, bucket awss3.IBucket, options *S3DestinationOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.kinesis_firehose.S3Destination",
		[]interface{}{bucket, options},
		s,
	)
}

func (s *jsiiProxy_S3Destination) Bind(scope constructs.IConstruct) *DeliveryStreamDestinationConfiguration {
	if err := s.validateBindParameters(scope); err != nil {
		panic(err)
	}
	var returns *DeliveryStreamDestinationConfiguration

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3Destination) BuildConfiguration(scope constructs.IConstruct) *awskinesisfirehose.CfnDeliveryStream_S3DestinationConfigurationProperty {
	if err := s.validateBuildConfigurationParameters(scope); err != nil {
		panic(err)
	}
	var returns *awskinesisfirehose.CfnDeliveryStream_S3DestinationConfigurationProperty

	_jsii_.Invoke(
		s,
		"buildConfiguration",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3Destination) RenderBackupConfiguration(scope constructs.IConstruct, enabled *bool) *BackupConfigurationResult {
	if err := s.validateRenderBackupConfigurationParameters(scope); err != nil {
		panic(err)
	}
	var returns *BackupConfigurationResult

	_jsii_.Invoke(
		s,
		"renderBackupConfiguration",
		[]interface{}{scope, enabled},
		&returns,
	)

	return returns
}

