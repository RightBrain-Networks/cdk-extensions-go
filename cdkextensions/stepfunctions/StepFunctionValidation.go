package stepfunctions

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"
)

type StepFunctionValidation interface {
}

// The jsii proxy struct for StepFunctionValidation
type jsiiProxy_StepFunctionValidation struct {
	_ byte // padding
}

func NewStepFunctionValidation() StepFunctionValidation {
	_init_.Initialize()

	j := jsiiProxy_StepFunctionValidation{}

	_jsii_.Create(
		"cdk-extensions.stepfunctions.StepFunctionValidation",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewStepFunctionValidation_Override(s StepFunctionValidation) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.stepfunctions.StepFunctionValidation",
		nil, // no parameters
		s,
	)
}

func StepFunctionValidation_IsIntrinsic(value *string) *bool {
	_init_.Initialize()

	if err := validateStepFunctionValidation_IsIntrinsicParameters(value); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-extensions.stepfunctions.StepFunctionValidation",
		"isIntrinsic",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func StepFunctionValidation_IsJsonPath(value *string) *bool {
	_init_.Initialize()

	if err := validateStepFunctionValidation_IsJsonPathParameters(value); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-extensions.stepfunctions.StepFunctionValidation",
		"isJsonPath",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func StepFunctionValidation_IsStatesExpression(value *string) *bool {
	_init_.Initialize()

	if err := validateStepFunctionValidation_IsStatesExpressionParameters(value); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-extensions.stepfunctions.StepFunctionValidation",
		"isStatesExpression",
		[]interface{}{value},
		&returns,
	)

	return returns
}

