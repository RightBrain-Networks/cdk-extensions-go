package k8saws

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
)

// Options for configuring the CloudWatch Logs Fluent Bit output plugin.
// See: [CloudWatch Logs Plugin Documention](https://docs.fluentbit.io/manual/pipeline/outputs/cloudwatch)
//
type FluentBitCloudWatchLogsOutputOptions struct {
	// The pattern to match for records that this output should apply to.
	Match FluentBitMatch `field:"optional" json:"match" yaml:"match"`
	// Automatically create the log group.
	// Default: false.
	//
	AutoCreateGroup *bool `field:"optional" json:"autoCreateGroup" yaml:"autoCreateGroup"`
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
	// Specify a custom endpoint for the CloudWatch Logs API.
	Endpoint *string `field:"optional" json:"endpoint" yaml:"endpoint"`
	// An optional parameter that can be used to tell CloudWatch the format of the data.
	//
	// A value of json/emf enables CloudWatch to extract custom
	// metrics embedded in a JSON payload.
	// See: [Embedded Metric Format](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Embedded_Metric_Format_Specification.html)
	//
	LogFormat *string `field:"optional" json:"logFormat" yaml:"logFormat"`
	// The CloudWatch Log Group configuration for output records.
	LogGroup FluentBitLogGroupOutput `field:"optional" json:"logGroup" yaml:"logGroup"`
	// Template for Log Group name using Fluent Bit record_accessor syntax.
	//
	// This field is optional and if configured it overrides the configured Log
	// Group.
	//
	// If the template translation fails, an error is logged and the provided
	// Log Group (which is still required) is used instead.
	// See: [Fluent Bit record accessor snytax](https://docs.fluentbit.io/manual/administration/configuring-fluent-bit/classic-mode/record-accessor)
	//
	LogGroupTemplate *string `field:"optional" json:"logGroupTemplate" yaml:"logGroupTemplate"`
	// By default, the whole log record will be sent to CloudWatch.
	//
	// If you
	// specify a key name with this option, then only the value of that key
	// will be sent to CloudWatch.
	LogKey *string `field:"optional" json:"logKey" yaml:"logKey"`
	// If set to a number greater than zero, and newly create log group's retention policy is set to this many days.
	LogRetention awslogs.RetentionDays `field:"optional" json:"logRetention" yaml:"logRetention"`
	// The CloudWatch LogStream configuration for outbound records.
	LogStream FluentBitLogStreamOutput `field:"optional" json:"logStream" yaml:"logStream"`
	// Template for Log Stream name using Fluent Bit record accessor syntax.
	//
	// This field is optional and if configured it overrides the other log
	// stream options. If the template translation fails, an error is logged
	// and the logStream or logStreamPrefix are used instead (and thus one of
	// those fields is still required to be configured).
	// See: [Fluent Bit record accessor snytax](https://docs.fluentbit.io/manual/administration/configuring-fluent-bit/classic-mode/record-accessor)
	//
	LogStreamTemplate *string `field:"optional" json:"logStreamTemplate" yaml:"logStreamTemplate"`
	// A list of lists containing the dimension keys that will be applied to all metrics.
	//
	// The values within a dimension set MUST also be members on
	// the root-node.
	// See: [Dimensions](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/cloudwatch_concepts.html#Dimension)
	//
	MetricDimensions *[]*string `field:"optional" json:"metricDimensions" yaml:"metricDimensions"`
	// An optional string representing the CloudWatch namespace for the metrics.
	// See: [Metric Tutorial](https://docs.fluentbit.io/manual/pipeline/outputs/cloudwatch#metrics-tutorial)
	//
	MetricNamespace *string `field:"optional" json:"metricNamespace" yaml:"metricNamespace"`
	// The AWS region.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// ARN of an IAM role to assume (for cross account access).
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// Specify a custom STS endpoint for the AWS STS API.
	StsEndpoint *string `field:"optional" json:"stsEndpoint" yaml:"stsEndpoint"`
}

