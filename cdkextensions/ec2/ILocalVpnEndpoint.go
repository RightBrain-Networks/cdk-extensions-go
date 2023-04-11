package ec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

type ILocalVpnEndpoint interface {
	// Produces a configuration that can be used when configuring the local end of a VPN connection.
	Bind(scope constructs.IConstruct) *LocalVpnEndpointConfiguration
}

// The jsii proxy for ILocalVpnEndpoint
type jsiiProxy_ILocalVpnEndpoint struct {
	_ byte // padding
}

func (i *jsiiProxy_ILocalVpnEndpoint) Bind(scope constructs.IConstruct) *LocalVpnEndpointConfiguration {
	if err := i.validateBindParameters(scope); err != nil {
		panic(err)
	}
	var returns *LocalVpnEndpointConfiguration

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

