package k8saws

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/constructs-go/constructs/v10"
)

type FluentBitElasticsearchOutput interface {
	FluentBitOutputPluginBase
	// Enable AWS Sigv4 Authentication for Amazon Elasticsearch Service.
	AwsAuth() *bool
	// External ID for the AWS IAM Role specified with `awsRole`.
	AwsExternalId() *string
	// Specify the AWS region for Elasticsearch Service.
	AwsRegion() *string
	// AWS IAM Role to assume to put records to your Amazon cluster.
	AwsRole() awsiam.IRole
	// Specify the custom sts endpoint to be used with STS API for Amazon Elasticsearch Service.
	AwsStsEndpoint() *string
	// Specify the buffer size used to read the response from the Elasticsearch HTTP service.
	//
	// This option is useful for debugging purposes where is
	// required to read full responses, note that response size grows depending
	// of the number of records inserted.
	BufferSize() ElasticsearchOutputBufferSize
	// Specify the credentials to use to connect to Elastic's Elasticsearch Service running on Elastic Cloud.
	CloudAuth() *string
	// If you are using Elastic's Elasticsearch Service you can specify the cloud_id of the cluster running.
	CloudId() *string
	// Set payload compression mechanism.
	Compress() ElasticsearchCompressionFormat
	// Use current time for index generation instead of message record.
	CurrentTimeIndex() *bool
	// When enabled, generate `_id` for outgoing records.
	//
	// This prevents duplicate
	// records when retrying.
	GenerateId() *bool
	// IP address or hostname of the target Elasticsearch instance.
	Host() *string
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
	// Elasticsearch accepts new data on HTTP query path "/_bulk".
	//
	// But it is
	// also possible to serve Elasticsearch behind a reverse proxy on a
	// subpath. This option defines such path on the fluent-bit side. It
	// simply adds a path prefix in the indexing HTTP POST URI.
	Path() *string
	// Elasticsearch allows to setup filters called pipelines.
	//
	// This option
	// allows to define which pipeline the database should use.
	Pipeline() *string
	// The type of fluent bit plugin.
	PluginType() *string
	// TCP port of the target Elasticsearch instance.
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
	// When enabled print the Elasticsearch API calls to stdout when Elasticsearch returns an error (for diag only).
	TraceError() *bool
	// When enabled print the Elasticsearch API calls to stdout (for diag only).
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
	Bind(_scope constructs.IConstruct) *ResolvedFluentBitConfiguration
	// Renders a Fluent Bit configuration file for the plugin.
	//
	// Returns: A rendered plugin configuration file.
	RenderConfigFile(config *map[string]interface{}) *string
}

// The jsii proxy struct for FluentBitElasticsearchOutput
type jsiiProxy_FluentBitElasticsearchOutput struct {
	jsiiProxy_FluentBitOutputPluginBase
}

func (j *jsiiProxy_FluentBitElasticsearchOutput) AwsAuth() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"awsAuth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitElasticsearchOutput) AwsExternalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsExternalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitElasticsearchOutput) AwsRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitElasticsearchOutput) AwsRole() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"awsRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitElasticsearchOutput) AwsStsEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsStsEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitElasticsearchOutput) BufferSize() ElasticsearchOutputBufferSize {
	var returns ElasticsearchOutputBufferSize
	_jsii_.Get(
		j,
		"bufferSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitElasticsearchOutput) CloudAuth() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudAuth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitElasticsearchOutput) CloudId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitElasticsearchOutput) Compress() ElasticsearchCompressionFormat {
	var returns ElasticsearchCompressionFormat
	_jsii_.Get(
		j,
		"compress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitElasticsearchOutput) CurrentTimeIndex() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"currentTimeIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitElasticsearchOutput) GenerateId() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"generateId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitElasticsearchOutput) Host() *string {
	var returns *string
	_jsii_.Get(
		j,
		"host",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitElasticsearchOutput) HttpPasswd() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpPasswd",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitElasticsearchOutput) HttpUser() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpUser",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitElasticsearchOutput) IdKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitElasticsearchOutput) IncludeTagKey() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"includeTagKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitElasticsearchOutput) Index() *string {
	var returns *string
	_jsii_.Get(
		j,
		"index",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitElasticsearchOutput) LogstashDateFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logstashDateFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitElasticsearchOutput) LogstashFormat() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"logstashFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitElasticsearchOutput) LogstashPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logstashPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitElasticsearchOutput) LogstashPrefixKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logstashPrefixKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitElasticsearchOutput) Match() FluentBitMatch {
	var returns FluentBitMatch
	_jsii_.Get(
		j,
		"match",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitElasticsearchOutput) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitElasticsearchOutput) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitElasticsearchOutput) Pipeline() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pipeline",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitElasticsearchOutput) PluginType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pluginType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitElasticsearchOutput) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitElasticsearchOutput) ReplaceDots() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"replaceDots",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitElasticsearchOutput) SuppressTypeName() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"suppressTypeName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitElasticsearchOutput) TagKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitElasticsearchOutput) TimeKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitElasticsearchOutput) TimeKeyFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeKeyFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitElasticsearchOutput) TimeKeyNanos() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"timeKeyNanos",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitElasticsearchOutput) TraceError() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"traceError",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitElasticsearchOutput) TraceOutput() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"traceOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitElasticsearchOutput) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitElasticsearchOutput) Workers() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"workers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitElasticsearchOutput) WriteOperation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"writeOperation",
		&returns,
	)
	return returns
}


// Creates a new instance of the FluentBitKinesisFirehoseOutput class.
func NewFluentBitElasticsearchOutput(options *FluentBitElasticsearchOutputOptions) FluentBitElasticsearchOutput {
	_init_.Initialize()

	if err := validateNewFluentBitElasticsearchOutputParameters(options); err != nil {
		panic(err)
	}
	j := jsiiProxy_FluentBitElasticsearchOutput{}

	_jsii_.Create(
		"cdk-extensions.k8s_aws.FluentBitElasticsearchOutput",
		[]interface{}{options},
		&j,
	)

	return &j
}

// Creates a new instance of the FluentBitKinesisFirehoseOutput class.
func NewFluentBitElasticsearchOutput_Override(f FluentBitElasticsearchOutput, options *FluentBitElasticsearchOutputOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.k8s_aws.FluentBitElasticsearchOutput",
		[]interface{}{options},
		f,
	)
}

func (f *jsiiProxy_FluentBitElasticsearchOutput) Bind(_scope constructs.IConstruct) *ResolvedFluentBitConfiguration {
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

func (f *jsiiProxy_FluentBitElasticsearchOutput) RenderConfigFile(config *map[string]interface{}) *string {
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

