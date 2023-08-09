//go:build no_runtime_type_checking

package config

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RemediationConfiguration) validateAddParameterParameters(key *string) error {
	return nil
}

func (r *jsiiProxy_RemediationConfiguration) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (r *jsiiProxy_RemediationConfiguration) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (r *jsiiProxy_RemediationConfiguration) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateRemediationConfiguration_FromRemediationConfigurationArnParameters(scope constructs.IConstruct, id *string, arn *string) error {
	return nil
}

func validateRemediationConfiguration_FromRemediationConfigurationAttributesParameters(scope constructs.IConstruct, id *string, attrs *RemediationConfigurationAttributes) error {
	return nil
}

func validateRemediationConfiguration_FromRemediationConfigurationNameParameters(scope constructs.IConstruct, id *string, name *string) error {
	return nil
}

func validateRemediationConfiguration_IsConstructParameters(x interface{}) error {
	return nil
}

func validateRemediationConfiguration_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateRemediationConfiguration_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewRemediationConfigurationParameters(scope constructs.IConstruct, id *string, props *RemediationConfigurationProps) error {
	return nil
}

