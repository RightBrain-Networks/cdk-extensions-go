package networkmanager

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type IGlobalNetwork interface {
	GlobalNetworkArn() *string
	GlobalNetworkId() *string
}

// The jsii proxy for IGlobalNetwork
type jsiiProxy_IGlobalNetwork struct {
	_ byte // padding
}

func (j *jsiiProxy_IGlobalNetwork) GlobalNetworkArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"globalNetworkArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGlobalNetwork) GlobalNetworkId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"globalNetworkId",
		&returns,
	)
	return returns
}

