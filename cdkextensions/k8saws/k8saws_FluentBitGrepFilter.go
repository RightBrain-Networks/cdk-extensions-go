package k8saws

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/constructs-go/constructs/v10"
)

// A Fluent Bit filter that allows log records to be kept or discarded based on whether they match a given regular expression or not.
type FluentBitGrepFilter interface {
	FluentBitFilterPluginBase
	// The pattern to match for records that this output should apply to.
	Match() FluentBitMatch
	// The name of the fluent bit plugin.
	Name() *string
	// The pattern to use for filtering records processed by the plugin.
	Pattern() *FluentBitGrepRegex
	// The type of fluent bit plugin.
	PluginType() *string
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

// The jsii proxy struct for FluentBitGrepFilter
type jsiiProxy_FluentBitGrepFilter struct {
	jsiiProxy_FluentBitFilterPluginBase
}

func (j *jsiiProxy_FluentBitGrepFilter) Match() FluentBitMatch {
	var returns FluentBitMatch
	_jsii_.Get(
		j,
		"match",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitGrepFilter) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitGrepFilter) Pattern() *FluentBitGrepRegex {
	var returns *FluentBitGrepRegex
	_jsii_.Get(
		j,
		"pattern",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitGrepFilter) PluginType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pluginType",
		&returns,
	)
	return returns
}


// Creates a new instance of the FluentBitKinesisFirehoseOutput class.
func NewFluentBitGrepFilter(options *FluentBitGrepFilterOptions) FluentBitGrepFilter {
	_init_.Initialize()

	if err := validateNewFluentBitGrepFilterParameters(options); err != nil {
		panic(err)
	}
	j := jsiiProxy_FluentBitGrepFilter{}

	_jsii_.Create(
		"cdk-extensions.k8s_aws.FluentBitGrepFilter",
		[]interface{}{options},
		&j,
	)

	return &j
}

// Creates a new instance of the FluentBitKinesisFirehoseOutput class.
func NewFluentBitGrepFilter_Override(f FluentBitGrepFilter, options *FluentBitGrepFilterOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.k8s_aws.FluentBitGrepFilter",
		[]interface{}{options},
		f,
	)
}

func (f *jsiiProxy_FluentBitGrepFilter) Bind(_scope constructs.IConstruct) *ResolvedFluentBitConfiguration {
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

func (f *jsiiProxy_FluentBitGrepFilter) RenderConfigFile(config *map[string]interface{}) *string {
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

