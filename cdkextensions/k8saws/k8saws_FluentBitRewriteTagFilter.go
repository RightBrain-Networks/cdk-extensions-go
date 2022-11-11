package k8saws

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/vibe-io/cdk-extensions-go/cdkextensions/core"
)

// A Fluent Bit filter that allows parsing of fields in event records.
type FluentBitRewriteTagFilter interface {
	FluentBitFilterPluginBase
	// Set a limit on the amount of memory the tag rewrite emitter can consume if the outputs provide backpressure.
	EmitterMemBufLimit() core.DataSize
	// When the filter emits a record under the new Tag, there is an internal emitter plugin that takes care of the job.
	//
	// Since this emitter expose
	// metrics as any other component of the pipeline, you can use this
	// property to configure an optional name for it.
	EmitterName() *string
	// Define a buffering mechanism for the new records created.
	//
	// Note these records are part of the emitter plugin.
	EmitterStorageType() EmitterStorageType
	// The pattern to match for records that this output should apply to.
	Match() FluentBitMatch
	// The name of the fluent bit plugin.
	Name() *string
	// The type of fluent bit plugin.
	PluginType() *string
	// Collection of rules defining matching criteria and the format of the tag for the matching record.
	Rules() *[]*RewriteTagRule
	// Adds a new rule to apply to matched log entries.
	//
	// Returns: The parser filter that the parser plugin was registered with.
	AddRule(rule *RewriteTagRule) FluentBitRewriteTagFilter
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

// The jsii proxy struct for FluentBitRewriteTagFilter
type jsiiProxy_FluentBitRewriteTagFilter struct {
	jsiiProxy_FluentBitFilterPluginBase
}

func (j *jsiiProxy_FluentBitRewriteTagFilter) EmitterMemBufLimit() core.DataSize {
	var returns core.DataSize
	_jsii_.Get(
		j,
		"emitterMemBufLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitRewriteTagFilter) EmitterName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emitterName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitRewriteTagFilter) EmitterStorageType() EmitterStorageType {
	var returns EmitterStorageType
	_jsii_.Get(
		j,
		"emitterStorageType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitRewriteTagFilter) Match() FluentBitMatch {
	var returns FluentBitMatch
	_jsii_.Get(
		j,
		"match",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitRewriteTagFilter) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitRewriteTagFilter) PluginType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pluginType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitRewriteTagFilter) Rules() *[]*RewriteTagRule {
	var returns *[]*RewriteTagRule
	_jsii_.Get(
		j,
		"rules",
		&returns,
	)
	return returns
}


// Creates a new instance of the FluentBitRewriteTagFilter class.
func NewFluentBitRewriteTagFilter(options *FluentBitRewriteTagFilterOptions) FluentBitRewriteTagFilter {
	_init_.Initialize()

	if err := validateNewFluentBitRewriteTagFilterParameters(options); err != nil {
		panic(err)
	}
	j := jsiiProxy_FluentBitRewriteTagFilter{}

	_jsii_.Create(
		"cdk-extensions.k8s_aws.FluentBitRewriteTagFilter",
		[]interface{}{options},
		&j,
	)

	return &j
}

// Creates a new instance of the FluentBitRewriteTagFilter class.
func NewFluentBitRewriteTagFilter_Override(f FluentBitRewriteTagFilter, options *FluentBitRewriteTagFilterOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.k8s_aws.FluentBitRewriteTagFilter",
		[]interface{}{options},
		f,
	)
}

func FluentBitRewriteTagFilter_PLUGIN_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"cdk-extensions.k8s_aws.FluentBitRewriteTagFilter",
		"PLUGIN_NAME",
		&returns,
	)
	return returns
}

func (f *jsiiProxy_FluentBitRewriteTagFilter) AddRule(rule *RewriteTagRule) FluentBitRewriteTagFilter {
	if err := f.validateAddRuleParameters(rule); err != nil {
		panic(err)
	}
	var returns FluentBitRewriteTagFilter

	_jsii_.Invoke(
		f,
		"addRule",
		[]interface{}{rule},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FluentBitRewriteTagFilter) Bind(_scope constructs.IConstruct) *ResolvedFluentBitConfiguration {
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

func (f *jsiiProxy_FluentBitRewriteTagFilter) RenderConfigFile(config *map[string]interface{}) *string {
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

