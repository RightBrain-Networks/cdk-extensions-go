package athena

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"
)

type WorkGroupState interface {
	Name() *string
}

// The jsii proxy struct for WorkGroupState
type jsiiProxy_WorkGroupState struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkGroupState) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}


func WorkGroupState_Of(name *string) WorkGroupState {
	_init_.Initialize()

	if err := validateWorkGroupState_OfParameters(name); err != nil {
		panic(err)
	}
	var returns WorkGroupState

	_jsii_.StaticInvoke(
		"cdk-extensions.athena.WorkGroupState",
		"of",
		[]interface{}{name},
		&returns,
	)

	return returns
}

func WorkGroupState_DISABLED() WorkGroupState {
	_init_.Initialize()
	var returns WorkGroupState
	_jsii_.StaticGet(
		"cdk-extensions.athena.WorkGroupState",
		"DISABLED",
		&returns,
	)
	return returns
}

func WorkGroupState_ENABLED() WorkGroupState {
	_init_.Initialize()
	var returns WorkGroupState
	_jsii_.StaticGet(
		"cdk-extensions.athena.WorkGroupState",
		"ENABLED",
		&returns,
	)
	return returns
}

