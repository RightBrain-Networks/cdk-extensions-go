package k8saws

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/constructs-go/constructs/v10"
)

// A Fluent Bit filter that allows parsing of fields in event records.
type FluentBitParserFilter interface {
	FluentBitFilterPluginBase
	// Specify field name in record to parse.
	KeyName() *string
	// The pattern to match for records that this output should apply to.
	Match() FluentBitMatch
	// The name of the fluent bit plugin.
	Name() *string
	// Collection of the parsers that should be used to evaluate the filter.
	Parsers() *[]IFluentBitParserPlugin
	// The type of fluent bit plugin.
	PluginType() *string
	// Keep original `keyName` field in the parsed result.
	//
	// If `false`, the field will be removed.
	// Default: false.
	//
	PreserveKey() *bool
	// Keep all other original fields in the parsed result.
	//
	// If `false`, all other original fields will be removed.
	// Default: false.
	//
	ReserveData() *bool
	// Adds a new parser to apply to matched log entries.
	//
	// Returns: The parser filter that the parser plugin was registered with.
	AddParser(parser IFluentBitParserPlugin) FluentBitParserFilter
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

// The jsii proxy struct for FluentBitParserFilter
type jsiiProxy_FluentBitParserFilter struct {
	jsiiProxy_FluentBitFilterPluginBase
}

func (j *jsiiProxy_FluentBitParserFilter) KeyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitParserFilter) Match() FluentBitMatch {
	var returns FluentBitMatch
	_jsii_.Get(
		j,
		"match",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitParserFilter) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitParserFilter) Parsers() *[]IFluentBitParserPlugin {
	var returns *[]IFluentBitParserPlugin
	_jsii_.Get(
		j,
		"parsers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitParserFilter) PluginType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pluginType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitParserFilter) PreserveKey() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"preserveKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitParserFilter) ReserveData() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"reserveData",
		&returns,
	)
	return returns
}


// Creates a new instance of the FluentBitParserFilter class.
func NewFluentBitParserFilter(options *FluentBitParserFilterOptions) FluentBitParserFilter {
	_init_.Initialize()

	if err := validateNewFluentBitParserFilterParameters(options); err != nil {
		panic(err)
	}
	j := jsiiProxy_FluentBitParserFilter{}

	_jsii_.Create(
		"cdk-extensions.k8s_aws.FluentBitParserFilter",
		[]interface{}{options},
		&j,
	)

	return &j
}

// Creates a new instance of the FluentBitParserFilter class.
func NewFluentBitParserFilter_Override(f FluentBitParserFilter, options *FluentBitParserFilterOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.k8s_aws.FluentBitParserFilter",
		[]interface{}{options},
		f,
	)
}

func (f *jsiiProxy_FluentBitParserFilter) AddParser(parser IFluentBitParserPlugin) FluentBitParserFilter {
	if err := f.validateAddParserParameters(parser); err != nil {
		panic(err)
	}
	var returns FluentBitParserFilter

	_jsii_.Invoke(
		f,
		"addParser",
		[]interface{}{parser},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FluentBitParserFilter) Bind(_scope constructs.IConstruct) *ResolvedFluentBitConfiguration {
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

func (f *jsiiProxy_FluentBitParserFilter) RenderConfigFile(config *map[string]interface{}) *string {
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

