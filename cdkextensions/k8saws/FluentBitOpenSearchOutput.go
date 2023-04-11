package k8saws

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsopensearchservice"
	"github.com/aws/constructs-go/constructs/v10"
)

type FluentBitOpenSearchOutput interface {
	FluentBitOutputPluginBase
	// Enable AWS Sigv4 Authentication for Amazon OpenSearch Service.
	AwsAuth() *bool
	// External ID for the AWS IAM Role specified with `awsRole`.
	AwsExternalId() *string
	// Specify the AWS region for Amazon OpenSearch Service.
	AwsRegion() *string
	// AWS IAM Role to assume to put records to your Amazon cluster.
	AwsRole() awsiam.IRole
	// Specify the custom sts endpoint to be used with STS API for Amazon OpenSearch Service.
	AwsStsEndpoint() *string
	// Specify the buffer size used to read the response from the OpenSearch HTTP service.
	//
	// This option is useful for debugging purposes where is
	// required to read full responses, note that response size grows depending
	// of the number of records inserted.
	BufferSize() OpenSearchOutputBufferSize
	// Use current time for index generation instead of message record.
	CurrentTimeIndex() *bool
	// The Opensearch domain to which logs should be shipped.
	Domain() awsopensearchservice.IDomain
	// When enabled, generate `_id` for outgoing records.
	//
	// This prevents duplicate
	// records when retrying.
	GenerateId() *bool
	// Password for user defined in `httpUser`.
	HttpPasswd() *string
	// Optional username credential for access.
	HttpUser() *string
	// If set, `_id` will be the value of the key from incoming record and `generateId` option is ignored.
	IdKey() *string
	// When enabled, it append the Tag name to the record.
	IncludeTagKey() *bool
	// Index name.
	Index() *string
	// Time format (based on strftime) to generate the second part of the Index name.
	// See: [strftime](http://man7.org/linux/man-pages/man3/strftime.3.html)
	//
	LogstashDateFormat() *string
	// Enable Logstash format compatibility.
	LogstashFormat() *bool
	// When `logstashFormat` is enabled, the Index name is composed using a prefix and the date, e.g: If `logstashPrefix` is equals to 'mydata' your index will become 'mydata-YYYY.MM.DD'.
	//
	// The last string appended belongs to the date when the data is being
	// generated.
	LogstashPrefix() *string
	// When included: the value in the record that belongs to the key will be looked up and over-write the `logstashPrefix` for index generation.
	//
	// If
	// the key/value is not found in the record then the `logstashPrefix` option
	// will act as a fallback.
	//
	// Nested keys are not supported (if desired, you can use the nest filter
	// plugin to remove nesting).
	LogstashPrefixKey() *string
	// The pattern to match for records that this output should apply to.
	Match() FluentBitMatch
	// The name of the fluent bit plugin.
	Name() *string
	// OpenSearch accepts new data on HTTP query path "/_bulk".
	//
	// But it is also
	// possible to serve OpenSearch behind a reverse proxy on a subpath. This
	// option defines such path on the fluent-bit side. It simply adds a path
	// prefix in the indexing HTTP POST URI.
	Path() *string
	// OpenSearch allows to setup filters called pipelines.
	//
	// This option allows
	// to define which pipeline the database should use.
	Pipeline() *string
	// The type of fluent bit plugin.
	PluginType() *string
	// TCP port of the target OpenSearch instance.
	Port() *float64
	// When enabled, replace field name dots with underscore.
	ReplaceDots() *bool
	// When enabled, mapping types is removed and `type` option is ignored.
	SuppressTypeName() *bool
	// When `includeTagKey` is enabled, this property defines the key name for the tag.
	TagKey() *string
	// When `logstashFormat` is enabled, each record will get a new timestamp field.
	//
	// The`timeKey` property defines the name of that field.
	TimeKey() *string
	// When `logstashFormat` is enabled, this property defines the format of the timestamp.
	// See: [strftime](http://man7.org/linux/man-pages/man3/strftime.3.html)
	//
	TimeKeyFormat() *string
	// When `logstashFormat` is enabled, enabling this property sends nanosecond precision timestamps.
	TimeKeyNanos() *bool
	// When enabled print the OpenSearch API calls to stdout when OpenSearch returns an error (for diag only).
	TraceError() *bool
	// When enabled print the OpenSearch API calls to stdout (for diag only).
	TraceOutput() *bool
	// Type name.
	Type() *string
	// Enables dedicated thread(s) for this output.
	Workers() *float64
	// Operation to use to write in bulk requests.
	WriteOperation() *string
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

// The jsii proxy struct for FluentBitOpenSearchOutput
type jsiiProxy_FluentBitOpenSearchOutput struct {
	jsiiProxy_FluentBitOutputPluginBase
}

func (j *jsiiProxy_FluentBitOpenSearchOutput) AwsAuth() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"awsAuth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitOpenSearchOutput) AwsExternalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsExternalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitOpenSearchOutput) AwsRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitOpenSearchOutput) AwsRole() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"awsRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitOpenSearchOutput) AwsStsEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsStsEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitOpenSearchOutput) BufferSize() OpenSearchOutputBufferSize {
	var returns OpenSearchOutputBufferSize
	_jsii_.Get(
		j,
		"bufferSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitOpenSearchOutput) CurrentTimeIndex() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"currentTimeIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitOpenSearchOutput) Domain() awsopensearchservice.IDomain {
	var returns awsopensearchservice.IDomain
	_jsii_.Get(
		j,
		"domain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitOpenSearchOutput) GenerateId() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"generateId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitOpenSearchOutput) HttpPasswd() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpPasswd",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitOpenSearchOutput) HttpUser() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpUser",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitOpenSearchOutput) IdKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitOpenSearchOutput) IncludeTagKey() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"includeTagKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitOpenSearchOutput) Index() *string {
	var returns *string
	_jsii_.Get(
		j,
		"index",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitOpenSearchOutput) LogstashDateFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logstashDateFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitOpenSearchOutput) LogstashFormat() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"logstashFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitOpenSearchOutput) LogstashPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logstashPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitOpenSearchOutput) LogstashPrefixKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logstashPrefixKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitOpenSearchOutput) Match() FluentBitMatch {
	var returns FluentBitMatch
	_jsii_.Get(
		j,
		"match",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitOpenSearchOutput) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitOpenSearchOutput) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitOpenSearchOutput) Pipeline() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pipeline",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitOpenSearchOutput) PluginType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pluginType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitOpenSearchOutput) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitOpenSearchOutput) ReplaceDots() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"replaceDots",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitOpenSearchOutput) SuppressTypeName() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"suppressTypeName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitOpenSearchOutput) TagKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitOpenSearchOutput) TimeKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitOpenSearchOutput) TimeKeyFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeKeyFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitOpenSearchOutput) TimeKeyNanos() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"timeKeyNanos",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitOpenSearchOutput) TraceError() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"traceError",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitOpenSearchOutput) TraceOutput() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"traceOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitOpenSearchOutput) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitOpenSearchOutput) Workers() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"workers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitOpenSearchOutput) WriteOperation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"writeOperation",
		&returns,
	)
	return returns
}


// Creates a new instance of the FluentBitOpenSearchOutput class.
func NewFluentBitOpenSearchOutput(options *FluentBitOpenSearchOutputOptions) FluentBitOpenSearchOutput {
	_init_.Initialize()

	if err := validateNewFluentBitOpenSearchOutputParameters(options); err != nil {
		panic(err)
	}
	j := jsiiProxy_FluentBitOpenSearchOutput{}

	_jsii_.Create(
		"cdk-extensions.k8s_aws.FluentBitOpenSearchOutput",
		[]interface{}{options},
		&j,
	)

	return &j
}

// Creates a new instance of the FluentBitOpenSearchOutput class.
func NewFluentBitOpenSearchOutput_Override(f FluentBitOpenSearchOutput, options *FluentBitOpenSearchOutputOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.k8s_aws.FluentBitOpenSearchOutput",
		[]interface{}{options},
		f,
	)
}

func (f *jsiiProxy_FluentBitOpenSearchOutput) Bind(scope constructs.IConstruct) *ResolvedFluentBitConfiguration {
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

func (f *jsiiProxy_FluentBitOpenSearchOutput) RenderConfigFile(config *map[string]interface{}) *string {
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

