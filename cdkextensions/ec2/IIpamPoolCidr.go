package ec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type IIpamPoolCidr interface {
	IpamPoolCidrId() *string
	IpamPoolCidrState() *string
}

// The jsii proxy for IIpamPoolCidr
type jsiiProxy_IIpamPoolCidr struct {
	_ byte // padding
}

func (j *jsiiProxy_IIpamPoolCidr) IpamPoolCidrId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipamPoolCidrId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IIpamPoolCidr) IpamPoolCidrState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipamPoolCidrState",
		&returns,
	)
	return returns
}

