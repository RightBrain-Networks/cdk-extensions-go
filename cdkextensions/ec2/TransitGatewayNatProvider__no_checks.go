//go:build no_runtime_type_checking

package ec2

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TransitGatewayNatProvider) validateConfigureNatParameters(options *awsec2.ConfigureNatOptions) error {
	return nil
}

func (t *jsiiProxy_TransitGatewayNatProvider) validateConfigureSubnetParameters(subnet awsec2.PrivateSubnet) error {
	return nil
}

func validateTransitGatewayNatProvider_GatewayParameters(props *awsec2.NatGatewayProps) error {
	return nil
}

func validateTransitGatewayNatProvider_InstanceParameters(props *awsec2.NatInstanceProps) error {
	return nil
}

func validateNewTransitGatewayNatProviderParameters(options *TransitGatewayNatProviderOptions) error {
	return nil
}

