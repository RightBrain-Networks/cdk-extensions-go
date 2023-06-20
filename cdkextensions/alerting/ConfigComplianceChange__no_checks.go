//go:build no_runtime_type_checking

package alerting

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ConfigComplianceChange) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (c *jsiiProxy_ConfigComplianceChange) validateBindParameters(_node constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_ConfigComplianceChange) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (c *jsiiProxy_ConfigComplianceChange) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (c *jsiiProxy_ConfigComplianceChange) validateRegisterIssueTriggerParameters(id *string, options *ConfigComplianceChangeRuleOptions) error {
	return nil
}

func validateConfigComplianceChange_IsConstructParameters(x interface{}) error {
	return nil
}

func validateConfigComplianceChange_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateConfigComplianceChange_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewConfigComplianceChangeParameters(scope constructs.IConstruct, id *string, props *ConfigComplianceChangeProps) error {
	return nil
}

