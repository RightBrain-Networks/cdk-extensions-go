//go:build no_runtime_type_checking

package resourcegroups

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_Group) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (g *jsiiProxy_Group) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (g *jsiiProxy_Group) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateGroup_FromGroupArnParameters(scope constructs.IConstruct, id *string, arn *string) error {
	return nil
}

func validateGroup_FromGroupAttributesParameters(scope constructs.IConstruct, id *string, attrs *GroupAttributes) error {
	return nil
}

func validateGroup_FromGroupNameParameters(scope constructs.IConstruct, id *string, name *string) error {
	return nil
}

func validateGroup_IsConstructParameters(x interface{}) error {
	return nil
}

func validateGroup_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateGroup_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewGroupParameters(scope constructs.IConstruct, id *string, props *GroupProps) error {
	return nil
}

