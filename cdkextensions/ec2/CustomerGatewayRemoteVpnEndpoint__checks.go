//go:build !no_runtime_type_checking

package ec2

import (
	"fmt"

	"github.com/aws/constructs-go/constructs/v10"
)

func (c *jsiiProxy_CustomerGatewayRemoteVpnEndpoint) validateBindParameters(_scope constructs.IConstruct) error {
	if _scope == nil {
		return fmt.Errorf("parameter _scope is required, but nil was provided")
	}

	return nil
}

func validateNewCustomerGatewayRemoteVpnEndpointParameters(customerGateway ICustomerGateway) error {
	if customerGateway == nil {
		return fmt.Errorf("parameter customerGateway is required, but nil was provided")
	}

	return nil
}

