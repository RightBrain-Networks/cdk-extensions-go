package k8saws

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskinesis"
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents configuration for outputing logs from Fluent Bit to Kinesis Data Streams.
type FluentBitKinesisOutput interface {
	FluentBitOutputPluginBase
	// Immediately retry failed requests to AWS services once.
	//
	// This option does
	// not affect the normal Fluent Bit retry mechanism with backoff. Instead,
	// it enables an immediate retry with no delay for networking errors, which
	// may help improve throughput when there are transient/random networking
	// issues.
	AutoRetryRequests() *bool
	// Specify a custom endpoint for the Firehose API.
	Endpoint() *string
	// By default, the whole log record will be sent to Firehose.
	//
	// If you
	// specify a key name with this option, then only the value of that key
	// will be sent to Firehose.
	LogKey() *string
	// The pattern to match for records that this output should apply to.
	Match() FluentBitMatch
	// The name of the fluent bit plugin.
	Name() *string
	// The type of fluent bit plugin.
	PluginType() *string
	// The AWS region.
	Region() *string
	// ARN of an IAM role to assume (for cross account access).
	Role() awsiam.IRole
	// The name of the Kinesis Streams Delivery stream that you want log records sent to.
	Stream() awskinesis.IStream
	// Specify a custom STS endpoint for the AWS STS API.
	StsEndpoint() *string
	// Add the timestamp to the record under this key.
	TimeKey() *string
	// A strftime compliant format string for the timestamp.
	TimeKeyFormat() *string
	// Builds a configuration for this plugin and returns the details for consumtion by a resource that is configuring logging.
	//
	// Returns: A configuration for the plugin that con be used by the resource
	// configuring logging.
	Bind(scope constructs.IConstruct) *ResolvedFluentBitConfiguration
	// Renders a Fluent Bit configuration file for the plugin.
	//
	// Returns: A rendered plugin configuration file.
	RenderConfigFile(config *map[string]interface{}) *string
}

// The jsii proxy struct for FluentBitKinesisOutput
type jsiiProxy_FluentBitKinesisOutput struct {
	jsiiProxy_FluentBitOutputPluginBase
}

func (j *jsiiProxy_FluentBitKinesisOutput) AutoRetryRequests() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"autoRetryRequests",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitKinesisOutput) Endpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitKinesisOutput) LogKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitKinesisOutput) Match() FluentBitMatch {
	var returns FluentBitMatch
	_jsii_.Get(
		j,
		"match",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitKinesisOutput) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitKinesisOutput) PluginType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pluginType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitKinesisOutput) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitKinesisOutput) Role() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitKinesisOutput) Stream() awskinesis.IStream {
	var returns awskinesis.IStream
	_jsii_.Get(
		j,
		"stream",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitKinesisOutput) StsEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stsEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitKinesisOutput) TimeKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitKinesisOutput) TimeKeyFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeKeyFormat",
		&returns,
	)
	return returns
}


// Creates a new instance of the FluentBitKinesisOutput class.
func NewFluentBitKinesisOutput(options *FluentBitKinesisOutputOptions) FluentBitKinesisOutput {
	_init_.Initialize()

	if err := validateNewFluentBitKinesisOutputParameters(options); err != nil {
		panic(err)
	}
	j := jsiiProxy_FluentBitKinesisOutput{}

	_jsii_.Create(
		"cdk-extensions.k8s_aws.FluentBitKinesisOutput",
		[]interface{}{options},
		&j,
	)

	return &j
}

// Creates a new instance of the FluentBitKinesisOutput class.
func NewFluentBitKinesisOutput_Override(f FluentBitKinesisOutput, options *FluentBitKinesisOutputOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.k8s_aws.FluentBitKinesisOutput",
		[]interface{}{options},
		f,
	)
}

func (f *jsiiProxy_FluentBitKinesisOutput) Bind(scope constructs.IConstruct) *ResolvedFluentBitConfiguration {
	if err := f.validateBindParameters(scope); err != nil {
		panic(err)
	}
	var returns *ResolvedFluentBitConfiguration

	_jsii_.Invoke(
		f,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FluentBitKinesisOutput) RenderConfigFile(config *map[string]interface{}) *string {
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

