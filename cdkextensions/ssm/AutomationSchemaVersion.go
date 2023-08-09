package ssm

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"
)

type AutomationSchemaVersion interface {
	Version() *string
}

// The jsii proxy struct for AutomationSchemaVersion
type jsiiProxy_AutomationSchemaVersion struct {
	_ byte // padding
}

func (j *jsiiProxy_AutomationSchemaVersion) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}


func AutomationSchemaVersion_Of(version *string) AutomationSchemaVersion {
	_init_.Initialize()

	if err := validateAutomationSchemaVersion_OfParameters(version); err != nil {
		panic(err)
	}
	var returns AutomationSchemaVersion

	_jsii_.StaticInvoke(
		"cdk-extensions.ssm.AutomationSchemaVersion",
		"of",
		[]interface{}{version},
		&returns,
	)

	return returns
}

func AutomationSchemaVersion_VER_0_3() AutomationSchemaVersion {
	_init_.Initialize()
	var returns AutomationSchemaVersion
	_jsii_.StaticGet(
		"cdk-extensions.ssm.AutomationSchemaVersion",
		"VER_0_3",
		&returns,
	)
	return returns
}

