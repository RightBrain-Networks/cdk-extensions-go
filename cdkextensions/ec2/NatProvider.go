package ec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

type NatProvider interface {
}

// The jsii proxy struct for NatProvider
type jsiiProxy_NatProvider struct {
	_ byte // padding
}

func NewNatProvider() NatProvider {
	_init_.Initialize()

	j := jsiiProxy_NatProvider{}

	_jsii_.Create(
		"cdk-extensions.ec2.NatProvider",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewNatProvider_Override(n NatProvider) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.ec2.NatProvider",
		nil, // no parameters
		n,
	)
}

func NatProvider_Gateway(props *awsec2.NatGatewayProps) awsec2.NatProvider {
	_init_.Initialize()

	if err := validateNatProvider_GatewayParameters(props); err != nil {
		panic(err)
	}
	var returns awsec2.NatProvider

	_jsii_.StaticInvoke(
		"cdk-extensions.ec2.NatProvider",
		"gateway",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func NatProvider_Instance(props *awsec2.NatInstanceProps) awsec2.NatProvider {
	_init_.Initialize()

	if err := validateNatProvider_InstanceParameters(props); err != nil {
		panic(err)
	}
	var returns awsec2.NatProvider

	_jsii_.StaticInvoke(
		"cdk-extensions.ec2.NatProvider",
		"instance",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func NatProvider_TransitGateway(props *TransitGatewayNatProviderOptions) awsec2.NatProvider {
	_init_.Initialize()

	if err := validateNatProvider_TransitGatewayParameters(props); err != nil {
		panic(err)
	}
	var returns awsec2.NatProvider

	_jsii_.StaticInvoke(
		"cdk-extensions.ec2.NatProvider",
		"transitGateway",
		[]interface{}{props},
		&returns,
	)

	return returns
}

