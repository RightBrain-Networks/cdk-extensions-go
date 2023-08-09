//go:build no_runtime_type_checking

package ssm

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_Document) validateAddRequirementParameters(requirement *DocumentRequirement) error {
	return nil
}

func (d *jsiiProxy_Document) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (d *jsiiProxy_Document) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (d *jsiiProxy_Document) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateDocument_IsConstructParameters(x interface{}) error {
	return nil
}

func validateDocument_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateDocument_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewDocumentParameters(scope constructs.IConstruct, id *string, props *DocumentProps) error {
	return nil
}

