//go:build no_runtime_type_checking

package ec2

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IIpam) validateAddScopeParameters(id *string, options *PrivateIpamScopeOptions) error {
	return nil
}

func (i *jsiiProxy_IIpam) validateAssociateResourceDiscoveryParameters(resourceDiscovery IIpamResourceDiscovery) error {
	return nil
}

