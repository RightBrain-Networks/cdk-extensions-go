package k8saws

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Options for configuring the Elasticsearch Fluent Bit output plugin.
// See: [Opensearch Plugin Documention](https://docs.fluentbit.io/manual/pipeline/outputs/elasticsearch)
//
type FluentBitElasticsearchOutputOptions struct {
	// The pattern to match for records that this output should apply to.
	Match FluentBitMatch `field:"optional" json:"match" yaml:"match"`
	// IP address or hostname of the target Elasticsearch instance.
	Host *string `field:"required" json:"host" yaml:"host"`
	// Enable AWS Sigv4 Authentication for Amazon Elasticsearch Service.
	AwsAuth *bool `field:"optional" json:"awsAuth" yaml:"awsAuth"`
	// External ID for the AWS IAM Role specified with `awsRole`.
	AwsExternalId *string `field:"optional" json:"awsExternalId" yaml:"awsExternalId"`
	// Specify the AWS region for Elasticsearch Service.
	AwsRegion *string `field:"optional" json:"awsRegion" yaml:"awsRegion"`
	// AWS IAM Role to assume to put records to your Amazon cluster.
	AwsRole awsiam.IRole `field:"optional" json:"awsRole" yaml:"awsRole"`
	// Specify the custom sts endpoint to be used with STS API for Amazon Elasticsearch Service.
	AwsStsEndpoint *string `field:"optional" json:"awsStsEndpoint" yaml:"awsStsEndpoint"`
	// Specify the buffer size used to read the response from the Elasticsearch HTTP service.
	//
	// This option is useful for debugging purposes where is
	// required to read full responses, note that response size grows depending
	// of the number of records inserted.
	BufferSize ElasticsearchOutputBufferSize `field:"optional" json:"bufferSize" yaml:"bufferSize"`
	// Specify the credentials to use to connect to Elastic's Elasticsearch Service running on Elastic Cloud.
	CloudAuth *string `field:"optional" json:"cloudAuth" yaml:"cloudAuth"`
	// If you are using Elastic's Elasticsearch Service you can specify the cloud_id of the cluster running.
	CloudId *string `field:"optional" json:"cloudId" yaml:"cloudId"`
	// Set payload compression mechanism.
	Compress ElasticsearchCompressionFormat `field:"optional" json:"compress" yaml:"compress"`
	// Use current time for index generation instead of message record.
	CurrentTimeIndex *bool `field:"optional" json:"currentTimeIndex" yaml:"currentTimeIndex"`
	// When enabled, generate `_id` for outgoing records.
	//
	// This prevents duplicate
	// records when retrying.
	GenerateId *bool `field:"optional" json:"generateId" yaml:"generateId"`
	// Password for user defined in `httpUser`.
	HttpPasswd *string `field:"optional" json:"httpPasswd" yaml:"httpPasswd"`
	// Optional username credential for access.
	HttpUser *string `field:"optional" json:"httpUser" yaml:"httpUser"`
	// If set, `_id` will be the value of the key from incoming record and `generateId` option is ignored.
	IdKey *string `field:"optional" json:"idKey" yaml:"idKey"`
	// When enabled, it append the Tag name to the record.
	IncludeTagKey *bool `field:"optional" json:"includeTagKey" yaml:"includeTagKey"`
	// Index name.
	Index *string `field:"optional" json:"index" yaml:"index"`
	// Time format (based on strftime) to generate the second part of the Index name.
	// See: [strftime](http://man7.org/linux/man-pages/man3/strftime.3.html)
	//
	LogstashDateFormat *string `field:"optional" json:"logstashDateFormat" yaml:"logstashDateFormat"`
	// Enable Logstash format compatibility.
	LogstashFormat *bool `field:"optional" json:"logstashFormat" yaml:"logstashFormat"`
	// When `logstashFormat` is enabled, the Index name is composed using a prefix and the date, e.g: If `logstashPrefix` is equals to 'mydata' your index will become 'mydata-YYYY.MM.DD'.
	//
	// The last string appended belongs to the date when the data is being
	// generated.
	LogstashPrefix *string `field:"optional" json:"logstashPrefix" yaml:"logstashPrefix"`
	// When included: the value in the record that belongs to the key will be looked up and over-write the `logstashPrefix` for index generation.
	//
	// If
	// the key/value is not found in the record then the `logstashPrefix` option
	// will act as a fallback.
	//
	// Nested keys are not supported (if desired, you can use the nest filter
	// plugin to remove nesting).
	LogstashPrefixKey *string `field:"optional" json:"logstashPrefixKey" yaml:"logstashPrefixKey"`
	// Elasticsearch accepts new data on HTTP query path "/_bulk".
	//
	// But it is
	// also possible to serve Elasticsearch behind a reverse proxy on a
	// subpath. This option defines such path on the fluent-bit side. It
	// simply adds a path prefix in the indexing HTTP POST URI..
	Path *string `field:"optional" json:"path" yaml:"path"`
	// Elasticsearch allows to setup filters called pipelines.
	//
	// This option
	// allows to define which pipeline the database should use.
	Pipeline *string `field:"optional" json:"pipeline" yaml:"pipeline"`
	// TCP port of the target Elasticsearch instance.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// When enabled, replace field name dots with underscore.
	ReplaceDots *bool `field:"optional" json:"replaceDots" yaml:"replaceDots"`
	// When enabled, mapping types is removed and `type` option is ignored.
	SuppressTypeName *bool `field:"optional" json:"suppressTypeName" yaml:"suppressTypeName"`
	// When `includeTagKey` is enabled, this property defines the key name for the tag.
	TagKey *string `field:"optional" json:"tagKey" yaml:"tagKey"`
	// When `logstashFormat` is enabled, each record will get a new timestamp field.
	//
	// The`timeKey` property defines the name of that field.
	TimeKey *string `field:"optional" json:"timeKey" yaml:"timeKey"`
	// When `logstashFormat` is enabled, this property defines the format of the timestamp.
	// See: [strftime](http://man7.org/linux/man-pages/man3/strftime.3.html)
	//
	TimeKeyFormat *string `field:"optional" json:"timeKeyFormat" yaml:"timeKeyFormat"`
	// When `logstashFormat` is enabled, enabling this property sends nanosecond precision timestamps.
	TimeKeyNanos *bool `field:"optional" json:"timeKeyNanos" yaml:"timeKeyNanos"`
	// When enabled print the Elasticsearch API calls to stdout when Elasticsearch returns an error (for diag only).
	TraceError *bool `field:"optional" json:"traceError" yaml:"traceError"`
	// When enabled print the Elasticsearch API calls to stdout (for diag only).
	TraceOutput *bool `field:"optional" json:"traceOutput" yaml:"traceOutput"`
	// Type name.
	Type *string `field:"optional" json:"type" yaml:"type"`
	// Enables dedicated thread(s) for this output.
	Workers *float64 `field:"optional" json:"workers" yaml:"workers"`
	// Operation to use to write in bulk requests.
	WriteOperation *string `field:"optional" json:"writeOperation" yaml:"writeOperation"`
}

