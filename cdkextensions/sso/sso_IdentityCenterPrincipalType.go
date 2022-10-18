package sso

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"
)

// Provides a wrapper around the accepted values for the IAM Identity Center [Assignment.PrincipalType attribute](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sso-assignment.html#cfn-sso-assignment-principaltype).
//
// Accepted values are provided as static properties that can be used when
// configuring an assignment.
type IdentityCenterPrincipalType interface {
	// The name for a type of IAM Identity Center Principal.
	Name() *string
}

// The jsii proxy struct for IdentityCenterPrincipalType
type jsiiProxy_IdentityCenterPrincipalType struct {
	_ byte // padding
}

func (j *jsiiProxy_IdentityCenterPrincipalType) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}


// An escape hatch method that allows specifying a custom principal types in the even more options are added and the provided static types are yet to catch up.
//
// It is recommended that the provided static types be used when possible
// instead of calling `of`.
func IdentityCenterPrincipalType_Of(name *string) IdentityCenterPrincipalType {
	_init_.Initialize()

	if err := validateIdentityCenterPrincipalType_OfParameters(name); err != nil {
		panic(err)
	}
	var returns IdentityCenterPrincipalType

	_jsii_.StaticInvoke(
		"cdk-extensions.sso.IdentityCenterPrincipalType",
		"of",
		[]interface{}{name},
		&returns,
	)

	return returns
}

func IdentityCenterPrincipalType_GROUP() IdentityCenterPrincipalType {
	_init_.Initialize()
	var returns IdentityCenterPrincipalType
	_jsii_.StaticGet(
		"cdk-extensions.sso.IdentityCenterPrincipalType",
		"GROUP",
		&returns,
	)
	return returns
}

func IdentityCenterPrincipalType_USER() IdentityCenterPrincipalType {
	_init_.Initialize()
	var returns IdentityCenterPrincipalType
	_jsii_.StaticGet(
		"cdk-extensions.sso.IdentityCenterPrincipalType",
		"USER",
		&returns,
	)
	return returns
}

