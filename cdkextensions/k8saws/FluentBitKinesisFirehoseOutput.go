package k8saws

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/vibe-io/cdk-extensions-go/cdkextensions/kinesisfirehose"
)

// Represents configuration for outputing logs from Fluent Bit to Kinesis Firehose.
type FluentBitKinesisFirehoseOutput interface {
	FluentBitOutputPluginBase
	// Immediately retry failed requests to AWS services once.
	//
	// This option does
	// not affect the normal Fluent Bit retry mechanism with backoff. Instead,
	// it enables an immediate retry with no delay for networking errors, which
	// may help improve throughput when there are transient/random networking
	// issues.
	AutoRetryRequests() *bool
	// Compression type for Firehose records.
	//
	// Each log record is individually
	// compressed and sent to Firehose.
	Compression() KinesisFirehoseCompressionFormat
	// The Kinesis Firehose Delivery stream that you want log records sent to.
	DeliveryStream() kinesisfirehose.IDeliveryStream
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

// The jsii proxy struct for FluentBitKinesisFirehoseOutput
type jsiiProxy_FluentBitKinesisFirehoseOutput struct {
	jsiiProxy_FluentBitOutputPluginBase
}

func (j *jsiiProxy_FluentBitKinesisFirehoseOutput) AutoRetryRequests() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"autoRetryRequests",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitKinesisFirehoseOutput) Compression() KinesisFirehoseCompressionFormat {
	var returns KinesisFirehoseCompressionFormat
	_jsii_.Get(
		j,
		"compression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitKinesisFirehoseOutput) DeliveryStream() kinesisfirehose.IDeliveryStream {
	var returns kinesisfirehose.IDeliveryStream
	_jsii_.Get(
		j,
		"deliveryStream",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitKinesisFirehoseOutput) Endpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitKinesisFirehoseOutput) LogKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitKinesisFirehoseOutput) Match() FluentBitMatch {
	var returns FluentBitMatch
	_jsii_.Get(
		j,
		"match",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitKinesisFirehoseOutput) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitKinesisFirehoseOutput) PluginType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pluginType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitKinesisFirehoseOutput) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitKinesisFirehoseOutput) Role() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitKinesisFirehoseOutput) StsEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stsEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitKinesisFirehoseOutput) TimeKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitKinesisFirehoseOutput) TimeKeyFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeKeyFormat",
		&returns,
	)
	return returns
}


// Creates a new instance of the FluentBitKinesisFirehoseOutput class.
func NewFluentBitKinesisFirehoseOutput(options *FluentBitKinesisFirehoseOutputOptions) FluentBitKinesisFirehoseOutput {
	_init_.Initialize()

	if err := validateNewFluentBitKinesisFirehoseOutputParameters(options); err != nil {
		panic(err)
	}
	j := jsiiProxy_FluentBitKinesisFirehoseOutput{}

	_jsii_.Create(
		"cdk-extensions.k8s_aws.FluentBitKinesisFirehoseOutput",
		[]interface{}{options},
		&j,
	)

	return &j
}

// Creates a new instance of the FluentBitKinesisFirehoseOutput class.
func NewFluentBitKinesisFirehoseOutput_Override(f FluentBitKinesisFirehoseOutput, options *FluentBitKinesisFirehoseOutputOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.k8s_aws.FluentBitKinesisFirehoseOutput",
		[]interface{}{options},
		f,
	)
}

func (f *jsiiProxy_FluentBitKinesisFirehoseOutput) Bind(scope constructs.IConstruct) *ResolvedFluentBitConfiguration {
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

func (f *jsiiProxy_FluentBitKinesisFirehoseOutput) RenderConfigFile(config *map[string]interface{}) *string {
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

