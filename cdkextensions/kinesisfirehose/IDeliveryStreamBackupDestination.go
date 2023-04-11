package kinesisfirehose

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

type IDeliveryStreamBackupDestination interface {
	RenderBackupConfiguration(scope constructs.IConstruct, enabled *bool) *BackupConfigurationResult
}

// The jsii proxy for IDeliveryStreamBackupDestination
type jsiiProxy_IDeliveryStreamBackupDestination struct {
	_ byte // padding
}

func (i *jsiiProxy_IDeliveryStreamBackupDestination) RenderBackupConfiguration(scope constructs.IConstruct, enabled *bool) *BackupConfigurationResult {
	if err := i.validateRenderBackupConfigurationParameters(scope); err != nil {
		panic(err)
	}
	var returns *BackupConfigurationResult

	_jsii_.Invoke(
		i,
		"renderBackupConfiguration",
		[]interface{}{scope, enabled},
		&returns,
	)

	return returns
}

