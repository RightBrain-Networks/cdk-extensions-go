package sso

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type IInstance interface {
	// The ARN of the IAM Identity Center instance under which the operation will be executed.
	//
	// For more information about ARNs, see Amazon Resource Names (ARNs) and AWS Service Namespaces in the AWS General Reference.
	// See: [AWS::SSO::Assignment](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sso-assignment.html#cfn-sso-assignment-instancearn)
	//
	InstanceArn() *string
	// The ID of the IAM Identity Center instance under which the operation will be executed.
	InstanceId() *string
}

// The jsii proxy for IInstance
type jsiiProxy_IInstance struct {
	_ byte // padding
}

func (j *jsiiProxy_IInstance) InstanceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IInstance) InstanceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceId",
		&returns,
	)
	return returns
}

