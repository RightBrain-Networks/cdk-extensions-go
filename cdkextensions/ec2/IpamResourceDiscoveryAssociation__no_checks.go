//go:build no_runtime_type_checking

package ec2

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IpamResourceDiscoveryAssociation) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (i *jsiiProxy_IpamResourceDiscoveryAssociation) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (i *jsiiProxy_IpamResourceDiscoveryAssociation) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateIpamResourceDiscoveryAssociation_IsConstructParameters(x interface{}) error {
	return nil
}

func validateIpamResourceDiscoveryAssociation_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateIpamResourceDiscoveryAssociation_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewIpamResourceDiscoveryAssociationParameters(scope constructs.IConstruct, id *string, props *IpamResourceDiscoveryAssociationProps) error {
	return nil
}

