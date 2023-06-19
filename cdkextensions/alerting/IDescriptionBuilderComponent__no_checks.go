//go:build no_runtime_type_checking

package alerting

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IDescriptionBuilderComponent) validateAddIteratorParameters(id *string, props *DescriptionBuilderIteratorProps) error {
	return nil
}

func (i *jsiiProxy_IDescriptionBuilderComponent) validateAddReferenceParameters(id *string, props *AddReferenceProps) error {
	return nil
}

func (i *jsiiProxy_IDescriptionBuilderComponent) validateSetDelimiterParameters(id *string, props *SetDelimiterProps) error {
	return nil
}

func (i *jsiiProxy_IDescriptionBuilderComponent) validateWriteParameters(id *string, props *WriteProps) error {
	return nil
}

