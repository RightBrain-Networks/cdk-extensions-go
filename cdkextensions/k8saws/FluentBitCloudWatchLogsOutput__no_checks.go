//go:build no_runtime_type_checking

package k8saws

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FluentBitCloudWatchLogsOutput) validateBindParameters(scope constructs.IConstruct) error {
	return nil
}

func (f *jsiiProxy_FluentBitCloudWatchLogsOutput) validateRenderConfigFileParameters(config *map[string]interface{}) error {
	return nil
}

func validateNewFluentBitCloudWatchLogsOutputParameters(options *FluentBitCloudWatchLogsOutputOptions) error {
	return nil
}

