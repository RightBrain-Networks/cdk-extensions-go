package k8saws

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/constructs-go/constructs/v10"
)

// A Fluent Bit filter that allows appending fields or excluding specific fields.
type FluentBitRecordModifierFilter interface {
	FluentBitFilterPluginBase
	// Collection of tags that are allowed on a matched input record.
	//
	// If a tag is not matched it is removed.
	Allow() *[]*string
	// The pattern to match for records that this output should apply to.
	Match() FluentBitMatch
	// The name of the fluent bit plugin.
	Name() *string
	// The type of fluent bit plugin.
	PluginType() *string
	// Collection of the records to be appending to matched input.
	Records() *[]*AppendedRecord
	// Collection of tags to exclude from a matched input record.
	//
	// If a tag is matched it is removed.
	Remove() *[]*string
	// Adds a tag to be allowed on a matched input record.
	//
	// If a tag is not matched it is removed.
	//
	// Returns: The record modifier filter that the tag plugin was registered
	// with.
	AddAllow(tag *string) FluentBitRecordModifierFilter
	// Add a record to be appended to matched events.
	//
	// Returns: The record modifier filter that the tag plugin was registered
	// with.
	AddRecord(record *AppendedRecord) FluentBitRecordModifierFilter
	// Adds a tag to be removed on a matched input record.
	//
	// If a tag is matched it is removed.
	//
	// Returns: The record modifier filter that the tag plugin was registered
	// with.
	AddRemove(tag *string) FluentBitRecordModifierFilter
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

// The jsii proxy struct for FluentBitRecordModifierFilter
type jsiiProxy_FluentBitRecordModifierFilter struct {
	jsiiProxy_FluentBitFilterPluginBase
}

func (j *jsiiProxy_FluentBitRecordModifierFilter) Allow() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitRecordModifierFilter) Match() FluentBitMatch {
	var returns FluentBitMatch
	_jsii_.Get(
		j,
		"match",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitRecordModifierFilter) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitRecordModifierFilter) PluginType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pluginType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitRecordModifierFilter) Records() *[]*AppendedRecord {
	var returns *[]*AppendedRecord
	_jsii_.Get(
		j,
		"records",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitRecordModifierFilter) Remove() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"remove",
		&returns,
	)
	return returns
}


// Creates a new instance of the FluentBitRecordModifierFilter class.
func NewFluentBitRecordModifierFilter(options *FluentBitRecordModifierFilterOptions) FluentBitRecordModifierFilter {
	_init_.Initialize()

	if err := validateNewFluentBitRecordModifierFilterParameters(options); err != nil {
		panic(err)
	}
	j := jsiiProxy_FluentBitRecordModifierFilter{}

	_jsii_.Create(
		"cdk-extensions.k8s_aws.FluentBitRecordModifierFilter",
		[]interface{}{options},
		&j,
	)

	return &j
}

// Creates a new instance of the FluentBitRecordModifierFilter class.
func NewFluentBitRecordModifierFilter_Override(f FluentBitRecordModifierFilter, options *FluentBitRecordModifierFilterOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.k8s_aws.FluentBitRecordModifierFilter",
		[]interface{}{options},
		f,
	)
}

func (f *jsiiProxy_FluentBitRecordModifierFilter) AddAllow(tag *string) FluentBitRecordModifierFilter {
	if err := f.validateAddAllowParameters(tag); err != nil {
		panic(err)
	}
	var returns FluentBitRecordModifierFilter

	_jsii_.Invoke(
		f,
		"addAllow",
		[]interface{}{tag},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FluentBitRecordModifierFilter) AddRecord(record *AppendedRecord) FluentBitRecordModifierFilter {
	if err := f.validateAddRecordParameters(record); err != nil {
		panic(err)
	}
	var returns FluentBitRecordModifierFilter

	_jsii_.Invoke(
		f,
		"addRecord",
		[]interface{}{record},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FluentBitRecordModifierFilter) AddRemove(tag *string) FluentBitRecordModifierFilter {
	if err := f.validateAddRemoveParameters(tag); err != nil {
		panic(err)
	}
	var returns FluentBitRecordModifierFilter

	_jsii_.Invoke(
		f,
		"addRemove",
		[]interface{}{tag},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FluentBitRecordModifierFilter) Bind(_scope constructs.IConstruct) *ResolvedFluentBitConfiguration {
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

func (f *jsiiProxy_FluentBitRecordModifierFilter) RenderConfigFile(config *map[string]interface{}) *string {
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

