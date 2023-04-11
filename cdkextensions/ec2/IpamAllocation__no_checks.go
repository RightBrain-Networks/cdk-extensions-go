//go:build no_runtime_type_checking

package ec2

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IpamAllocation) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (i *jsiiProxy_IpamAllocation) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (i *jsiiProxy_IpamAllocation) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateIpamAllocation_IsConstructParameters(x interface{}) error {
	return nil
}

func validateIpamAllocation_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateIpamAllocation_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewIpamAllocationParameters(scope constructs.IConstruct, id *string, props *IpamAllocationProps) error {
	return nil
}

