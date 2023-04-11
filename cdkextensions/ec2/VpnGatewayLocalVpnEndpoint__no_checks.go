//go:build no_runtime_type_checking

package ec2

// Building without runtime type checking enabled, so all the below just return nil

func (v *jsiiProxy_VpnGatewayLocalVpnEndpoint) validateBindParameters(_scope constructs.IConstruct) error {
	return nil
}

func validateNewVpnGatewayLocalVpnEndpointParameters(vpnGateway awsec2.IVpnGateway) error {
	return nil
}

