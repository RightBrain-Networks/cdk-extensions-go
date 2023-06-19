//go:build no_runtime_type_checking

package alerting

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DescriptionBuilderSection) validateAddIteratorParameters(id *string, props *DescriptionBuilderIteratorProps) error {
	return nil
}

func (d *jsiiProxy_DescriptionBuilderSection) validateAddReferenceParameters(id *string, props *AddReferenceProps) error {
	return nil
}

func (d *jsiiProxy_DescriptionBuilderSection) validateAddReferenceCheckParameters(ref *string) error {
	return nil
}

func (d *jsiiProxy_DescriptionBuilderSection) validateBuildIdParameters(prefix *string) error {
	return nil
}

func (d *jsiiProxy_DescriptionBuilderSection) validateRegisterBuilderParameters(builder IDescriptionBuilderComponent) error {
	return nil
}

func (d *jsiiProxy_DescriptionBuilderSection) validateRegisterChainableParameters(chainable awsstepfunctions.IChainable) error {
	return nil
}

func (d *jsiiProxy_DescriptionBuilderSection) validateSetDelimiterParameters(id *string, props *SetDelimiterProps) error {
	return nil
}

func (d *jsiiProxy_DescriptionBuilderSection) validateWriteParameters(id *string, props *WriteProps) error {
	return nil
}

func validateDescriptionBuilderSection_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewDescriptionBuilderSectionParameters(scope constructs.IConstruct, id *string, props *DescriptionBuilderSectionProps) error {
	return nil
}

