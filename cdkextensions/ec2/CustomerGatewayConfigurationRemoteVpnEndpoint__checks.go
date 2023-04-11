//go:build !no_runtime_type_checking

package ec2

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

func (c *jsiiProxy_CustomerGatewayConfigurationRemoteVpnEndpoint) validateBindParameters(scope constructs.IConstruct) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	return nil
}

func validateNewCustomerGatewayConfigurationRemoteVpnEndpointParameters(configuration *CustomerGatewayProps) error {
	if configuration == nil {
		return fmt.Errorf("parameter configuration is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(configuration, func() string { return "parameter configuration" }); err != nil {
		return err
	}

	return nil
}

