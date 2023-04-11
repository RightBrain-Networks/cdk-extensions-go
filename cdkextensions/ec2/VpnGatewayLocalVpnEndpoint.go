package ec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/constructs-go/constructs/v10"
)

// Specifies a VPN connection endpoint which routes to a VPN gateway on the AWS side.
type VpnGatewayLocalVpnEndpoint interface {
	ILocalVpnEndpoint
	// The VPN gateway that serves as the local end of a VPN connection.
	VpnGateway() awsec2.IVpnGateway
	// Produces a configuration that can be used when configuring the local end of a VPN connection.
	Bind(_scope constructs.IConstruct) *LocalVpnEndpointConfiguration
}

// The jsii proxy struct for VpnGatewayLocalVpnEndpoint
type jsiiProxy_VpnGatewayLocalVpnEndpoint struct {
	jsiiProxy_ILocalVpnEndpoint
}

func (j *jsiiProxy_VpnGatewayLocalVpnEndpoint) VpnGateway() awsec2.IVpnGateway {
	var returns awsec2.IVpnGateway
	_jsii_.Get(
		j,
		"vpnGateway",
		&returns,
	)
	return returns
}


// Creates a new instance of the VpnGatewayLocalVpnEndpoint class.
func NewVpnGatewayLocalVpnEndpoint(vpnGateway awsec2.IVpnGateway) VpnGatewayLocalVpnEndpoint {
	_init_.Initialize()

	if err := validateNewVpnGatewayLocalVpnEndpointParameters(vpnGateway); err != nil {
		panic(err)
	}
	j := jsiiProxy_VpnGatewayLocalVpnEndpoint{}

	_jsii_.Create(
		"cdk-extensions.ec2.VpnGatewayLocalVpnEndpoint",
		[]interface{}{vpnGateway},
		&j,
	)

	return &j
}

// Creates a new instance of the VpnGatewayLocalVpnEndpoint class.
func NewVpnGatewayLocalVpnEndpoint_Override(v VpnGatewayLocalVpnEndpoint, vpnGateway awsec2.IVpnGateway) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.ec2.VpnGatewayLocalVpnEndpoint",
		[]interface{}{vpnGateway},
		v,
	)
}

func (v *jsiiProxy_VpnGatewayLocalVpnEndpoint) Bind(_scope constructs.IConstruct) *LocalVpnEndpointConfiguration {
	if err := v.validateBindParameters(_scope); err != nil {
		panic(err)
	}
	var returns *LocalVpnEndpointConfiguration

	_jsii_.Invoke(
		v,
		"bind",
		[]interface{}{_scope},
		&returns,
	)

	return returns
}

