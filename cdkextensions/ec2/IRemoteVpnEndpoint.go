package ec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// An object that can be used to retrieve the details for the remote end of a VPN connection.
type IRemoteVpnEndpoint interface {
	// Produces a configuration that can be used when configuring the remote end of a VPN connection.
	Bind(scope constructs.IConstruct) *RemoteVpnEndpointConfiguration
}

// The jsii proxy for IRemoteVpnEndpoint
type jsiiProxy_IRemoteVpnEndpoint struct {
	_ byte // padding
}

func (i *jsiiProxy_IRemoteVpnEndpoint) Bind(scope constructs.IConstruct) *RemoteVpnEndpointConfiguration {
	if err := i.validateBindParameters(scope); err != nil {
		panic(err)
	}
	var returns *RemoteVpnEndpointConfiguration

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

