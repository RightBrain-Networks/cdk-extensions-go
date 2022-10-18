package sso

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type IGroup interface {
	GroupId() *string
}

// The jsii proxy for IGroup
type jsiiProxy_IGroup struct {
	_ byte // padding
}

func (j *jsiiProxy_IGroup) GroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupId",
		&returns,
	)
	return returns
}

