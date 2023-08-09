package ssm

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"
)

type DocumentUpdateMethod interface {
	Value() *string
}

// The jsii proxy struct for DocumentUpdateMethod
type jsiiProxy_DocumentUpdateMethod struct {
	_ byte // padding
}

func (j *jsiiProxy_DocumentUpdateMethod) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func DocumentUpdateMethod_Of(value *string) DocumentUpdateMethod {
	_init_.Initialize()

	if err := validateDocumentUpdateMethod_OfParameters(value); err != nil {
		panic(err)
	}
	var returns DocumentUpdateMethod

	_jsii_.StaticInvoke(
		"cdk-extensions.ssm.DocumentUpdateMethod",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func DocumentUpdateMethod_NEW_VERSION() DocumentUpdateMethod {
	_init_.Initialize()
	var returns DocumentUpdateMethod
	_jsii_.StaticGet(
		"cdk-extensions.ssm.DocumentUpdateMethod",
		"NEW_VERSION",
		&returns,
	)
	return returns
}

func DocumentUpdateMethod_REPLACE() DocumentUpdateMethod {
	_init_.Initialize()
	var returns DocumentUpdateMethod
	_jsii_.StaticGet(
		"cdk-extensions.ssm.DocumentUpdateMethod",
		"REPLACE",
		&returns,
	)
	return returns
}

