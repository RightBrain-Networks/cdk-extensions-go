package ssm

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"
)

type DocumentType interface {
	Name() *string
}

// The jsii proxy struct for DocumentType
type jsiiProxy_DocumentType struct {
	_ byte // padding
}

func (j *jsiiProxy_DocumentType) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}


func DocumentType_Of(value *string) DocumentType {
	_init_.Initialize()

	if err := validateDocumentType_OfParameters(value); err != nil {
		panic(err)
	}
	var returns DocumentType

	_jsii_.StaticInvoke(
		"cdk-extensions.ssm.DocumentType",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func DocumentType_APPLICATION_CONFIGURATION() DocumentType {
	_init_.Initialize()
	var returns DocumentType
	_jsii_.StaticGet(
		"cdk-extensions.ssm.DocumentType",
		"APPLICATION_CONFIGURATION",
		&returns,
	)
	return returns
}

func DocumentType_APPLICATION_CONFIGURATION_SCHEMA() DocumentType {
	_init_.Initialize()
	var returns DocumentType
	_jsii_.StaticGet(
		"cdk-extensions.ssm.DocumentType",
		"APPLICATION_CONFIGURATION_SCHEMA",
		&returns,
	)
	return returns
}

func DocumentType_AUTOMATION() DocumentType {
	_init_.Initialize()
	var returns DocumentType
	_jsii_.StaticGet(
		"cdk-extensions.ssm.DocumentType",
		"AUTOMATION",
		&returns,
	)
	return returns
}

func DocumentType_AUTOMATION_CHANGE_TEMPLATE() DocumentType {
	_init_.Initialize()
	var returns DocumentType
	_jsii_.StaticGet(
		"cdk-extensions.ssm.DocumentType",
		"AUTOMATION_CHANGE_TEMPLATE",
		&returns,
	)
	return returns
}

func DocumentType_COMMAND() DocumentType {
	_init_.Initialize()
	var returns DocumentType
	_jsii_.StaticGet(
		"cdk-extensions.ssm.DocumentType",
		"COMMAND",
		&returns,
	)
	return returns
}

func DocumentType_DEPLOYMENT_STRATEGY() DocumentType {
	_init_.Initialize()
	var returns DocumentType
	_jsii_.StaticGet(
		"cdk-extensions.ssm.DocumentType",
		"DEPLOYMENT_STRATEGY",
		&returns,
	)
	return returns
}

func DocumentType_PACKAGE() DocumentType {
	_init_.Initialize()
	var returns DocumentType
	_jsii_.StaticGet(
		"cdk-extensions.ssm.DocumentType",
		"PACKAGE",
		&returns,
	)
	return returns
}

func DocumentType_POLICY() DocumentType {
	_init_.Initialize()
	var returns DocumentType
	_jsii_.StaticGet(
		"cdk-extensions.ssm.DocumentType",
		"POLICY",
		&returns,
	)
	return returns
}

func DocumentType_SESSION() DocumentType {
	_init_.Initialize()
	var returns DocumentType
	_jsii_.StaticGet(
		"cdk-extensions.ssm.DocumentType",
		"SESSION",
		&returns,
	)
	return returns
}

