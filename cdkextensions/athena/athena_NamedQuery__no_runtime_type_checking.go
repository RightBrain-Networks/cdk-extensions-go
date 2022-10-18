//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package athena

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_NamedQuery) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (n *jsiiProxy_NamedQuery) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (n *jsiiProxy_NamedQuery) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateNamedQuery_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNamedQuery_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNamedQuery_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewNamedQueryParameters(scope constructs.Construct, id *string, props *NamedQueryProps) error {
	return nil
}

