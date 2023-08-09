//go:build no_runtime_type_checking

package configrules

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IamPasswordPolicy) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (i *jsiiProxy_IamPasswordPolicy) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (i *jsiiProxy_IamPasswordPolicy) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (i *jsiiProxy_IamPasswordPolicy) validateOnComplianceChangeParameters(id *string, options *awsevents.OnEventOptions) error {
	return nil
}

func (i *jsiiProxy_IamPasswordPolicy) validateOnEventParameters(id *string, options *awsevents.OnEventOptions) error {
	return nil
}

func (i *jsiiProxy_IamPasswordPolicy) validateOnReEvaluationStatusParameters(id *string, options *awsevents.OnEventOptions) error {
	return nil
}

func validateIamPasswordPolicy_FromConfigRuleNameParameters(scope constructs.Construct, id *string, configRuleName *string) error {
	return nil
}

func validateIamPasswordPolicy_IsConstructParameters(x interface{}) error {
	return nil
}

func validateIamPasswordPolicy_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateIamPasswordPolicy_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewIamPasswordPolicyParameters(scope constructs.IConstruct, id *string, props *IamPasswordPolicyProps) error {
	return nil
}

