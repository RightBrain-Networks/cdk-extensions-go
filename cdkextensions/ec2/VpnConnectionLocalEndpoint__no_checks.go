//go:build no_runtime_type_checking

package ec2

// Building without runtime type checking enabled, so all the below just return nil

func validateVpnConnectionLocalEndpoint_FromTransitGatewayParameters(transitGateway ITransitGateway) error {
	return nil
}

func validateVpnConnectionLocalEndpoint_FromVpnGatewayParameters(vpnGateway awsec2.IVpnGateway) error {
	return nil
}

