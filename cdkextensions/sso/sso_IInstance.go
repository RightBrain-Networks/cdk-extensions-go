package sso

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type IInstance interface {
	InstanceArn() *string
	InstanceId() *string
}

// The jsii proxy for IInstance
type jsiiProxy_IInstance struct {
	_ byte // padding
}

func (j *jsiiProxy_IInstance) InstanceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IInstance) InstanceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceId",
		&returns,
	)
	return returns
}

