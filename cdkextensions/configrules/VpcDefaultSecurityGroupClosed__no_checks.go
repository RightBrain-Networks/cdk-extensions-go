//go:build no_runtime_type_checking

package configrules

// Building without runtime type checking enabled, so all the below just return nil

func (v *jsiiProxy_VpcDefaultSecurityGroupClosed) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (v *jsiiProxy_VpcDefaultSecurityGroupClosed) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (v *jsiiProxy_VpcDefaultSecurityGroupClosed) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (v *jsiiProxy_VpcDefaultSecurityGroupClosed) validateOnComplianceChangeParameters(id *string, options *awsevents.OnEventOptions) error {
	return nil
}

func (v *jsiiProxy_VpcDefaultSecurityGroupClosed) validateOnEventParameters(id *string, options *awsevents.OnEventOptions) error {
	return nil
}

func (v *jsiiProxy_VpcDefaultSecurityGroupClosed) validateOnReEvaluationStatusParameters(id *string, options *awsevents.OnEventOptions) error {
	return nil
}

func validateVpcDefaultSecurityGroupClosed_FromConfigRuleNameParameters(scope constructs.Construct, id *string, configRuleName *string) error {
	return nil
}

func validateVpcDefaultSecurityGroupClosed_IsConstructParameters(x interface{}) error {
	return nil
}

func validateVpcDefaultSecurityGroupClosed_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateVpcDefaultSecurityGroupClosed_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewVpcDefaultSecurityGroupClosedParameters(scope constructs.IConstruct, id *string, props *VpcDefaultSecurityGroupClosedProps) error {
	return nil
}

