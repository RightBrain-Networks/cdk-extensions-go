//go:build no_runtime_type_checking

package ec2

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IIpamResourceDiscovery) validateAddIpamParameters(id *string, options *IpamProps) error {
	return nil
}

func (i *jsiiProxy_IIpamResourceDiscovery) validateAssociateIpamParameters(ipam IIpam) error {
	return nil
}

