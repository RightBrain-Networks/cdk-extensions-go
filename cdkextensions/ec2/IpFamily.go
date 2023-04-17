package ec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"
)

type IpFamily interface {
	Name() *string
}

// The jsii proxy struct for IpFamily
type jsiiProxy_IpFamily struct {
	_ byte // padding
}

func (j *jsiiProxy_IpFamily) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}


func IpFamily_Of(name *string) IpFamily {
	_init_.Initialize()

	if err := validateIpFamily_OfParameters(name); err != nil {
		panic(err)
	}
	var returns IpFamily

	_jsii_.StaticInvoke(
		"cdk-extensions.ec2.IpFamily",
		"of",
		[]interface{}{name},
		&returns,
	)

	return returns
}

func IpFamily_IPV4() IpFamily {
	_init_.Initialize()
	var returns IpFamily
	_jsii_.StaticGet(
		"cdk-extensions.ec2.IpFamily",
		"IPV4",
		&returns,
	)
	return returns
}

func IpFamily_IPV6() IpFamily {
	_init_.Initialize()
	var returns IpFamily
	_jsii_.StaticGet(
		"cdk-extensions.ec2.IpFamily",
		"IPV6",
		&returns,
	)
	return returns
}

