//go:build no_runtime_type_checking

package ec2

// Building without runtime type checking enabled, so all the below just return nil

func validateVpnConnectionRemoteEndpoint_FromConfigurationParameters(configuration *CustomerGatewayProps) error {
	return nil
}

func validateVpnConnectionRemoteEndpoint_FromCustomerGatewayParameters(customerGateway ICustomerGateway) error {
	return nil
}

