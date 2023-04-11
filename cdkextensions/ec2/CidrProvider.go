package ec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"
)

type CidrProvider interface {
}

// The jsii proxy struct for CidrProvider
type jsiiProxy_CidrProvider struct {
	_ byte // padding
}

func NewCidrProvider() CidrProvider {
	_init_.Initialize()

	j := jsiiProxy_CidrProvider{}

	_jsii_.Create(
		"cdk-extensions.ec2.CidrProvider",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewCidrProvider_Override(c CidrProvider) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.ec2.CidrProvider",
		nil, // no parameters
		c,
	)
}

func CidrProvider_Cidr(cidr *string) ICidrProvider {
	_init_.Initialize()

	if err := validateCidrProvider_CidrParameters(cidr); err != nil {
		panic(err)
	}
	var returns ICidrProvider

	_jsii_.StaticInvoke(
		"cdk-extensions.ec2.CidrProvider",
		"cidr",
		[]interface{}{cidr},
		&returns,
	)

	return returns
}

func CidrProvider_IpamPool(pool IIpamPool, netmask *float64) ICidrProvider {
	_init_.Initialize()

	if err := validateCidrProvider_IpamPoolParameters(pool, netmask); err != nil {
		panic(err)
	}
	var returns ICidrProvider

	_jsii_.StaticInvoke(
		"cdk-extensions.ec2.CidrProvider",
		"ipamPool",
		[]interface{}{pool, netmask},
		&returns,
	)

	return returns
}

