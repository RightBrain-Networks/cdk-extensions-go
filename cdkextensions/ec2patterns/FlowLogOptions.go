package ec2patterns

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/vibe-io/cdk-extensions-go/cdkextensions/ec2"
)

type FlowLogOptions struct {
	// Specifies the type of destination to which the flow log data is to be published.
	//
	// Flow log data can be published to CloudWatch Logs or Amazon S3.
	Destination awsec2.FlowLogDestination `field:"optional" json:"destination" yaml:"destination"`
	// The fields to include in the flow log record, in the order in which they should appear.
	//
	// If multiple fields are specified, they will be separated by spaces. For full control over the literal log format
	// string, pass a single field constructed with `LogFormat.custom()`.
	//
	// See https://docs.aws.amazon.com/vpc/latest/userguide/flow-logs.html#flow-log-records
	LogFormat *[]awsec2.LogFormat `field:"optional" json:"logFormat" yaml:"logFormat"`
	// The maximum interval of time during which a flow of packets is captured and aggregated into a flow log record.
	MaxAggregationInterval awsec2.FlowLogMaxAggregationInterval `field:"optional" json:"maxAggregationInterval" yaml:"maxAggregationInterval"`
	// The type of traffic to log.
	//
	// You can log traffic that the resource accepts or rejects, or all traffic.
	TrafficType awsec2.FlowLogTrafficType `field:"optional" json:"trafficType" yaml:"trafficType"`
	LogFormatDefinition ec2.FlowLogFormat `field:"optional" json:"logFormatDefinition" yaml:"logFormatDefinition"`
}

