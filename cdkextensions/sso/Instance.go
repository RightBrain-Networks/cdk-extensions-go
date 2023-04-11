package sso

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/constructs-go/constructs/v10"
)

type Instance interface {
}

// The jsii proxy struct for Instance
type jsiiProxy_Instance struct {
	_ byte // padding
}

func NewInstance() Instance {
	_init_.Initialize()

	j := jsiiProxy_Instance{}

	_jsii_.Create(
		"cdk-extensions.sso.Instance",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewInstance_Override(i Instance) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.sso.Instance",
		nil, // no parameters
		i,
	)
}

func Instance_FromArn(scope constructs.IConstruct, id *string, arn *string) IInstance {
	_init_.Initialize()

	if err := validateInstance_FromArnParameters(scope, id, arn); err != nil {
		panic(err)
	}
	var returns IInstance

	_jsii_.StaticInvoke(
		"cdk-extensions.sso.Instance",
		"fromArn",
		[]interface{}{scope, id, arn},
		&returns,
	)

	return returns
}

func Instance_FromInstanceId(scope constructs.IConstruct, id *string, instanceId *string) IInstance {
	_init_.Initialize()

	if err := validateInstance_FromInstanceIdParameters(scope, id, instanceId); err != nil {
		panic(err)
	}
	var returns IInstance

	_jsii_.StaticInvoke(
		"cdk-extensions.sso.Instance",
		"fromInstanceId",
		[]interface{}{scope, id, instanceId},
		&returns,
	)

	return returns
}

