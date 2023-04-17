package athena

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"
)

type AnalyticsEngine interface {
}

// The jsii proxy struct for AnalyticsEngine
type jsiiProxy_AnalyticsEngine struct {
	_ byte // padding
}

func NewAnalyticsEngine() AnalyticsEngine {
	_init_.Initialize()

	j := jsiiProxy_AnalyticsEngine{}

	_jsii_.Create(
		"cdk-extensions.athena.AnalyticsEngine",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewAnalyticsEngine_Override(a AnalyticsEngine) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.athena.AnalyticsEngine",
		nil, // no parameters
		a,
	)
}

func AnalyticsEngine_ApacheSpark(options ApacheSparkEngineOptions) IAnalyticsEngine {
	_init_.Initialize()

	if err := validateAnalyticsEngine_ApacheSparkParameters(options); err != nil {
		panic(err)
	}
	var returns IAnalyticsEngine

	_jsii_.StaticInvoke(
		"cdk-extensions.athena.AnalyticsEngine",
		"apacheSpark",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func AnalyticsEngine_AthenaSql(options *AthenaSqlEngineOptions) IAnalyticsEngine {
	_init_.Initialize()

	if err := validateAnalyticsEngine_AthenaSqlParameters(options); err != nil {
		panic(err)
	}
	var returns IAnalyticsEngine

	_jsii_.StaticInvoke(
		"cdk-extensions.athena.AnalyticsEngine",
		"athenaSql",
		[]interface{}{options},
		&returns,
	)

	return returns
}

