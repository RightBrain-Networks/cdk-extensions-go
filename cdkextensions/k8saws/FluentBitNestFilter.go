package k8saws

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/constructs-go/constructs/v10"
)

// A Fluent Bit filter that allows operating on or with nested data.
type FluentBitNestFilter interface {
	FluentBitFilterPluginBase
	// Prefix affected keys with this string.
	AddPrefix() *string
	// The pattern to match for records that this output should apply to.
	Match() FluentBitMatch
	// The name of the fluent bit plugin.
	Name() *string
	// Operation specific details for the plugin.
	Operation() NestFilterOperation
	// The type of fluent bit plugin.
	PluginType() *string
	// Remove prefix from affected keys if it matches this string.
	RemovePrefix() *string
	// Builds a configuration for this plugin and returns the details for consumtion by a resource that is configuring logging.
	//
	// Returns: A configuration for the plugin that con be used by the resource
	// configuring logging.
	Bind(_scope constructs.IConstruct) *ResolvedFluentBitConfiguration
	// Renders a Fluent Bit configuration file for the plugin.
	//
	// Returns: A rendered plugin configuration file.
	RenderConfigFile(config *map[string]interface{}) *string
}

// The jsii proxy struct for FluentBitNestFilter
type jsiiProxy_FluentBitNestFilter struct {
	jsiiProxy_FluentBitFilterPluginBase
}

func (j *jsiiProxy_FluentBitNestFilter) AddPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"addPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitNestFilter) Match() FluentBitMatch {
	var returns FluentBitMatch
	_jsii_.Get(
		j,
		"match",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitNestFilter) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitNestFilter) Operation() NestFilterOperation {
	var returns NestFilterOperation
	_jsii_.Get(
		j,
		"operation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitNestFilter) PluginType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pluginType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitNestFilter) RemovePrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"removePrefix",
		&returns,
	)
	return returns
}


// Creates a new instance of the FluentBitNestFilter class.
func NewFluentBitNestFilter(options *FluentBitNestFilterOptions) FluentBitNestFilter {
	_init_.Initialize()

	if err := validateNewFluentBitNestFilterParameters(options); err != nil {
		panic(err)
	}
	j := jsiiProxy_FluentBitNestFilter{}

	_jsii_.Create(
		"cdk-extensions.k8s_aws.FluentBitNestFilter",
		[]interface{}{options},
		&j,
	)

	return &j
}

// Creates a new instance of the FluentBitNestFilter class.
func NewFluentBitNestFilter_Override(f FluentBitNestFilter, options *FluentBitNestFilterOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.k8s_aws.FluentBitNestFilter",
		[]interface{}{options},
		f,
	)
}

func (f *jsiiProxy_FluentBitNestFilter) Bind(_scope constructs.IConstruct) *ResolvedFluentBitConfiguration {
	if err := f.validateBindParameters(_scope); err != nil {
		panic(err)
	}
	var returns *ResolvedFluentBitConfiguration

	_jsii_.Invoke(
		f,
		"bind",
		[]interface{}{_scope},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FluentBitNestFilter) RenderConfigFile(config *map[string]interface{}) *string {
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

