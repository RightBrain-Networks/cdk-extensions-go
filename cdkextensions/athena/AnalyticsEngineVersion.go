package athena

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"
)

type AnalyticsEngineVersion interface {
	Name() *string
}

// The jsii proxy struct for AnalyticsEngineVersion
type jsiiProxy_AnalyticsEngineVersion struct {
	_ byte // padding
}

func (j *jsiiProxy_AnalyticsEngineVersion) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}


func NewAnalyticsEngineVersion_Override(a AnalyticsEngineVersion) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.athena.AnalyticsEngineVersion",
		nil, // no parameters
		a,
	)
}

