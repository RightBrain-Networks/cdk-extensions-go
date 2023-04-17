//go:build no_runtime_type_checking

package ec2

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IPublicIpamScope) validateAddAwsProvidedIpv6PoolParameters(id *string, options *AddAwsProvidedIpv6PoolOptions) error {
	return nil
}

func (i *jsiiProxy_IPublicIpamScope) validateAddByoipIpv4PoolParameters(id *string, options *AddByoipIpv4PoolOptions) error {
	return nil
}

func (i *jsiiProxy_IPublicIpamScope) validateAddByoipIpv6PoolParameters(id *string, options *AddByoipIpv6PoolOptions) error {
	return nil
}

