package resourcegroups

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2"
)

type StackReference interface {
}

// The jsii proxy struct for StackReference
type jsiiProxy_StackReference struct {
	_ byte // padding
}

func NewStackReference() StackReference {
	_init_.Initialize()

	j := jsiiProxy_StackReference{}

	_jsii_.Create(
		"cdk-extensions.resourcegroups.StackReference",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewStackReference_Override(s StackReference) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.resourcegroups.StackReference",
		nil, // no parameters
		s,
	)
}

func StackReference_FromStack(stack awscdk.Stack) IStackReference {
	_init_.Initialize()

	if err := validateStackReference_FromStackParameters(stack); err != nil {
		panic(err)
	}
	var returns IStackReference

	_jsii_.StaticInvoke(
		"cdk-extensions.resourcegroups.StackReference",
		"fromStack",
		[]interface{}{stack},
		&returns,
	)

	return returns
}

func StackReference_FromStackId(stackId *string) IStackReference {
	_init_.Initialize()

	if err := validateStackReference_FromStackIdParameters(stackId); err != nil {
		panic(err)
	}
	var returns IStackReference

	_jsii_.StaticInvoke(
		"cdk-extensions.resourcegroups.StackReference",
		"fromStackId",
		[]interface{}{stackId},
		&returns,
	)

	return returns
}

