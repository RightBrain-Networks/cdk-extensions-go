package ec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/constructs-go/constructs/v10"
)

// Specifies a remote VPN endpoint device that has its details configured in an existing customer gateway.
type CustomerGatewayRemoteVpnEndpoint interface {
	IRemoteVpnEndpoint
	// The customer gateway that is configured with the details of the remote endpoint device.
	CustomerGateway() ICustomerGateway
	// Produces a configuration that can be used when configuring the remote end of a VPN connection.
	Bind(_scope constructs.IConstruct) *RemoteVpnEndpointConfiguration
}

// The jsii proxy struct for CustomerGatewayRemoteVpnEndpoint
type jsiiProxy_CustomerGatewayRemoteVpnEndpoint struct {
	jsiiProxy_IRemoteVpnEndpoint
}

func (j *jsiiProxy_CustomerGatewayRemoteVpnEndpoint) CustomerGateway() ICustomerGateway {
	var returns ICustomerGateway
	_jsii_.Get(
		j,
		"customerGateway",
		&returns,
	)
	return returns
}


// Creates a new instance of the CustomerGatewayRemoteVpnEndpoint class.
func NewCustomerGatewayRemoteVpnEndpoint(customerGateway ICustomerGateway) CustomerGatewayRemoteVpnEndpoint {
	_init_.Initialize()

	if err := validateNewCustomerGatewayRemoteVpnEndpointParameters(customerGateway); err != nil {
		panic(err)
	}
	j := jsiiProxy_CustomerGatewayRemoteVpnEndpoint{}

	_jsii_.Create(
		"cdk-extensions.ec2.CustomerGatewayRemoteVpnEndpoint",
		[]interface{}{customerGateway},
		&j,
	)

	return &j
}

// Creates a new instance of the CustomerGatewayRemoteVpnEndpoint class.
func NewCustomerGatewayRemoteVpnEndpoint_Override(c CustomerGatewayRemoteVpnEndpoint, customerGateway ICustomerGateway) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.ec2.CustomerGatewayRemoteVpnEndpoint",
		[]interface{}{customerGateway},
		c,
	)
}

func (c *jsiiProxy_CustomerGatewayRemoteVpnEndpoint) Bind(_scope constructs.IConstruct) *RemoteVpnEndpointConfiguration {
	if err := c.validateBindParameters(_scope); err != nil {
		panic(err)
	}
	var returns *RemoteVpnEndpointConfiguration

	_jsii_.Invoke(
		c,
		"bind",
		[]interface{}{_scope},
		&returns,
	)

	return returns
}

