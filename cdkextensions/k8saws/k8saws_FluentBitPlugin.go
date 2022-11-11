package k8saws

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/constructs-go/constructs/v10"
)

// A Fluent Bit plugin that allows for configuration of options and can be used to configure logging from containers.
type FluentBitPlugin interface {
	IFluentBitPlugin
	// The name of the fluent bit plugin.
	Name() *string
	// The type of fluent bit plugin.
	PluginType() *string
	// Builds a configuration for this plugin and returns the details for consumtion by a resource that is configuring logging.
	//
	// Returns: A configuration for the plugin that con be used by the resource
	// configuring logging.
	Bind(scope constructs.IConstruct) *ResolvedFluentBitConfiguration
	// Returns: A fluent bit config file representation of the passed properties.
	RenderConfigFile(config *map[string]interface{}) *string
}

// The jsii proxy struct for FluentBitPlugin
type jsiiProxy_FluentBitPlugin struct {
	jsiiProxy_IFluentBitPlugin
}

func (j *jsiiProxy_FluentBitPlugin) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitPlugin) PluginType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pluginType",
		&returns,
	)
	return returns
}


// Creates a new instance of the FluentBitPlugin class.
func NewFluentBitPlugin_Override(f FluentBitPlugin, options *FluentBitPluginCommonOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.k8s_aws.FluentBitPlugin",
		[]interface{}{options},
		f,
	)
}

func (f *jsiiProxy_FluentBitPlugin) Bind(scope constructs.IConstruct) *ResolvedFluentBitConfiguration {
	if err := f.validateBindParameters(scope); err != nil {
		panic(err)
	}
	var returns *ResolvedFluentBitConfiguration

	_jsii_.Invoke(
		f,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FluentBitPlugin) RenderConfigFile(config *map[string]interface{}) *string {
	if err := f.validateRenderConfigFileParameters(config); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		f,
		"renderConfigFile",
		[]interface{}{config},
		&returns,
	)

	return returns
}

