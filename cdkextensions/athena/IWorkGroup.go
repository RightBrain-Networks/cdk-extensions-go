package athena

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type IWorkGroup interface {
	WorkGroupArn() *string
	WorkGroupCreationTime() *string
	WorkGroupEffectiveEngineVersion() *string
	WorkGroupName() *string
}

// The jsii proxy for IWorkGroup
type jsiiProxy_IWorkGroup struct {
	_ byte // padding
}

func (j *jsiiProxy_IWorkGroup) WorkGroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workGroupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IWorkGroup) WorkGroupCreationTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workGroupCreationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IWorkGroup) WorkGroupEffectiveEngineVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workGroupEffectiveEngineVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IWorkGroup) WorkGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workGroupName",
		&returns,
	)
	return returns
}

