package sso

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Represents an IAM Identity Center permission set resource.
type IPermissionSet interface {
	PermissionSetArn() *string
}

// The jsii proxy for IPermissionSet
type jsiiProxy_IPermissionSet struct {
	_ byte // padding
}

func (j *jsiiProxy_IPermissionSet) PermissionSetArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"permissionSetArn",
		&returns,
	)
	return returns
}

