//go:build no_runtime_type_checking

package ec2

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TieredSubnets) validateAllocateSubnetsCidrParameters(input *awsec2.AllocateCidrRequest) error {
	return nil
}

func validateNewTieredSubnetsParameters(options *TieredSubnetsOptions) error {
	return nil
}

