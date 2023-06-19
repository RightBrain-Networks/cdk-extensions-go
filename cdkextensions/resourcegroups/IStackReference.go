package resourcegroups

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
)

type IStackReference interface {
	StackConstruct() awscdk.Stack
	StackId() *string
}

// The jsii proxy for IStackReference
type jsiiProxy_IStackReference struct {
	_ byte // padding
}

func (j *jsiiProxy_IStackReference) StackConstruct() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stackConstruct",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStackReference) StackId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stackId",
		&returns,
	)
	return returns
}

