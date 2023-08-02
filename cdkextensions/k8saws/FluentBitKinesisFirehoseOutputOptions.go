package k8saws

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/vibe-io/cdk-extensions-go/cdkextensions/kinesisfirehose"
)

// Options for configuring the Kinesis Firehose Fluent Bit output plugin.
// See: [Kinesis Firehose Plugin Documention](https://docs.fluentbit.io/manual/pipeline/outputs/firehose)
//
type FluentBitKinesisFirehoseOutputOptions struct {
	// The pattern to match for records that this output should apply to.
	Match FluentBitMatch `field:"optional" json:"match" yaml:"match"`
	// Immediately retry failed requests to AWS services once.
	//
	// This option does
	// not affect the normal Fluent Bit retry mechanism with backoff. Instead,
	// it enables an immediate retry with no delay for networking errors, which
	// may help improve throughput when there are transient/random networking
	// issues.
	// Default: true.
	//
	AutoRetryRequests *bool `field:"optional" json:"autoRetryRequests" yaml:"autoRetryRequests"`
	// Compression type for Firehose records.
	//
	// Each log record is individually
	// compressed and sent to Firehose.
	Compression KinesisFirehoseCompressionFormat `field:"optional" json:"compression" yaml:"compression"`
	// The Kinesis Firehose Delivery stream that you want log records sent to.
	DeliveryStream kinesisfirehose.IDeliveryStream `field:"optional" json:"deliveryStream" yaml:"deliveryStream"`
	// Specify a custom endpoint for the Firehose API.
	Endpoint *string `field:"optional" json:"endpoint" yaml:"endpoint"`
	// By default, the whole log record will be sent to Firehose.
	//
	// If you
	// specify a key name with this option, then only the value of that key
	// will be sent to Firehose.
	LogKey *string `field:"optional" json:"logKey" yaml:"logKey"`
	// The AWS region.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// ARN of an IAM role to assume (for cross account access).
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// Specify a custom STS endpoint for the AWS STS API.
	StsEndpoint *string `field:"optional" json:"stsEndpoint" yaml:"stsEndpoint"`
	// Add the timestamp to the record under this key.
	TimeKey *string `field:"optional" json:"timeKey" yaml:"timeKey"`
	// A strftime compliant format string for the timestamp.
	// Default: '%Y-%m-%dT%H:%M:%S'.
	//
	TimeKeyFormat *string `field:"optional" json:"timeKeyFormat" yaml:"timeKeyFormat"`
}

