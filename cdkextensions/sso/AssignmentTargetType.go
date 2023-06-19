package sso

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"
)

// Provides a wrapper around the accepted values for the IAM Identity Center [Assignment.TargetType attribute](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sso-assignment.html#cfn-sso-assignment-targettype).
//
// Accepted values are provided as static properties that can be used when
// configuring an assignment.
type AssignmentTargetType interface {
	// The name describing the type of target.
	Name() *string
}

// The jsii proxy struct for AssignmentTargetType
type jsiiProxy_AssignmentTargetType struct {
	_ byte // padding
}

func (j *jsiiProxy_AssignmentTargetType) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}


// An escape hatch method that allows specifying a custom target type in the even more options are added and the provided static types are yet to catch up.
//
// It is recommended that the provided static types be used when possible
// instead of calling `of`.
//
// Returns: An {@link AssignmentTargetType } object representing the specified type.
func AssignmentTargetType_Of(name *string) AssignmentTargetType {
	_init_.Initialize()

	if err := validateAssignmentTargetType_OfParameters(name); err != nil {
		panic(err)
	}
	var returns AssignmentTargetType

	_jsii_.StaticInvoke(
		"cdk-extensions.sso.AssignmentTargetType",
		"of",
		[]interface{}{name},
		&returns,
	)

	return returns
}

func AssignmentTargetType_AWS_ACCOUNT() AssignmentTargetType {
	_init_.Initialize()
	var returns AssignmentTargetType
	_jsii_.StaticGet(
		"cdk-extensions.sso.AssignmentTargetType",
		"AWS_ACCOUNT",
		&returns,
	)
	return returns
}

