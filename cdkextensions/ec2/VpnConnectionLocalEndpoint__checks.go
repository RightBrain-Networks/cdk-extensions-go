//go:build !no_runtime_type_checking

package ec2

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

func validateVpnConnectionLocalEndpoint_FromTransitGatewayParameters(transitGateway ITransitGateway) error {
	if transitGateway == nil {
		return fmt.Errorf("parameter transitGateway is required, but nil was provided")
	}

	return nil
}

func validateVpnConnectionLocalEndpoint_FromVpnGatewayParameters(vpnGateway awsec2.IVpnGateway) error {
	if vpnGateway == nil {
		return fmt.Errorf("parameter vpnGateway is required, but nil was provided")
	}

	return nil
}

