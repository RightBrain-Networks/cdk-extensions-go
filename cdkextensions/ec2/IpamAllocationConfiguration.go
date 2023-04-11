package ec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"
)

type IpamAllocationConfiguration interface {
}

// The jsii proxy struct for IpamAllocationConfiguration
type jsiiProxy_IpamAllocationConfiguration struct {
	_ byte // padding
}

func NewIpamAllocationConfiguration() IpamAllocationConfiguration {
	_init_.Initialize()

	j := jsiiProxy_IpamAllocationConfiguration{}

	_jsii_.Create(
		"cdk-extensions.ec2.IpamAllocationConfiguration",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewIpamAllocationConfiguration_Override(i IpamAllocationConfiguration) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.ec2.IpamAllocationConfiguration",
		nil, // no parameters
		i,
	)
}

func IpamAllocationConfiguration_Auto() IIpamAllocationConfiguration {
	_init_.Initialize()

	var returns IIpamAllocationConfiguration

	_jsii_.StaticInvoke(
		"cdk-extensions.ec2.IpamAllocationConfiguration",
		"auto",
		nil, // no parameters
		&returns,
	)

	return returns
}

func IpamAllocationConfiguration_Cidr(cidr *string) IIpamAllocationConfiguration {
	_init_.Initialize()

	if err := validateIpamAllocationConfiguration_CidrParameters(cidr); err != nil {
		panic(err)
	}
	var returns IIpamAllocationConfiguration

	_jsii_.StaticInvoke(
		"cdk-extensions.ec2.IpamAllocationConfiguration",
		"cidr",
		[]interface{}{cidr},
		&returns,
	)

	return returns
}

func IpamAllocationConfiguration_Netmask(length *float64) IIpamAllocationConfiguration {
	_init_.Initialize()

	if err := validateIpamAllocationConfiguration_NetmaskParameters(length); err != nil {
		panic(err)
	}
	var returns IIpamAllocationConfiguration

	_jsii_.StaticInvoke(
		"cdk-extensions.ec2.IpamAllocationConfiguration",
		"netmask",
		[]interface{}{length},
		&returns,
	)

	return returns
}

