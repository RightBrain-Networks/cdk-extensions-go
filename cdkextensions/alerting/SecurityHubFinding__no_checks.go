//go:build no_runtime_type_checking

package alerting

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SecurityHubFinding) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (s *jsiiProxy_SecurityHubFinding) validateBindParameters(_scope constructs.IConstruct) error {
	return nil
}

func (s *jsiiProxy_SecurityHubFinding) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (s *jsiiProxy_SecurityHubFinding) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (s *jsiiProxy_SecurityHubFinding) validateRegisterIssueTriggerParameters(id *string, options *SecurityHubFindingEventOptions) error {
	return nil
}

func validateSecurityHubFinding_IsConstructParameters(x interface{}) error {
	return nil
}

func validateSecurityHubFinding_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateSecurityHubFinding_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewSecurityHubFindingParameters(scope constructs.IConstruct, id *string, props *SecurityHubFindingProps) error {
	return nil
}

