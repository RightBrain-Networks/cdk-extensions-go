//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package k8saws

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FluentBitModifyFilter) validateAddConditionParameters(condition ModifyCondition) error {
	return nil
}

func (f *jsiiProxy_FluentBitModifyFilter) validateAddOperationParameters(operation ModifyOperation) error {
	return nil
}

func (f *jsiiProxy_FluentBitModifyFilter) validateBindParameters(_scope constructs.IConstruct) error {
	return nil
}

func (f *jsiiProxy_FluentBitModifyFilter) validateRenderConfigFileParameters(config *map[string]interface{}) error {
	return nil
}

func validateNewFluentBitModifyFilterParameters(options *FluentBitModifyFilterOptions) error {
	return nil
}

