package ssm

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"
)

type DocumentContent interface {
}

// The jsii proxy struct for DocumentContent
type jsiiProxy_DocumentContent struct {
	_ byte // padding
}

func NewDocumentContent() DocumentContent {
	_init_.Initialize()

	j := jsiiProxy_DocumentContent{}

	_jsii_.Create(
		"cdk-extensions.ssm.DocumentContent",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewDocumentContent_Override(d DocumentContent) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.ssm.DocumentContent",
		nil, // no parameters
		d,
	)
}

func DocumentContent_FromObject(props *ObjectContentProps) IDocumentContent {
	_init_.Initialize()

	if err := validateDocumentContent_FromObjectParameters(props); err != nil {
		panic(err)
	}
	var returns IDocumentContent

	_jsii_.StaticInvoke(
		"cdk-extensions.ssm.DocumentContent",
		"fromObject",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func DocumentContent_FromString(props *StringContentProps) IDocumentContent {
	_init_.Initialize()

	if err := validateDocumentContent_FromStringParameters(props); err != nil {
		panic(err)
	}
	var returns IDocumentContent

	_jsii_.StaticInvoke(
		"cdk-extensions.ssm.DocumentContent",
		"fromString",
		[]interface{}{props},
		&returns,
	)

	return returns
}

