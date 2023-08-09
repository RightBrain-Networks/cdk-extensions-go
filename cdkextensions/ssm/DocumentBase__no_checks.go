//go:build no_runtime_type_checking

package ssm

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DocumentBase) validateAddRequirementParameters(requirement *DocumentRequirement) error {
	return nil
}

func (d *jsiiProxy_DocumentBase) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (d *jsiiProxy_DocumentBase) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (d *jsiiProxy_DocumentBase) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateDocumentBase_IsConstructParameters(x interface{}) error {
	return nil
}

func validateDocumentBase_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateDocumentBase_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewDocumentBaseParameters(scope constructs.IConstruct, id *string, props *DocumentBaseProps) error {
	return nil
}

