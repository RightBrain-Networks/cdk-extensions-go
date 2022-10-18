package sso

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"
)

// Represents a resource that can have permissions granted for using IAM Identity Center such as an AWS account.
type AssignmentTarget interface {
	// The unique identifier for the resource for which permissions will be granted.
	TargetId() *string
	// The type of resource for which permissions will be granted.
	TargetType() AssignmentTargetType
}

// The jsii proxy struct for AssignmentTarget
type jsiiProxy_AssignmentTarget struct {
	_ byte // padding
}

func (j *jsiiProxy_AssignmentTarget) TargetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AssignmentTarget) TargetType() AssignmentTargetType {
	var returns AssignmentTargetType
	_jsii_.Get(
		j,
		"targetType",
		&returns,
	)
	return returns
}


// Creates an assignment target that represents an AWS account.
//
// Returns: An AssignmentTarget representing the AWS account.
func AssignmentTarget_AwsAccount(accountId *string) AssignmentTarget {
	_init_.Initialize()

	if err := validateAssignmentTarget_AwsAccountParameters(accountId); err != nil {
		panic(err)
	}
	var returns AssignmentTarget

	_jsii_.StaticInvoke(
		"cdk-extensions.sso.AssignmentTarget",
		"awsAccount",
		[]interface{}{accountId},
		&returns,
	)

	return returns
}

// An escape hatch method that allows specifying a custom target for an assignment in the event new target options are added and the provided methods for configuring targets are yet to catch up.
//
// It is recommended that the provided static methods be used whenever
// possible for configuring assignment targets instead of calling `of`.
func AssignmentTarget_Of(targetType AssignmentTargetType, targetId *string) AssignmentTarget {
	_init_.Initialize()

	if err := validateAssignmentTarget_OfParameters(targetType, targetId); err != nil {
		panic(err)
	}
	var returns AssignmentTarget

	_jsii_.StaticInvoke(
		"cdk-extensions.sso.AssignmentTarget",
		"of",
		[]interface{}{targetType, targetId},
		&returns,
	)

	return returns
}

