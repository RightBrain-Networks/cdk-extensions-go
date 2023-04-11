//go:build no_runtime_type_checking

package ec2

// Building without runtime type checking enabled, so all the below just return nil

func validateNatProvider_GatewayParameters(props *awsec2.NatGatewayProps) error {
	return nil
}

func validateNatProvider_InstanceParameters(props *awsec2.NatInstanceProps) error {
	return nil
}

func validateNatProvider_TransitGatewayParameters(props *TransitGatewayNatProviderOptions) error {
	return nil
}

