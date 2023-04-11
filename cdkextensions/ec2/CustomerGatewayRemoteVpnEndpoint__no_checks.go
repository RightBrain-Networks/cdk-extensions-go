//go:build no_runtime_type_checking

package ec2

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CustomerGatewayRemoteVpnEndpoint) validateBindParameters(_scope constructs.IConstruct) error {
	return nil
}

func validateNewCustomerGatewayRemoteVpnEndpointParameters(customerGateway ICustomerGateway) error {
	return nil
}

