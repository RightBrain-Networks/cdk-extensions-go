package sso

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Represents an entity that can be granted permissions via IAM Identity Center.
type IIdentityCenterPrincipal interface {
	// The unique ID that identifies the entity withing IAM Identity Center.
	PrincipalId() *string
	// The type of entity being represented.
	PrincipalType() IdentityCenterPrincipalType
}

// The jsii proxy for IIdentityCenterPrincipal
type jsiiProxy_IIdentityCenterPrincipal struct {
	_ byte // padding
}

func (j *jsiiProxy_IIdentityCenterPrincipal) PrincipalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"principalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IIdentityCenterPrincipal) PrincipalType() IdentityCenterPrincipalType {
	var returns IdentityCenterPrincipalType
	_jsii_.Get(
		j,
		"principalType",
		&returns,
	)
	return returns
}

