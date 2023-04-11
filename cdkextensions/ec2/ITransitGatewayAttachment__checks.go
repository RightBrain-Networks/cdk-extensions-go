//go:build !no_runtime_type_checking

package ec2

import (
	"fmt"
)

func (i *jsiiProxy_ITransitGatewayAttachment) validateAddRouteParameters(cidr *string, routeTable ITransitGatewayRouteTable) error {
	if cidr == nil {
		return fmt.Errorf("parameter cidr is required, but nil was provided")
	}

	if routeTable == nil {
		return fmt.Errorf("parameter routeTable is required, but nil was provided")
	}

	return nil
}

