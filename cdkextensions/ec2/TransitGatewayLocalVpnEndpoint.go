package ec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/constructs-go/constructs/v10"
)

// Specifies a VPN connection endpoint which routes to a transit gateway on the AWS side.
type TransitGatewayLocalVpnEndpoint interface {
	ILocalVpnEndpoint
	// The transit gateway that serves as the local end of a VPN connection.
	TransitGateway() ITransitGateway
	// Produces a configuration that can be used when configuring the local end of a VPN connection.
	Bind(_scope constructs.IConstruct) *LocalVpnEndpointConfiguration
}

// The jsii proxy struct for TransitGatewayLocalVpnEndpoint
type jsiiProxy_TransitGatewayLocalVpnEndpoint struct {
	jsiiProxy_ILocalVpnEndpoint
}

func (j *jsiiProxy_TransitGatewayLocalVpnEndpoint) TransitGateway() ITransitGateway {
	var returns ITransitGateway
	_jsii_.Get(
		j,
		"transitGateway",
		&returns,
	)
	return returns
}


// Creates a new instance of the TransitGatewayLocalVpnEndpoint class.
func NewTransitGatewayLocalVpnEndpoint(transitGateway ITransitGateway) TransitGatewayLocalVpnEndpoint {
	_init_.Initialize()

	if err := validateNewTransitGatewayLocalVpnEndpointParameters(transitGateway); err != nil {
		panic(err)
	}
	j := jsiiProxy_TransitGatewayLocalVpnEndpoint{}

	_jsii_.Create(
		"cdk-extensions.ec2.TransitGatewayLocalVpnEndpoint",
		[]interface{}{transitGateway},
		&j,
	)

	return &j
}

// Creates a new instance of the TransitGatewayLocalVpnEndpoint class.
func NewTransitGatewayLocalVpnEndpoint_Override(t TransitGatewayLocalVpnEndpoint, transitGateway ITransitGateway) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.ec2.TransitGatewayLocalVpnEndpoint",
		[]interface{}{transitGateway},
		t,
	)
}

func (t *jsiiProxy_TransitGatewayLocalVpnEndpoint) Bind(_scope constructs.IConstruct) *LocalVpnEndpointConfiguration {
	if err := t.validateBindParameters(_scope); err != nil {
		panic(err)
	}
	var returns *LocalVpnEndpointConfiguration

	_jsii_.Invoke(
		t,
		"bind",
		[]interface{}{_scope},
		&returns,
	)

	return returns
}

