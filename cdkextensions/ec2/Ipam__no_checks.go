//go:build no_runtime_type_checking

package ec2

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_Ipam) validateAddRegionParameters(region *string) error {
	return nil
}

func (i *jsiiProxy_Ipam) validateAddScopeParameters(id *string, options *IpamScopeOptions) error {
	return nil
}

func (i *jsiiProxy_Ipam) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (i *jsiiProxy_Ipam) validateAssociateResourceDiscoveryParameters(resourceDiscovery IIpamResourceDiscovery) error {
	return nil
}

func (i *jsiiProxy_Ipam) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (i *jsiiProxy_Ipam) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateIpam_FromIpamArnParameters(scope constructs.IConstruct, id *string, ipamArn *string) error {
	return nil
}

func validateIpam_FromIpamAttributesParameters(scope constructs.IConstruct, id *string, attrs *IpamAttributes) error {
	return nil
}

func validateIpam_FromIpamIdParameters(scope constructs.IConstruct, id *string, ipamId *string) error {
	return nil
}

func validateIpam_IsConstructParameters(x interface{}) error {
	return nil
}

func validateIpam_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateIpam_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewIpamParameters(scope constructs.IConstruct, id *string, props *IpamProps) error {
	return nil
}

