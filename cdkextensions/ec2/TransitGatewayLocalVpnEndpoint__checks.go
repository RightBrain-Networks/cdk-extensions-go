//go:build !no_runtime_type_checking

package ec2

import (
	"fmt"

	"github.com/aws/constructs-go/constructs/v10"
)

func (t *jsiiProxy_TransitGatewayLocalVpnEndpoint) validateBindParameters(_scope constructs.IConstruct) error {
	if _scope == nil {
		return fmt.Errorf("parameter _scope is required, but nil was provided")
	}

	return nil
}

func validateNewTransitGatewayLocalVpnEndpointParameters(transitGateway ITransitGateway) error {
	if transitGateway == nil {
		return fmt.Errorf("parameter transitGateway is required, but nil was provided")
	}

	return nil
}

