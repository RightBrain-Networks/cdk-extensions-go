package k8saws

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/constructs-go/constructs/v10"
)

// A Fluent Bit filter that parsed inbound messages using regular expressions.
type FluentBitRegexParser interface {
	FluentBitParserPluginBase
	// The data format that the parser extracts records from.
	Format() *string
	// The name of the fluent bit plugin.
	Name() *string
	// The type of fluent bit plugin.
	PluginType() *string
	// The regular expression to use to parse the incoming records.
	//
	// Use regex group names to define the name of fields being captured.
	Regex() *string
	// If enabled, the parser ignores empty value of the record.
	SkipEmptyValues() *bool
	// Defines the format of the timestamp on the inbound record.
	// See: [strftime](http://man7.org/linux/man-pages/man3/strftime.3.html)
	//
	TimeFormat() *string
	// The key under which timestamp information for the inbound record is given.
	TimeKey() *string
	// Maps group names matched by the regex to the data types they should be interpreted as.
	Types() *map[string]ParserPluginDataType
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

// The jsii proxy struct for FluentBitRegexParser
type jsiiProxy_FluentBitRegexParser struct {
	jsiiProxy_FluentBitParserPluginBase
}

func (j *jsiiProxy_FluentBitRegexParser) Format() *string {
	var returns *string
	_jsii_.Get(
		j,
		"format",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitRegexParser) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitRegexParser) PluginType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pluginType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitRegexParser) Regex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitRegexParser) SkipEmptyValues() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"skipEmptyValues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitRegexParser) TimeFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitRegexParser) TimeKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitRegexParser) Types() *map[string]ParserPluginDataType {
	var returns *map[string]ParserPluginDataType
	_jsii_.Get(
		j,
		"types",
		&returns,
	)
	return returns
}


// Creates a new instance of the FluentBitLtsvParser class.
func NewFluentBitRegexParser(name *string, options *FluentBitRegexParserOptions) FluentBitRegexParser {
	_init_.Initialize()

	if err := validateNewFluentBitRegexParserParameters(name, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_FluentBitRegexParser{}

	_jsii_.Create(
		"cdk-extensions.k8s_aws.FluentBitRegexParser",
		[]interface{}{name, options},
		&j,
	)

	return &j
}

// Creates a new instance of the FluentBitLtsvParser class.
func NewFluentBitRegexParser_Override(f FluentBitRegexParser, name *string, options *FluentBitRegexParserOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.k8s_aws.FluentBitRegexParser",
		[]interface{}{name, options},
		f,
	)
}

func (f *jsiiProxy_FluentBitRegexParser) Bind(_scope constructs.IConstruct) *ResolvedFluentBitConfiguration {
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

func (f *jsiiProxy_FluentBitRegexParser) RenderConfigFile(config *map[string]interface{}) *string {
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

