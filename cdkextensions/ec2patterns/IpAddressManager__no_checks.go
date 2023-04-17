//go:build no_runtime_type_checking

package ec2patterns

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IpAddressManager) validateAddRegionParameters(region *string) error {
	return nil
}

func (i *jsiiProxy_IpAddressManager) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (i *jsiiProxy_IpAddressManager) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (i *jsiiProxy_IpAddressManager) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (i *jsiiProxy_IpAddressManager) validateGetVpcConfigurationParameters(scope constructs.IConstruct, id *string, options *GetVpcConfigurationOptions) error {
	return nil
}

func (i *jsiiProxy_IpAddressManager) validatePrivateVpcPoolForEnvironmentParameters(account *string, region *string) error {
	return nil
}

func (i *jsiiProxy_IpAddressManager) validatePrivateVpcPoolForRegionParameters(region *string) error {
	return nil
}

func (i *jsiiProxy_IpAddressManager) validateRegisterAccountParameters(account *string, pool ec2.IIpamPool) error {
	return nil
}

func (i *jsiiProxy_IpAddressManager) validateRegisterCidrParameters(scope constructs.IConstruct, id *string, cidr *string) error {
	return nil
}

func validateIpAddressManager_IsConstructParameters(x interface{}) error {
	return nil
}

func validateIpAddressManager_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateIpAddressManager_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewIpAddressManagerParameters(scope constructs.IConstruct, id *string, props *IpAddressManagerProps) error {
	return nil
}

