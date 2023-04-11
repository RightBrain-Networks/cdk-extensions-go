package k8saws

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/constructs-go/constructs/v10"
)

// A Fluent Bit filter that allows changing records using rules and conditions.
type FluentBitModifyFilter interface {
	FluentBitFilterPluginBase
	// Collection of conditions to apply for the filter.
	Conditions() *[]ModifyCondition
	// The pattern to match for records that this output should apply to.
	Match() FluentBitMatch
	// The name of the fluent bit plugin.
	Name() *string
	// Collection of operations to apply for the filter.
	Operations() *[]ModifyOperation
	// The type of fluent bit plugin.
	PluginType() *string
	// Adds a new condition to the modify filter.
	//
	// All conditions must evaluate to `true` in order for operations are
	// performed.
	//
	// If one or more conditions do not evaluate to true, no conditions are
	// performed.
	//
	// Returns: The modify filter to which the condition was added.
	AddCondition(condition ModifyCondition) FluentBitModifyFilter
	// Adds a new operation to the modify filter.
	//
	// Returns: The modify filter to which the operation was added.
	AddOperation(operation ModifyOperation) FluentBitModifyFilter
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

// The jsii proxy struct for FluentBitModifyFilter
type jsiiProxy_FluentBitModifyFilter struct {
	jsiiProxy_FluentBitFilterPluginBase
}

func (j *jsiiProxy_FluentBitModifyFilter) Conditions() *[]ModifyCondition {
	var returns *[]ModifyCondition
	_jsii_.Get(
		j,
		"conditions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitModifyFilter) Match() FluentBitMatch {
	var returns FluentBitMatch
	_jsii_.Get(
		j,
		"match",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitModifyFilter) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitModifyFilter) Operations() *[]ModifyOperation {
	var returns *[]ModifyOperation
	_jsii_.Get(
		j,
		"operations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitModifyFilter) PluginType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pluginType",
		&returns,
	)
	return returns
}


// Creates a new instance of the FluentBitModifyFilter class.
func NewFluentBitModifyFilter(options *FluentBitModifyFilterOptions) FluentBitModifyFilter {
	_init_.Initialize()

	if err := validateNewFluentBitModifyFilterParameters(options); err != nil {
		panic(err)
	}
	j := jsiiProxy_FluentBitModifyFilter{}

	_jsii_.Create(
		"cdk-extensions.k8s_aws.FluentBitModifyFilter",
		[]interface{}{options},
		&j,
	)

	return &j
}

// Creates a new instance of the FluentBitModifyFilter class.
func NewFluentBitModifyFilter_Override(f FluentBitModifyFilter, options *FluentBitModifyFilterOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.k8s_aws.FluentBitModifyFilter",
		[]interface{}{options},
		f,
	)
}

func (f *jsiiProxy_FluentBitModifyFilter) AddCondition(condition ModifyCondition) FluentBitModifyFilter {
	if err := f.validateAddConditionParameters(condition); err != nil {
		panic(err)
	}
	var returns FluentBitModifyFilter

	_jsii_.Invoke(
		f,
		"addCondition",
		[]interface{}{condition},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FluentBitModifyFilter) AddOperation(operation ModifyOperation) FluentBitModifyFilter {
	if err := f.validateAddOperationParameters(operation); err != nil {
		panic(err)
	}
	var returns FluentBitModifyFilter

	_jsii_.Invoke(
		f,
		"addOperation",
		[]interface{}{operation},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FluentBitModifyFilter) Bind(_scope constructs.IConstruct) *ResolvedFluentBitConfiguration {
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

func (f *jsiiProxy_FluentBitModifyFilter) RenderConfigFile(config *map[string]interface{}) *string {
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

