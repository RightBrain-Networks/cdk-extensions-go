package config

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"
)

type RemediationTargetType interface {
	Value() *string
}

// The jsii proxy struct for RemediationTargetType
type jsiiProxy_RemediationTargetType struct {
	_ byte // padding
}

func (j *jsiiProxy_RemediationTargetType) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func RemediationTargetType_Of(value *string) RemediationTargetType {
	_init_.Initialize()

	if err := validateRemediationTargetType_OfParameters(value); err != nil {
		panic(err)
	}
	var returns RemediationTargetType

	_jsii_.StaticInvoke(
		"cdk-extensions.config.RemediationTargetType",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func RemediationTargetType_SSM_DOCUMENT() RemediationTargetType {
	_init_.Initialize()
	var returns RemediationTargetType
	_jsii_.StaticGet(
		"cdk-extensions.config.RemediationTargetType",
		"SSM_DOCUMENT",
		&returns,
	)
	return returns
}

