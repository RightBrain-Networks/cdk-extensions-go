//go:build no_runtime_type_checking

package ec2

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TransitGatewayLocalVpnEndpoint) validateBindParameters(_scope constructs.IConstruct) error {
	return nil
}

func validateNewTransitGatewayLocalVpnEndpointParameters(transitGateway ITransitGateway) error {
	return nil
}

