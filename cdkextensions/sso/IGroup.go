package sso

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type IGroup interface {
	// A GUID identifier for a group object in IAM Identity Center are GUIDs (For example, f81d4fae-7dec-11d0-a765-00a0c91e6bf6).
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

