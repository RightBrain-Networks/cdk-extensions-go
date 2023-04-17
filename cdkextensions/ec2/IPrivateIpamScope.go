package ec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type IPrivateIpamScope interface {
	IIpamScope
	AddPool() IIpamPool
}

// The jsii proxy for IPrivateIpamScope
type jsiiProxy_IPrivateIpamScope struct {
	jsiiProxy_IIpamScope
}

func (i *jsiiProxy_IPrivateIpamScope) AddPool() IIpamPool {
	var returns IIpamPool

	_jsii_.Invoke(
		i,
		"addPool",
		nil, // no parameters
		&returns,
	)

	return returns
}

