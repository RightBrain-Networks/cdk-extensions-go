//go:build no_runtime_type_checking

package ec2

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IpamPoolCidr) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (i *jsiiProxy_IpamPoolCidr) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (i *jsiiProxy_IpamPoolCidr) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateIpamPoolCidr_IsConstructParameters(x interface{}) error {
	return nil
}

func validateIpamPoolCidr_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateIpamPoolCidr_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewIpamPoolCidrParameters(scope constructs.IConstruct, id *string, props *IpamPoolCidrProps) error {
	return nil
}

