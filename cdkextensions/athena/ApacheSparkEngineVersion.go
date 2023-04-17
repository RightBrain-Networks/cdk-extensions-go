package athena

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"
)

type ApacheSparkEngineVersion interface {
	AnalyticsEngineVersion
	Name() *string
}

// The jsii proxy struct for ApacheSparkEngineVersion
type jsiiProxy_ApacheSparkEngineVersion struct {
	jsiiProxy_AnalyticsEngineVersion
}

func (j *jsiiProxy_ApacheSparkEngineVersion) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}


func ApacheSparkEngineVersion_Of(name *string) ApacheSparkEngineVersion {
	_init_.Initialize()

	if err := validateApacheSparkEngineVersion_OfParameters(name); err != nil {
		panic(err)
	}
	var returns ApacheSparkEngineVersion

	_jsii_.StaticInvoke(
		"cdk-extensions.athena.ApacheSparkEngineVersion",
		"of",
		[]interface{}{name},
		&returns,
	)

	return returns
}

func ApacheSparkEngineVersion_AUTO() ApacheSparkEngineVersion {
	_init_.Initialize()
	var returns ApacheSparkEngineVersion
	_jsii_.StaticGet(
		"cdk-extensions.athena.ApacheSparkEngineVersion",
		"AUTO",
		&returns,
	)
	return returns
}

func ApacheSparkEngineVersion_V3() ApacheSparkEngineVersion {
	_init_.Initialize()
	var returns ApacheSparkEngineVersion
	_jsii_.StaticGet(
		"cdk-extensions.athena.ApacheSparkEngineVersion",
		"V3",
		&returns,
	)
	return returns
}

