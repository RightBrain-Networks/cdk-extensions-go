//go:build !no_runtime_type_checking

package ec2

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateVpnConnectionRemoteEndpoint_FromConfigurationParameters(configuration *CustomerGatewayProps) error {
	if configuration == nil {
		return fmt.Errorf("parameter configuration is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(configuration, func() string { return "parameter configuration" }); err != nil {
		return err
	}

	return nil
}

func validateVpnConnectionRemoteEndpoint_FromCustomerGatewayParameters(customerGateway ICustomerGateway) error {
	if customerGateway == nil {
		return fmt.Errorf("parameter customerGateway is required, but nil was provided")
	}

	return nil
}

