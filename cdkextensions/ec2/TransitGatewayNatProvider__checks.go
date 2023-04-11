//go:build !no_runtime_type_checking

package ec2

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

func (t *jsiiProxy_TransitGatewayNatProvider) validateConfigureNatParameters(options *awsec2.ConfigureNatOptions) error {
	if options == nil {
		return fmt.Errorf("parameter options is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (t *jsiiProxy_TransitGatewayNatProvider) validateConfigureSubnetParameters(subnet awsec2.PrivateSubnet) error {
	if subnet == nil {
		return fmt.Errorf("parameter subnet is required, but nil was provided")
	}

	return nil
}

func validateTransitGatewayNatProvider_GatewayParameters(props *awsec2.NatGatewayProps) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func validateTransitGatewayNatProvider_InstanceParameters(props *awsec2.NatInstanceProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func validateNewTransitGatewayNatProviderParameters(options *TransitGatewayNatProviderOptions) error {
	if options == nil {
		return fmt.Errorf("parameter options is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

