package k8saws

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/constructs-go/constructs/v10"
)

// A Fluent Bit filter that parsed inbound messages in JSON format.
type FluentBitJsonParser interface {
	FluentBitParserPluginBase
	// The data format that the parser extracts records from.
	Format() *string
	// The name of the fluent bit plugin.
	Name() *string
	// The type of fluent bit plugin.
	PluginType() *string
	// Defines the format of the timestamp on the inbound record.
	// See: [strftime](http://man7.org/linux/man-pages/man3/strftime.3.html)
	//
	TimeFormat() *string
	// The key under which timestamp information for the inbound record is given.
	TimeKey() *string
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

// The jsii proxy struct for FluentBitJsonParser
type jsiiProxy_FluentBitJsonParser struct {
	jsiiProxy_FluentBitParserPluginBase
}

func (j *jsiiProxy_FluentBitJsonParser) Format() *string {
	var returns *string
	_jsii_.Get(
		j,
		"format",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitJsonParser) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitJsonParser) PluginType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pluginType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitJsonParser) TimeFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitJsonParser) TimeKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeKey",
		&returns,
	)
	return returns
}


// Creates a new instance of the FluentBitJsonParser class.
func NewFluentBitJsonParser(name *string, options *FluentBitJsonParserOptions) FluentBitJsonParser {
	_init_.Initialize()

	if err := validateNewFluentBitJsonParserParameters(name, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_FluentBitJsonParser{}

	_jsii_.Create(
		"cdk-extensions.k8s_aws.FluentBitJsonParser",
		[]interface{}{name, options},
		&j,
	)

	return &j
}

// Creates a new instance of the FluentBitJsonParser class.
func NewFluentBitJsonParser_Override(f FluentBitJsonParser, name *string, options *FluentBitJsonParserOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.k8s_aws.FluentBitJsonParser",
		[]interface{}{name, options},
		f,
	)
}

func (f *jsiiProxy_FluentBitJsonParser) Bind(_scope constructs.IConstruct) *ResolvedFluentBitConfiguration {
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

func (f *jsiiProxy_FluentBitJsonParser) RenderConfigFile(config *map[string]interface{}) *string {
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

