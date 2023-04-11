//go:build no_runtime_type_checking

package sso

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_InstanceAccessControlAttributeConfiguration) validateAddAttributeParameters(key *string) error {
	return nil
}

func (i *jsiiProxy_InstanceAccessControlAttributeConfiguration) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (i *jsiiProxy_InstanceAccessControlAttributeConfiguration) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (i *jsiiProxy_InstanceAccessControlAttributeConfiguration) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateInstanceAccessControlAttributeConfiguration_IsConstructParameters(x interface{}) error {
	return nil
}

func validateInstanceAccessControlAttributeConfiguration_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateInstanceAccessControlAttributeConfiguration_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewInstanceAccessControlAttributeConfigurationParameters(scope constructs.Construct, id *string, props *InstanceAccessControlAttributeConfigurationProps) error {
	return nil
}

