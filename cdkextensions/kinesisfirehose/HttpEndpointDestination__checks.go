//go:build !no_runtime_type_checking

package kinesisfirehose

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

func (h *jsiiProxy_HttpEndpointDestination) validateAddCommonAttributeParameters(name *string, value *string) error {
	if name == nil {
		return fmt.Errorf("parameter name is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func (h *jsiiProxy_HttpEndpointDestination) validateAddProcessorParameters(processor DeliveryStreamProcessor) error {
	if processor == nil {
		return fmt.Errorf("parameter processor is required, but nil was provided")
	}

	return nil
}

func (h *jsiiProxy_HttpEndpointDestination) validateBindParameters(scope constructs.IConstruct) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	return nil
}

func (h *jsiiProxy_HttpEndpointDestination) validateBuildBackupConfigurationParameters(scope constructs.IConstruct) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	return nil
}

func (h *jsiiProxy_HttpEndpointDestination) validateGetOrCreateRoleParameters(scope constructs.IConstruct) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	return nil
}

func (h *jsiiProxy_HttpEndpointDestination) validateRenderProcessorConfigurationParameters(scope constructs.IConstruct) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	return nil
}

func validateNewHttpEndpointDestinationParameters(url *string, options *HttpEndpointDestinationOptions) error {
	if url == nil {
		return fmt.Errorf("parameter url is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

