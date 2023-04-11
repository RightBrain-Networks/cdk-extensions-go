package k8saws

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/constructs-go/constructs/v10"
)

// A Fluent Bit filter that parsed inbound messages in LTSV format.
type FluentBitLogfmtParser interface {
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

// The jsii proxy struct for FluentBitLogfmtParser
type jsiiProxy_FluentBitLogfmtParser struct {
	jsiiProxy_FluentBitParserPluginBase
}

func (j *jsiiProxy_FluentBitLogfmtParser) Format() *string {
	var returns *string
	_jsii_.Get(
		j,
		"format",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitLogfmtParser) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitLogfmtParser) PluginType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pluginType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitLogfmtParser) TimeFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitLogfmtParser) TimeKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitLogfmtParser) Types() *map[string]ParserPluginDataType {
	var returns *map[string]ParserPluginDataType
	_jsii_.Get(
		j,
		"types",
		&returns,
	)
	return returns
}


// Creates a new instance of the FluentBitLogfmtParser class.
func NewFluentBitLogfmtParser(name *string, options *FluentBitLogfmtParserOptions) FluentBitLogfmtParser {
	_init_.Initialize()

	if err := validateNewFluentBitLogfmtParserParameters(name, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_FluentBitLogfmtParser{}

	_jsii_.Create(
		"cdk-extensions.k8s_aws.FluentBitLogfmtParser",
		[]interface{}{name, options},
		&j,
	)

	return &j
}

// Creates a new instance of the FluentBitLogfmtParser class.
func NewFluentBitLogfmtParser_Override(f FluentBitLogfmtParser, name *string, options *FluentBitLogfmtParserOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.k8s_aws.FluentBitLogfmtParser",
		[]interface{}{name, options},
		f,
	)
}

func (f *jsiiProxy_FluentBitLogfmtParser) Bind(_scope constructs.IConstruct) *ResolvedFluentBitConfiguration {
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

func (f *jsiiProxy_FluentBitLogfmtParser) RenderConfigFile(config *map[string]interface{}) *string {
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

