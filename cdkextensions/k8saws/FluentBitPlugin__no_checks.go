//go:build no_runtime_type_checking

package k8saws

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FluentBitPlugin) validateBindParameters(scope constructs.IConstruct) error {
	return nil
}

func (f *jsiiProxy_FluentBitPlugin) validateRenderConfigFileParameters(config *map[string]interface{}) error {
	return nil
}

func validateNewFluentBitPluginParameters(options *FluentBitPluginCommonOptions) error {
	return nil
}

