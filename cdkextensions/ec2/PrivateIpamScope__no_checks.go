//go:build no_runtime_type_checking

package ec2

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PrivateIpamScope) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (p *jsiiProxy_PrivateIpamScope) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (p *jsiiProxy_PrivateIpamScope) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validatePrivateIpamScope_FromIpamScopeArnParameters(scope constructs.IConstruct, id *string, ipamScopeArn *string) error {
	return nil
}

func validatePrivateIpamScope_FromIpamScopeAttributesParameters(scope constructs.IConstruct, id *string, attrs *IpamScopeAttributes) error {
	return nil
}

func validatePrivateIpamScope_FromIpamScopeIdParameters(scope constructs.IConstruct, id *string, ipamScopeId *string) error {
	return nil
}

func validatePrivateIpamScope_IsConstructParameters(x interface{}) error {
	return nil
}

func validatePrivateIpamScope_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validatePrivateIpamScope_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewPrivateIpamScopeParameters(scope constructs.IConstruct, id *string, props *PrivateIpamScopeProps) error {
	return nil
}

