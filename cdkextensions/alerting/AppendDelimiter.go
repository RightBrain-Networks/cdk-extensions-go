package alerting

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"
)

type AppendDelimiter interface {
	Delimiter() *string
}

// The jsii proxy struct for AppendDelimiter
type jsiiProxy_AppendDelimiter struct {
	_ byte // padding
}

func (j *jsiiProxy_AppendDelimiter) Delimiter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delimiter",
		&returns,
	)
	return returns
}


func AppendDelimiter_Of(delimiter *string) AppendDelimiter {
	_init_.Initialize()

	if err := validateAppendDelimiter_OfParameters(delimiter); err != nil {
		panic(err)
	}
	var returns AppendDelimiter

	_jsii_.StaticInvoke(
		"cdk-extensions.alerting.AppendDelimiter",
		"of",
		[]interface{}{delimiter},
		&returns,
	)

	return returns
}

func AppendDelimiter_NEWLINE() AppendDelimiter {
	_init_.Initialize()
	var returns AppendDelimiter
	_jsii_.StaticGet(
		"cdk-extensions.alerting.AppendDelimiter",
		"NEWLINE",
		&returns,
	)
	return returns
}

func AppendDelimiter_PARAGRAPH() AppendDelimiter {
	_init_.Initialize()
	var returns AppendDelimiter
	_jsii_.StaticGet(
		"cdk-extensions.alerting.AppendDelimiter",
		"PARAGRAPH",
		&returns,
	)
	return returns
}

