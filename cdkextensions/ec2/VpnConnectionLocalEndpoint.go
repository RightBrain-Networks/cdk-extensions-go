package ec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

// Provides options for specifying the local side of a VPN connection.
type VpnConnectionLocalEndpoint interface {
}

// The jsii proxy struct for VpnConnectionLocalEndpoint
type jsiiProxy_VpnConnectionLocalEndpoint struct {
	_ byte // padding
}

func NewVpnConnectionLocalEndpoint() VpnConnectionLocalEndpoint {
	_init_.Initialize()

	j := jsiiProxy_VpnConnectionLocalEndpoint{}

	_jsii_.Create(
		"cdk-extensions.ec2.VpnConnectionLocalEndpoint",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewVpnConnectionLocalEndpoint_Override(v VpnConnectionLocalEndpoint) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.ec2.VpnConnectionLocalEndpoint",
		nil, // no parameters
		v,
	)
}

func VpnConnectionLocalEndpoint_FromTransitGateway(transitGateway ITransitGateway) TransitGatewayLocalVpnEndpoint {
	_init_.Initialize()

	if err := validateVpnConnectionLocalEndpoint_FromTransitGatewayParameters(transitGateway); err != nil {
		panic(err)
	}
	var returns TransitGatewayLocalVpnEndpoint

	_jsii_.StaticInvoke(
		"cdk-extensions.ec2.VpnConnectionLocalEndpoint",
		"fromTransitGateway",
		[]interface{}{transitGateway},
		&returns,
	)

	return returns
}

func VpnConnectionLocalEndpoint_FromVpnGateway(vpnGateway awsec2.IVpnGateway) VpnGatewayLocalVpnEndpoint {
	_init_.Initialize()

	if err := validateVpnConnectionLocalEndpoint_FromVpnGatewayParameters(vpnGateway); err != nil {
		panic(err)
	}
	var returns VpnGatewayLocalVpnEndpoint

	_jsii_.StaticInvoke(
		"cdk-extensions.ec2.VpnConnectionLocalEndpoint",
		"fromVpnGateway",
		[]interface{}{vpnGateway},
		&returns,
	)

	return returns
}

