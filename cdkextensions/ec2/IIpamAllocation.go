package ec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type IIpamAllocation interface {
	IpamAllocationCidr() *string
	IpamAllocationId() *string
}

// The jsii proxy for IIpamAllocation
type jsiiProxy_IIpamAllocation struct {
	_ byte // padding
}

func (j *jsiiProxy_IIpamAllocation) IpamAllocationCidr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipamAllocationCidr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IIpamAllocation) IpamAllocationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipamAllocationId",
		&returns,
	)
	return returns
}

