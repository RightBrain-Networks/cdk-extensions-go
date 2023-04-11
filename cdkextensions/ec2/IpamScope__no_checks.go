//go:build no_runtime_type_checking

package ec2

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IpamScope) validateAddPoolParameters(id *string, options *IpamPoolOptions) error {
	return nil
}

func (i *jsiiProxy_IpamScope) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (i *jsiiProxy_IpamScope) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (i *jsiiProxy_IpamScope) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateIpamScope_FromIpamScopeArnParameters(scope constructs.IConstruct, id *string, ipamScopeArn *string) error {
	return nil
}

func validateIpamScope_FromIpamScopeAttributesParameters(scope constructs.IConstruct, id *string, attrs *IpamScopeAttributes) error {
	return nil
}

func validateIpamScope_FromIpamScopeIdParameters(scope constructs.IConstruct, id *string, ipamScopeId *string) error {
	return nil
}

func validateIpamScope_IsConstructParameters(x interface{}) error {
	return nil
}

func validateIpamScope_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateIpamScope_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewIpamScopeParameters(scope constructs.IConstruct, id *string, props *IpamScopeProps) error {
	return nil
}

