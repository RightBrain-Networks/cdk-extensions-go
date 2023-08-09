package securityhub

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"
)

type ControlFindingGenerator interface {
	Value() *string
}

// The jsii proxy struct for ControlFindingGenerator
type jsiiProxy_ControlFindingGenerator struct {
	_ byte // padding
}

func (j *jsiiProxy_ControlFindingGenerator) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func ControlFindingGenerator_Of(value *string) ControlFindingGenerator {
	_init_.Initialize()

	if err := validateControlFindingGenerator_OfParameters(value); err != nil {
		panic(err)
	}
	var returns ControlFindingGenerator

	_jsii_.StaticInvoke(
		"cdk-extensions.securityhub.ControlFindingGenerator",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func ControlFindingGenerator_SECURITY_CONTROL() ControlFindingGenerator {
	_init_.Initialize()
	var returns ControlFindingGenerator
	_jsii_.StaticGet(
		"cdk-extensions.securityhub.ControlFindingGenerator",
		"SECURITY_CONTROL",
		&returns,
	)
	return returns
}

func ControlFindingGenerator_STANDARD_CONTROL() ControlFindingGenerator {
	_init_.Initialize()
	var returns ControlFindingGenerator
	_jsii_.StaticGet(
		"cdk-extensions.securityhub.ControlFindingGenerator",
		"STANDARD_CONTROL",
		&returns,
	)
	return returns
}

