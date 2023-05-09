package ec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"
)

type Ipv4CidrAssignment interface {
}

// The jsii proxy struct for Ipv4CidrAssignment
type jsiiProxy_Ipv4CidrAssignment struct {
	_ byte // padding
}

func NewIpv4CidrAssignment() Ipv4CidrAssignment {
	_init_.Initialize()

	j := jsiiProxy_Ipv4CidrAssignment{}

	_jsii_.Create(
		"cdk-extensions.ec2.Ipv4CidrAssignment",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewIpv4CidrAssignment_Override(i Ipv4CidrAssignment) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.ec2.Ipv4CidrAssignment",
		nil, // no parameters
		i,
	)
}

func Ipv4CidrAssignment_Custom(options *Ipv4CidrAssignmentCustomOptions) IIpv4CidrAssignment {
	_init_.Initialize()

	if err := validateIpv4CidrAssignment_CustomParameters(options); err != nil {
		panic(err)
	}
	var returns IIpv4CidrAssignment

	_jsii_.StaticInvoke(
		"cdk-extensions.ec2.Ipv4CidrAssignment",
		"custom",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func Ipv4CidrAssignment_IpamPool(options *Ipv4CidrAssignmentIpamPoolOptions) IIpv4CidrAssignment {
	_init_.Initialize()

	if err := validateIpv4CidrAssignment_IpamPoolParameters(options); err != nil {
		panic(err)
	}
	var returns IIpv4CidrAssignment

	_jsii_.StaticInvoke(
		"cdk-extensions.ec2.Ipv4CidrAssignment",
		"ipamPool",
		[]interface{}{options},
		&returns,
	)

	return returns
}

