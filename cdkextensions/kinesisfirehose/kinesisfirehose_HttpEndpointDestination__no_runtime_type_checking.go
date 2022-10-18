//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package kinesisfirehose

// Building without runtime type checking enabled, so all the below just return nil

func (h *jsiiProxy_HttpEndpointDestination) validateAddCommonAttributeParameters(name *string, value *string) error {
	return nil
}

func (h *jsiiProxy_HttpEndpointDestination) validateAddProcessorParameters(processor DeliveryStreamProcessor) error {
	return nil
}

func (h *jsiiProxy_HttpEndpointDestination) validateBindParameters(scope constructs.IConstruct) error {
	return nil
}

func (h *jsiiProxy_HttpEndpointDestination) validateBuildBackupConfigurationParameters(scope constructs.IConstruct) error {
	return nil
}

func (h *jsiiProxy_HttpEndpointDestination) validateGetOrCreateRoleParameters(scope constructs.IConstruct) error {
	return nil
}

func (h *jsiiProxy_HttpEndpointDestination) validateRenderProcessorConfigurationParameters(scope constructs.IConstruct) error {
	return nil
}

func validateNewHttpEndpointDestinationParameters(url *string, options *HttpEndpointDestinationOptions) error {
	return nil
}

