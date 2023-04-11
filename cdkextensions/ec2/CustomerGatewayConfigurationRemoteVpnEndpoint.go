package ec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/constructs-go/constructs/v10"
)

// Specifies a remote VPN endpoint device by directly specifyingits details.
type CustomerGatewayConfigurationRemoteVpnEndpoint interface {
	IRemoteVpnEndpoint
	// The details of the device on the remote end of the VPN connection.
	Configuration() *CustomerGatewayProps
	// The customer gateway that was created to represent the device on the remote end of the VPN connection.
	CustomerGateway() CustomerGateway
	// Produces a configuration that can be used when configuring the remote end of a VPN connection.
	Bind(scope constructs.IConstruct) *RemoteVpnEndpointConfiguration
}

// The jsii proxy struct for CustomerGatewayConfigurationRemoteVpnEndpoint
type jsiiProxy_CustomerGatewayConfigurationRemoteVpnEndpoint struct {
	jsiiProxy_IRemoteVpnEndpoint
}

func (j *jsiiProxy_CustomerGatewayConfigurationRemoteVpnEndpoint) Configuration() *CustomerGatewayProps {
	var returns *CustomerGatewayProps
	_jsii_.Get(
		j,
		"configuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerGatewayConfigurationRemoteVpnEndpoint) CustomerGateway() CustomerGateway {
	var returns CustomerGateway
	_jsii_.Get(
		j,
		"customerGateway",
		&returns,
	)
	return returns
}


// Creates a new instance of the CustomerGatewayConfigurationRemoteVpnEndpoint class.
func NewCustomerGatewayConfigurationRemoteVpnEndpoint(configuration *CustomerGatewayProps) CustomerGatewayConfigurationRemoteVpnEndpoint {
	_init_.Initialize()

	if err := validateNewCustomerGatewayConfigurationRemoteVpnEndpointParameters(configuration); err != nil {
		panic(err)
	}
	j := jsiiProxy_CustomerGatewayConfigurationRemoteVpnEndpoint{}

	_jsii_.Create(
		"cdk-extensions.ec2.CustomerGatewayConfigurationRemoteVpnEndpoint",
		[]interface{}{configuration},
		&j,
	)

	return &j
}

// Creates a new instance of the CustomerGatewayConfigurationRemoteVpnEndpoint class.
func NewCustomerGatewayConfigurationRemoteVpnEndpoint_Override(c CustomerGatewayConfigurationRemoteVpnEndpoint, configuration *CustomerGatewayProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.ec2.CustomerGatewayConfigurationRemoteVpnEndpoint",
		[]interface{}{configuration},
		c,
	)
}

func (c *jsiiProxy_CustomerGatewayConfigurationRemoteVpnEndpoint) Bind(scope constructs.IConstruct) *RemoteVpnEndpointConfiguration {
	if err := c.validateBindParameters(scope); err != nil {
		panic(err)
	}
	var returns *RemoteVpnEndpointConfiguration

	_jsii_.Invoke(
		c,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

