package sso

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type IUser interface {
	UserId() *string
}

// The jsii proxy for IUser
type jsiiProxy_IUser struct {
	_ byte // padding
}

func (j *jsiiProxy_IUser) UserId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userId",
		&returns,
	)
	return returns
}

