//go:build no_runtime_type_checking

package ec2

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IpamResourceDiscovery) validateAddIpamParameters(id *string, options *IpamProps) error {
	return nil
}

func (i *jsiiProxy_IpamResourceDiscovery) validateAddRegionParameters(region *string) error {
	return nil
}

func (i *jsiiProxy_IpamResourceDiscovery) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (i *jsiiProxy_IpamResourceDiscovery) validateAssociateIpamParameters(ipam IIpam) error {
	return nil
}

func (i *jsiiProxy_IpamResourceDiscovery) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (i *jsiiProxy_IpamResourceDiscovery) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateIpamResourceDiscovery_FromIpamResourceDiscoveryArnParameters(scope constructs.IConstruct, id *string, ipamResourceDiscoveryArn *string) error {
	return nil
}

func validateIpamResourceDiscovery_FromIpamResourceDiscoveryAttributesParameters(scope constructs.IConstruct, id *string, attrs *IpamResourceDiscoveryAttributes) error {
	return nil
}

func validateIpamResourceDiscovery_FromIpamResourceDiscoveryIdParameters(scope constructs.IConstruct, id *string, ipamResourceDiscoveryId *string) error {
	return nil
}

func validateIpamResourceDiscovery_IsConstructParameters(x interface{}) error {
	return nil
}

func validateIpamResourceDiscovery_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateIpamResourceDiscovery_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewIpamResourceDiscoveryParameters(scope constructs.IConstruct, id *string, props *IpamResourceDiscoveryProps) error {
	return nil
}

