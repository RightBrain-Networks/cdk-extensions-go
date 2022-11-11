package k8saws

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// Represents a Fluent Bit plugin that allows for configuration of options and can be used to configure logging from containers.
type IFluentBitPlugin interface {
	// Builds a configuration for this plugin and returns the details for consumtion by a resource that is configuring logging.
	Bind(scope constructs.IConstruct) *ResolvedFluentBitConfiguration
	// The name of the fluent bit plugin.
	Name() *string
	// The type of fluent bit plugin.
	PluginType() *string
}

// The jsii proxy for IFluentBitPlugin
type jsiiProxy_IFluentBitPlugin struct {
	_ byte // padding
}

func (i *jsiiProxy_IFluentBitPlugin) Bind(scope constructs.IConstruct) *ResolvedFluentBitConfiguration {
	if err := i.validateBindParameters(scope); err != nil {
		panic(err)
	}
	var returns *ResolvedFluentBitConfiguration

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IFluentBitPlugin) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFluentBitPlugin) PluginType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pluginType",
		&returns,
	)
	return returns
}

