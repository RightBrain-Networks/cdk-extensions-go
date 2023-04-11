//go:build no_runtime_type_checking

package ec2

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IIpamScope) validateAddPoolParameters(id *string, options *IpamPoolOptions) error {
	return nil
}

