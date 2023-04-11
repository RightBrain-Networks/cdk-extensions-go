//go:build no_runtime_type_checking

package ec2

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IpamPool) validateAddChildPoolParameters(id *string, options *AddChildPoolOptions) error {
	return nil
}

func (i *jsiiProxy_IpamPool) validateAddCidrToPoolParameters(id *string, options *AddCidrToPoolOptions) error {
	return nil
}

func (i *jsiiProxy_IpamPool) validateAddTagRestrictionParameters(key *string, value *string) error {
	return nil
}

func (i *jsiiProxy_IpamPool) validateAllocateCidrFromPoolParameters(id *string, options *AllocateCidrFromPoolOptions) error {
	return nil
}

func (i *jsiiProxy_IpamPool) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (i *jsiiProxy_IpamPool) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (i *jsiiProxy_IpamPool) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateIpamPool_IsConstructParameters(x interface{}) error {
	return nil
}

func validateIpamPool_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateIpamPool_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewIpamPoolParameters(scope constructs.IConstruct, id *string, props *IpamPoolProps) error {
	return nil
}

