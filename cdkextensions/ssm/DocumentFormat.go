package ssm

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"
)

type DocumentFormat interface {
	Value() *string
}

// The jsii proxy struct for DocumentFormat
type jsiiProxy_DocumentFormat struct {
	_ byte // padding
}

func (j *jsiiProxy_DocumentFormat) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func DocumentFormat_Of(value *string) DocumentFormat {
	_init_.Initialize()

	if err := validateDocumentFormat_OfParameters(value); err != nil {
		panic(err)
	}
	var returns DocumentFormat

	_jsii_.StaticInvoke(
		"cdk-extensions.ssm.DocumentFormat",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func DocumentFormat_JSON() DocumentFormat {
	_init_.Initialize()
	var returns DocumentFormat
	_jsii_.StaticGet(
		"cdk-extensions.ssm.DocumentFormat",
		"JSON",
		&returns,
	)
	return returns
}

func DocumentFormat_TEXT() DocumentFormat {
	_init_.Initialize()
	var returns DocumentFormat
	_jsii_.StaticGet(
		"cdk-extensions.ssm.DocumentFormat",
		"TEXT",
		&returns,
	)
	return returns
}

func DocumentFormat_YAML() DocumentFormat {
	_init_.Initialize()
	var returns DocumentFormat
	_jsii_.StaticGet(
		"cdk-extensions.ssm.DocumentFormat",
		"YAML",
		&returns,
	)
	return returns
}

