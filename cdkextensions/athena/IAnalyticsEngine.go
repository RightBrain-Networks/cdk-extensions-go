package athena

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

type IAnalyticsEngine interface {
	Bind(scope constructs.IConstruct, options *AnalyticsEngineBindProps) *AnalyticsEngineConfiguration
}

// The jsii proxy for IAnalyticsEngine
type jsiiProxy_IAnalyticsEngine struct {
	_ byte // padding
}

func (i *jsiiProxy_IAnalyticsEngine) Bind(scope constructs.IConstruct, options *AnalyticsEngineBindProps) *AnalyticsEngineConfiguration {
	if err := i.validateBindParameters(scope, options); err != nil {
		panic(err)
	}
	var returns *AnalyticsEngineConfiguration

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{scope, options},
		&returns,
	)

	return returns
}

