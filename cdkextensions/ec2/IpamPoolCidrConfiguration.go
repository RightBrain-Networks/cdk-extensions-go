package ec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"
)

type IpamPoolCidrConfiguration interface {
}

// The jsii proxy struct for IpamPoolCidrConfiguration
type jsiiProxy_IpamPoolCidrConfiguration struct {
	_ byte // padding
}

func NewIpamPoolCidrConfiguration() IpamPoolCidrConfiguration {
	_init_.Initialize()

	j := jsiiProxy_IpamPoolCidrConfiguration{}

	_jsii_.Create(
		"cdk-extensions.ec2.IpamPoolCidrConfiguration",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewIpamPoolCidrConfiguration_Override(i IpamPoolCidrConfiguration) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.ec2.IpamPoolCidrConfiguration",
		nil, // no parameters
		i,
	)
}

func IpamPoolCidrConfiguration_Cidr(cidr *string) IIpamPoolCidrConfiguration {
	_init_.Initialize()

	if err := validateIpamPoolCidrConfiguration_CidrParameters(cidr); err != nil {
		panic(err)
	}
	var returns IIpamPoolCidrConfiguration

	_jsii_.StaticInvoke(
		"cdk-extensions.ec2.IpamPoolCidrConfiguration",
		"cidr",
		[]interface{}{cidr},
		&returns,
	)

	return returns
}

func IpamPoolCidrConfiguration_Netmask(length *float64) IIpamPoolCidrConfiguration {
	_init_.Initialize()

	if err := validateIpamPoolCidrConfiguration_NetmaskParameters(length); err != nil {
		panic(err)
	}
	var returns IIpamPoolCidrConfiguration

	_jsii_.StaticInvoke(
		"cdk-extensions.ec2.IpamPoolCidrConfiguration",
		"netmask",
		[]interface{}{length},
		&returns,
	)

	return returns
}

