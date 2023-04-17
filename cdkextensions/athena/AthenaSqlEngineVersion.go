package athena

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"
)

type AthenaSqlEngineVersion interface {
	AnalyticsEngineVersion
	Name() *string
}

// The jsii proxy struct for AthenaSqlEngineVersion
type jsiiProxy_AthenaSqlEngineVersion struct {
	jsiiProxy_AnalyticsEngineVersion
}

func (j *jsiiProxy_AthenaSqlEngineVersion) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}


func AthenaSqlEngineVersion_Of(name *string) AthenaSqlEngineVersion {
	_init_.Initialize()

	if err := validateAthenaSqlEngineVersion_OfParameters(name); err != nil {
		panic(err)
	}
	var returns AthenaSqlEngineVersion

	_jsii_.StaticInvoke(
		"cdk-extensions.athena.AthenaSqlEngineVersion",
		"of",
		[]interface{}{name},
		&returns,
	)

	return returns
}

func AthenaSqlEngineVersion_AUTO() AthenaSqlEngineVersion {
	_init_.Initialize()
	var returns AthenaSqlEngineVersion
	_jsii_.StaticGet(
		"cdk-extensions.athena.AthenaSqlEngineVersion",
		"AUTO",
		&returns,
	)
	return returns
}

func AthenaSqlEngineVersion_V2() AthenaSqlEngineVersion {
	_init_.Initialize()
	var returns AthenaSqlEngineVersion
	_jsii_.StaticGet(
		"cdk-extensions.athena.AthenaSqlEngineVersion",
		"V2",
		&returns,
	)
	return returns
}

func AthenaSqlEngineVersion_V3() AthenaSqlEngineVersion {
	_init_.Initialize()
	var returns AthenaSqlEngineVersion
	_jsii_.StaticGet(
		"cdk-extensions.athena.AthenaSqlEngineVersion",
		"V3",
		&returns,
	)
	return returns
}

