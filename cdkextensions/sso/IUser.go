package sso

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type IUser interface {
	// A GUID identifier for a user object in IAM Identity Center (For example, f81d4fae-7dec-11d0-a765-00a0c91e6bf6).
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

