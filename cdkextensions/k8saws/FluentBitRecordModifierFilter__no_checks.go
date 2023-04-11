//go:build no_runtime_type_checking

package k8saws

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FluentBitRecordModifierFilter) validateAddAllowParameters(tag *string) error {
	return nil
}

func (f *jsiiProxy_FluentBitRecordModifierFilter) validateAddRecordParameters(record *AppendedRecord) error {
	return nil
}

func (f *jsiiProxy_FluentBitRecordModifierFilter) validateAddRemoveParameters(tag *string) error {
	return nil
}

func (f *jsiiProxy_FluentBitRecordModifierFilter) validateBindParameters(_scope constructs.IConstruct) error {
	return nil
}

func (f *jsiiProxy_FluentBitRecordModifierFilter) validateRenderConfigFileParameters(config *map[string]interface{}) error {
	return nil
}

func validateNewFluentBitRecordModifierFilterParameters(options *FluentBitRecordModifierFilterOptions) error {
	return nil
}

