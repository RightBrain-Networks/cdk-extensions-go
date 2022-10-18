//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package sso

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_GroupBase) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (g *jsiiProxy_GroupBase) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (g *jsiiProxy_GroupBase) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateGroupBase_IsConstructParameters(x interface{}) error {
	return nil
}

func validateGroupBase_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateGroupBase_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewGroupBaseParameters(scope constructs.IConstruct, id *string, props *awscdk.ResourceProps) error {
	return nil
}

