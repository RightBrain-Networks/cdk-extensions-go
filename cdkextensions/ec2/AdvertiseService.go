package ec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"
)

type AdvertiseService interface {
	Name() *string
}

// The jsii proxy struct for AdvertiseService
type jsiiProxy_AdvertiseService struct {
	_ byte // padding
}

func (j *jsiiProxy_AdvertiseService) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}


func AdvertiseService_Of(name *string) AdvertiseService {
	_init_.Initialize()

	if err := validateAdvertiseService_OfParameters(name); err != nil {
		panic(err)
	}
	var returns AdvertiseService

	_jsii_.StaticInvoke(
		"cdk-extensions.ec2.AdvertiseService",
		"of",
		[]interface{}{name},
		&returns,
	)

	return returns
}

func AdvertiseService_EC2() AdvertiseService {
	_init_.Initialize()
	var returns AdvertiseService
	_jsii_.StaticGet(
		"cdk-extensions.ec2.AdvertiseService",
		"EC2",
		&returns,
	)
	return returns
}

func AdvertiseService_NONE() AdvertiseService {
	_init_.Initialize()
	var returns AdvertiseService
	_jsii_.StaticGet(
		"cdk-extensions.ec2.AdvertiseService",
		"NONE",
		&returns,
	)
	return returns
}

