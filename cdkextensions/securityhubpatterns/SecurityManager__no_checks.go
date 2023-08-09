//go:build no_runtime_type_checking

package securityhubpatterns

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SecurityManager) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (s *jsiiProxy_SecurityManager) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (s *jsiiProxy_SecurityManager) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateSecurityManager_IsConstructParameters(x interface{}) error {
	return nil
}

func validateSecurityManager_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateSecurityManager_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewSecurityManagerParameters(scope constructs.IConstruct, id *string, props *SecurityManagerProps) error {
	return nil
}

