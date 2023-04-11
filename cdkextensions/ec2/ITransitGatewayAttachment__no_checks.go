//go:build no_runtime_type_checking

package ec2

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_ITransitGatewayAttachment) validateAddRouteParameters(cidr *string, routeTable ITransitGatewayRouteTable) error {
	return nil
}

