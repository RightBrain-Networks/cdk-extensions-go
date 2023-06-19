//go:build no_runtime_type_checking

package alerting

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DescriptionBuilderIterator) validateAddIteratorParameters(id *string, props *DescriptionBuilderIteratorProps) error {
	return nil
}

func (d *jsiiProxy_DescriptionBuilderIterator) validateAddReferenceParameters(id *string, props *AddReferenceProps) error {
	return nil
}

func (d *jsiiProxy_DescriptionBuilderIterator) validateBuildIdParameters(prefix *string) error {
	return nil
}

func (d *jsiiProxy_DescriptionBuilderIterator) validateRegisterBuilderParameters(builder IDescriptionBuilderComponent) error {
	return nil
}

func (d *jsiiProxy_DescriptionBuilderIterator) validateRegisterChainableParameters(chainable awsstepfunctions.IChainable) error {
	return nil
}

func (d *jsiiProxy_DescriptionBuilderIterator) validateSetDelimiterParameters(id *string, props *SetDelimiterProps) error {
	return nil
}

func (d *jsiiProxy_DescriptionBuilderIterator) validateWriteParameters(id *string, props *WriteProps) error {
	return nil
}

func validateDescriptionBuilderIterator_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewDescriptionBuilderIteratorParameters(scope constructs.IConstruct, id *string, props *DescriptionBuilderIteratorProps) error {
	return nil
}

