//go:build no_runtime_type_checking

package ec2

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IIpamPool) validateAddChildPoolParameters(id *string, options *AddChildPoolOptions) error {
	return nil
}

func (i *jsiiProxy_IIpamPool) validateAddCidrToPoolParameters(id *string, options *AddCidrToPoolOptions) error {
	return nil
}

func (i *jsiiProxy_IIpamPool) validateAllocateCidrFromPoolParameters(id *string, options *AllocateCidrFromPoolOptions) error {
	return nil
}

