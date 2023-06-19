package resourcegroups

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type IGroup interface {
	GroupArn() *string
	GroupName() *string
}

// The jsii proxy for IGroup
type jsiiProxy_IGroup struct {
	_ byte // padding
}

func (j *jsiiProxy_IGroup) GroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGroup) GroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupName",
		&returns,
	)
	return returns
}

