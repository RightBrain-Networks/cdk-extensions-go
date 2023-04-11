package k8saws

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
)

// A Fluent Bit filter that sets the average *Rate* of messages per *Interval*, based on leaky bucket and sliding window algorithm.
//
// In case of overflood,
// it will leak within certain rate.
type FluentBitThrottleFilter interface {
	FluentBitFilterPluginBase
	// Time interval.
	Interval() awscdk.Duration
	// The pattern to match for records that this output should apply to.
	Match() FluentBitMatch
	// The name of the fluent bit plugin.
	Name() *string
	// The type of fluent bit plugin.
	PluginType() *string
	// Whether to print status messages with current rate and the limits to information logs.
	PrintStatus() *bool
	// Amount of messages for the time.
	Rate() *float64
	// Amount of intervals to calculate average over.
	Window() *float64
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

// The jsii proxy struct for FluentBitThrottleFilter
type jsiiProxy_FluentBitThrottleFilter struct {
	jsiiProxy_FluentBitFilterPluginBase
}

func (j *jsiiProxy_FluentBitThrottleFilter) Interval() awscdk.Duration {
	var returns awscdk.Duration
	_jsii_.Get(
		j,
		"interval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitThrottleFilter) Match() FluentBitMatch {
	var returns FluentBitMatch
	_jsii_.Get(
		j,
		"match",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitThrottleFilter) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitThrottleFilter) PluginType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pluginType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitThrottleFilter) PrintStatus() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"printStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitThrottleFilter) Rate() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitThrottleFilter) Window() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"window",
		&returns,
	)
	return returns
}


// Creates a new instance of the FluentBitThrottleFilter class.
func NewFluentBitThrottleFilter(options *FluentBitThrottleFilterOptions) FluentBitThrottleFilter {
	_init_.Initialize()

	if err := validateNewFluentBitThrottleFilterParameters(options); err != nil {
		panic(err)
	}
	j := jsiiProxy_FluentBitThrottleFilter{}

	_jsii_.Create(
		"cdk-extensions.k8s_aws.FluentBitThrottleFilter",
		[]interface{}{options},
		&j,
	)

	return &j
}

// Creates a new instance of the FluentBitThrottleFilter class.
func NewFluentBitThrottleFilter_Override(f FluentBitThrottleFilter, options *FluentBitThrottleFilterOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.k8s_aws.FluentBitThrottleFilter",
		[]interface{}{options},
		f,
	)
}

func (f *jsiiProxy_FluentBitThrottleFilter) Bind(_scope constructs.IConstruct) *ResolvedFluentBitConfiguration {
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

func (f *jsiiProxy_FluentBitThrottleFilter) RenderConfigFile(config *map[string]interface{}) *string {
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

