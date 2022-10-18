//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package sso

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PermissionSet) validateAddCustomerManagedPolicyParameters(options *ReferenceOptions) error {
	return nil
}

func (p *jsiiProxy_PermissionSet) validateAddManagedPolicyParameters(policy awsiam.IManagedPolicy) error {
	return nil
}

func (p *jsiiProxy_PermissionSet) validateAddToPrincipalPolicyParameters(statement awsiam.PolicyStatement) error {
	return nil
}

func (p *jsiiProxy_PermissionSet) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (p *jsiiProxy_PermissionSet) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (p *jsiiProxy_PermissionSet) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validatePermissionSet_FromArnParameters(scope constructs.Construct, id *string, arn *string) error {
	return nil
}

func validatePermissionSet_IsConstructParameters(x interface{}) error {
	return nil
}

func validatePermissionSet_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validatePermissionSet_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewPermissionSetParameters(scope constructs.Construct, id *string, props *PermissionSetProps) error {
	return nil
}

