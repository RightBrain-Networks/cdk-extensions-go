//go:build !no_runtime_type_checking

package kinesisfirehose

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/constructs-go/constructs/v10"
)

func (e *jsiiProxy_ExtendedS3Destination) validateAddProcessorParameters(processor DeliveryStreamProcessor) error {
	if processor == nil {
		return fmt.Errorf("parameter processor is required, but nil was provided")
	}

	return nil
}

func (e *jsiiProxy_ExtendedS3Destination) validateBindParameters(scope constructs.IConstruct) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	return nil
}

func (e *jsiiProxy_ExtendedS3Destination) validateBuildConfigurationParameters(scope constructs.IConstruct) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	return nil
}

func (e *jsiiProxy_ExtendedS3Destination) validateRenderBackupConfigurationParameters(scope constructs.IConstruct) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	return nil
}

func (e *jsiiProxy_ExtendedS3Destination) validateRenderProcessorConfigurationParameters(scope constructs.IConstruct) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	return nil
}

func validateNewExtendedS3DestinationParameters(bucket awss3.IBucket, options *ExtendedS3DestinationOptions) error {
	if bucket == nil {
		return fmt.Errorf("parameter bucket is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

