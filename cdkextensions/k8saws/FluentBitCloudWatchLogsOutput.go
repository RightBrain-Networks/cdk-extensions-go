package k8saws

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents configuration for outputing logs from Fluent Bit to CloudWatch Logs.
type FluentBitCloudWatchLogsOutput interface {
	FluentBitOutputPluginBase
	// Automatically create the log group.
	AutoCreateGroup() *bool
	// Immediately retry failed requests to AWS services once.
	//
	// This option does
	// not affect the normal Fluent Bit retry mechanism with backoff. Instead,
	// it enables an immediate retry with no delay for networking errors, which
	// may help improve throughput when there are transient/random networking
	// issues.
	AutoRetryRequests() *bool
	// Specify a custom endpoint for the CloudWatch Logs API.
	Endpoint() *string
	// An optional parameter that can be used to tell CloudWatch the format of the data.
	//
	// A value of json/emf enables CloudWatch to extract custom
	// metrics embedded in a JSON payload.
	// See: [Embedded Metric Format](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Embedded_Metric_Format_Specification.html)
	//
	LogFormat() *string
	// The CloudWatch Log Group configuration for output records.
	LogGroup() FluentBitLogGroupOutput
	// Template for Log Group name using Fluent Bit record_accessor syntax.
	//
	// This field is optional and if configured it overrides the configured Log
	// Group.
	//
	// If the template translation fails, an error is logged and the provided
	// Log Group (which is still required) is used instead.
	// See: [Fluent Bit record accessor snytax](https://docs.fluentbit.io/manual/administration/configuring-fluent-bit/classic-mode/record-accessor)
	//
	LogGroupTemplate() *string
	// By default, the whole log record will be sent to CloudWatch.
	//
	// If you
	// specify a key name with this option, then only the value of that key
	// will be sent to CloudWatch.
	LogKey() *string
	// If set to a number greater than zero, and newly create log group's retention policy is set to this many days.
	LogRetention() awslogs.RetentionDays
	// The CloudWatch LogStream configuration for outbound records.
	LogStream() FluentBitLogStreamOutput
	// Template for Log Stream name using Fluent Bit record accessor syntax.
	//
	// This field is optional and if configured it overrides the other log
	// stream options. If the template translation fails, an error is logged
	// and the logStream or logStreamPrefix are used instead (and thus one of
	// those fields is still required to be configured).
	// See: [Fluent Bit record accessor snytax](https://docs.fluentbit.io/manual/administration/configuring-fluent-bit/classic-mode/record-accessor)
	//
	LogStreamTemplate() *string
	// The pattern to match for records that this output should apply to.
	Match() FluentBitMatch
	// A list of lists containing the dimension keys that will be applied to all metrics.
	//
	// The values within a dimension set MUST also be members on
	// the root-node.
	// See: [Dimensions](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/cloudwatch_concepts.html#Dimension)
	//
	MetricDimensions() *[]*string
	// An optional string representing the CloudWatch namespace for the metrics.
	// See: [Metric Tutorial](https://docs.fluentbit.io/manual/pipeline/outputs/cloudwatch#metrics-tutorial)
	//
	MetricNamespace() *string
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

// The jsii proxy struct for FluentBitCloudWatchLogsOutput
type jsiiProxy_FluentBitCloudWatchLogsOutput struct {
	jsiiProxy_FluentBitOutputPluginBase
}

func (j *jsiiProxy_FluentBitCloudWatchLogsOutput) AutoCreateGroup() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"autoCreateGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitCloudWatchLogsOutput) AutoRetryRequests() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"autoRetryRequests",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitCloudWatchLogsOutput) Endpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitCloudWatchLogsOutput) LogFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitCloudWatchLogsOutput) LogGroup() FluentBitLogGroupOutput {
	var returns FluentBitLogGroupOutput
	_jsii_.Get(
		j,
		"logGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitCloudWatchLogsOutput) LogGroupTemplate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logGroupTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitCloudWatchLogsOutput) LogKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitCloudWatchLogsOutput) LogRetention() awslogs.RetentionDays {
	var returns awslogs.RetentionDays
	_jsii_.Get(
		j,
		"logRetention",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitCloudWatchLogsOutput) LogStream() FluentBitLogStreamOutput {
	var returns FluentBitLogStreamOutput
	_jsii_.Get(
		j,
		"logStream",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitCloudWatchLogsOutput) LogStreamTemplate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logStreamTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitCloudWatchLogsOutput) Match() FluentBitMatch {
	var returns FluentBitMatch
	_jsii_.Get(
		j,
		"match",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitCloudWatchLogsOutput) MetricDimensions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"metricDimensions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitCloudWatchLogsOutput) MetricNamespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricNamespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitCloudWatchLogsOutput) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitCloudWatchLogsOutput) PluginType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pluginType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitCloudWatchLogsOutput) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitCloudWatchLogsOutput) Role() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitCloudWatchLogsOutput) StsEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stsEndpoint",
		&returns,
	)
	return returns
}


// Creates a new instance of the FluentBitCloudWatchLogsOutput class.
func NewFluentBitCloudWatchLogsOutput(options *FluentBitCloudWatchLogsOutputOptions) FluentBitCloudWatchLogsOutput {
	_init_.Initialize()

	if err := validateNewFluentBitCloudWatchLogsOutputParameters(options); err != nil {
		panic(err)
	}
	j := jsiiProxy_FluentBitCloudWatchLogsOutput{}

	_jsii_.Create(
		"cdk-extensions.k8s_aws.FluentBitCloudWatchLogsOutput",
		[]interface{}{options},
		&j,
	)

	return &j
}

// Creates a new instance of the FluentBitCloudWatchLogsOutput class.
func NewFluentBitCloudWatchLogsOutput_Override(f FluentBitCloudWatchLogsOutput, options *FluentBitCloudWatchLogsOutputOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.k8s_aws.FluentBitCloudWatchLogsOutput",
		[]interface{}{options},
		f,
	)
}

func (f *jsiiProxy_FluentBitCloudWatchLogsOutput) Bind(scope constructs.IConstruct) *ResolvedFluentBitConfiguration {
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

func (f *jsiiProxy_FluentBitCloudWatchLogsOutput) RenderConfigFile(config *map[string]interface{}) *string {
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

