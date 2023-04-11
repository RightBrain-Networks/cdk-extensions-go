package ec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"
)

type IpamConsumer interface {
	Name() *string
}

// The jsii proxy struct for IpamConsumer
type jsiiProxy_IpamConsumer struct {
	_ byte // padding
}

func (j *jsiiProxy_IpamConsumer) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}


func IpamConsumer_Of(name *string) IpamConsumer {
	_init_.Initialize()

	if err := validateIpamConsumer_OfParameters(name); err != nil {
		panic(err)
	}
	var returns IpamConsumer

	_jsii_.StaticInvoke(
		"cdk-extensions.ec2.IpamConsumer",
		"of",
		[]interface{}{name},
		&returns,
	)

	return returns
}

func IpamConsumer_EC2() IpamConsumer {
	_init_.Initialize()
	var returns IpamConsumer
	_jsii_.StaticGet(
		"cdk-extensions.ec2.IpamConsumer",
		"EC2",
		&returns,
	)
	return returns
}

