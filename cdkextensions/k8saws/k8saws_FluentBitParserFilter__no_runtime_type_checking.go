//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package k8saws

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FluentBitParserFilter) validateAddParserParameters(parser IFluentBitParserPlugin) error {
	return nil
}

func (f *jsiiProxy_FluentBitParserFilter) validateBindParameters(_scope constructs.IConstruct) error {
	return nil
}

func (f *jsiiProxy_FluentBitParserFilter) validateRenderConfigFileParameters(config *map[string]interface{}) error {
	return nil
}

func validateNewFluentBitParserFilterParameters(options *FluentBitParserFilterOptions) error {
	return nil
}

