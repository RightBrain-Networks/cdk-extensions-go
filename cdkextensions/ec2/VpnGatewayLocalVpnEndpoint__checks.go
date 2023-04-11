//go:build !no_runtime_type_checking

package ec2

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/constructs-go/constructs/v10"
)

func (v *jsiiProxy_VpnGatewayLocalVpnEndpoint) validateBindParameters(_scope constructs.IConstruct) error {
	if _scope == nil {
		return fmt.Errorf("parameter _scope is required, but nil was provided")
	}

	return nil
}

func validateNewVpnGatewayLocalVpnEndpointParameters(vpnGateway awsec2.IVpnGateway) error {
	if vpnGateway == nil {
		return fmt.Errorf("parameter vpnGateway is required, but nil was provided")
	}

	return nil
}

