package ec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"
)

type PublicIpSource interface {
	Name() *string
}

// The jsii proxy struct for PublicIpSource
type jsiiProxy_PublicIpSource struct {
	_ byte // padding
}

func (j *jsiiProxy_PublicIpSource) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}


func PublicIpSource_Of(name *string) PublicIpSource {
	_init_.Initialize()

	if err := validatePublicIpSource_OfParameters(name); err != nil {
		panic(err)
	}
	var returns PublicIpSource

	_jsii_.StaticInvoke(
		"cdk-extensions.ec2.PublicIpSource",
		"of",
		[]interface{}{name},
		&returns,
	)

	return returns
}

func PublicIpSource_AMAZON() PublicIpSource {
	_init_.Initialize()
	var returns PublicIpSource
	_jsii_.StaticGet(
		"cdk-extensions.ec2.PublicIpSource",
		"AMAZON",
		&returns,
	)
	return returns
}

func PublicIpSource_BYOIP() PublicIpSource {
	_init_.Initialize()
	var returns PublicIpSource
	_jsii_.StaticGet(
		"cdk-extensions.ec2.PublicIpSource",
		"BYOIP",
		&returns,
	)
	return returns
}

