package kinesisfirehose

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/constructs-go/constructs/v10"
)

type BackupConfiguration interface {
	Destination() IDeliveryStreamBackupDestination
	Enabled() *bool
	Bind(scope constructs.IConstruct) *BackupConfigurationResult
}

// The jsii proxy struct for BackupConfiguration
type jsiiProxy_BackupConfiguration struct {
	_ byte // padding
}

func (j *jsiiProxy_BackupConfiguration) Destination() IDeliveryStreamBackupDestination {
	var returns IDeliveryStreamBackupDestination
	_jsii_.Get(
		j,
		"destination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupConfiguration) Enabled() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}


func NewBackupConfiguration(options *BackupConfigurationOptions) BackupConfiguration {
	_init_.Initialize()

	if err := validateNewBackupConfigurationParameters(options); err != nil {
		panic(err)
	}
	j := jsiiProxy_BackupConfiguration{}

	_jsii_.Create(
		"cdk-extensions.kinesis_firehose.BackupConfiguration",
		[]interface{}{options},
		&j,
	)

	return &j
}

func NewBackupConfiguration_Override(b BackupConfiguration, options *BackupConfigurationOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.kinesis_firehose.BackupConfiguration",
		[]interface{}{options},
		b,
	)
}

func (b *jsiiProxy_BackupConfiguration) Bind(scope constructs.IConstruct) *BackupConfigurationResult {
	if err := b.validateBindParameters(scope); err != nil {
		panic(err)
	}
	var returns *BackupConfigurationResult

	_jsii_.Invoke(
		b,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

