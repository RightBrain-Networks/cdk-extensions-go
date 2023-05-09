package ec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"
)

type Ipv6CidrAssignment interface {
}

// The jsii proxy struct for Ipv6CidrAssignment
type jsiiProxy_Ipv6CidrAssignment struct {
	_ byte // padding
}

func NewIpv6CidrAssignment() Ipv6CidrAssignment {
	_init_.Initialize()

	j := jsiiProxy_Ipv6CidrAssignment{}

	_jsii_.Create(
		"cdk-extensions.ec2.Ipv6CidrAssignment",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewIpv6CidrAssignment_Override(i Ipv6CidrAssignment) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.ec2.Ipv6CidrAssignment",
		nil, // no parameters
		i,
	)
}

func Ipv6CidrAssignment_Custom(options *Ipv4CidrAssignmentCustomOptions) IIpv6CidrAssignment {
	_init_.Initialize()

	if err := validateIpv6CidrAssignment_CustomParameters(options); err != nil {
		panic(err)
	}
	var returns IIpv6CidrAssignment

	_jsii_.StaticInvoke(
		"cdk-extensions.ec2.Ipv6CidrAssignment",
		"custom",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func Ipv6CidrAssignment_IpamPool(options *Ipv6CidrAssignmentIpamPoolOptions) IIpv6CidrAssignment {
	_init_.Initialize()

	if err := validateIpv6CidrAssignment_IpamPoolParameters(options); err != nil {
		panic(err)
	}
	var returns IIpv6CidrAssignment

	_jsii_.StaticInvoke(
		"cdk-extensions.ec2.Ipv6CidrAssignment",
		"ipamPool",
		[]interface{}{options},
		&returns,
	)

	return returns
}

