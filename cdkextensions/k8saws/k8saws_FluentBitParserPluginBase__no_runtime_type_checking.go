//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package k8saws

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FluentBitParserPluginBase) validateBindParameters(scope constructs.IConstruct) error {
	return nil
}

func (f *jsiiProxy_FluentBitParserPluginBase) validateRenderConfigFileParameters(config *map[string]interface{}) error {
	return nil
}

func validateNewFluentBitParserPluginBaseParameters(name *string, format *string, _options *FluentBitParserPluginCommonOptions) error {
	return nil
}

