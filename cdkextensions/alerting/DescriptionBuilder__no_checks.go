//go:build no_runtime_type_checking

package alerting

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DescriptionBuilder) validateAddIteratorParameters(id *string, props *DescriptionBuilderIteratorProps) error {
	return nil
}

func (d *jsiiProxy_DescriptionBuilder) validateAddReferenceParameters(id *string, props *AddReferenceProps) error {
	return nil
}

func (d *jsiiProxy_DescriptionBuilder) validateAddSectionParameters(id *string, props *DescriptionBuilderSectionProps) error {
	return nil
}

func (d *jsiiProxy_DescriptionBuilder) validateBuildIdParameters(prefix *string) error {
	return nil
}

func (d *jsiiProxy_DescriptionBuilder) validateRegisterBuilderParameters(builder IDescriptionBuilderComponent) error {
	return nil
}

func (d *jsiiProxy_DescriptionBuilder) validateRegisterChainableParameters(chainable awsstepfunctions.IChainable) error {
	return nil
}

func (d *jsiiProxy_DescriptionBuilder) validateSetDelimiterParameters(id *string, props *SetDelimiterProps) error {
	return nil
}

func (d *jsiiProxy_DescriptionBuilder) validateWriteParameters(id *string, props *WriteProps) error {
	return nil
}

func validateDescriptionBuilder_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewDescriptionBuilderParameters(scope constructs.IConstruct, props *DescriptionBuilderProps) error {
	return nil
}

