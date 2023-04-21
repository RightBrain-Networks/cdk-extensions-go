//go:build !no_runtime_type_checking

package ec2

import (
	"fmt"
)

func (i *jsiiProxy_ITransitGatewayAttachment) validateAddRouteParameters(id *string, cidr *string, routeTable ITransitGatewayRouteTable) error {
	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if cidr == nil {
		return fmt.Errorf("parameter cidr is required, but nil was provided")
	}

	if routeTable == nil {
		return fmt.Errorf("parameter routeTable is required, but nil was provided")
	}

	return nil
}

