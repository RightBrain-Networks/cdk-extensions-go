package ec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"
)

// Provides options for specifying the remote side of a VPN connection.
type VpnConnectionRemoteEndpoint interface {
}

// The jsii proxy struct for VpnConnectionRemoteEndpoint
type jsiiProxy_VpnConnectionRemoteEndpoint struct {
	_ byte // padding
}

func NewVpnConnectionRemoteEndpoint() VpnConnectionRemoteEndpoint {
	_init_.Initialize()

	j := jsiiProxy_VpnConnectionRemoteEndpoint{}

	_jsii_.Create(
		"cdk-extensions.ec2.VpnConnectionRemoteEndpoint",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewVpnConnectionRemoteEndpoint_Override(v VpnConnectionRemoteEndpoint) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.ec2.VpnConnectionRemoteEndpoint",
		nil, // no parameters
		v,
	)
}

// Creates a remote connection using the configuration details provided.
//
// Returns: A configuration object representing a remote VPN destination.
func VpnConnectionRemoteEndpoint_FromConfiguration(configuration *CustomerGatewayProps) CustomerGatewayConfigurationRemoteVpnEndpoint {
	_init_.Initialize()

	if err := validateVpnConnectionRemoteEndpoint_FromConfigurationParameters(configuration); err != nil {
		panic(err)
	}
	var returns CustomerGatewayConfigurationRemoteVpnEndpoint

	_jsii_.StaticInvoke(
		"cdk-extensions.ec2.VpnConnectionRemoteEndpoint",
		"fromConfiguration",
		[]interface{}{configuration},
		&returns,
	)

	return returns
}

// Creates a remote connection using a customer gateway.
//
// Returns: A configuration object representing a remote VPN destination.
func VpnConnectionRemoteEndpoint_FromCustomerGateway(customerGateway ICustomerGateway) CustomerGatewayRemoteVpnEndpoint {
	_init_.Initialize()

	if err := validateVpnConnectionRemoteEndpoint_FromCustomerGatewayParameters(customerGateway); err != nil {
		panic(err)
	}
	var returns CustomerGatewayRemoteVpnEndpoint

	_jsii_.StaticInvoke(
		"cdk-extensions.ec2.VpnConnectionRemoteEndpoint",
		"fromCustomerGateway",
		[]interface{}{customerGateway},
		&returns,
	)

	return returns
}

