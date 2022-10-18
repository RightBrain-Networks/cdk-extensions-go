package ec2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

// Configuration for the FlowLog class.
type FlowLogProps struct {
	// The AWS account ID this resource belongs to.
	Account *string `field:"optional" json:"account" yaml:"account"`
	// ARN to deduce region and account from.
	//
	// The ARN is parsed and the account and region are taken from the ARN.
	// This should be used for imported resources.
	//
	// Cannot be supplied together with either `account` or `region`.
	EnvironmentFromArn *string `field:"optional" json:"environmentFromArn" yaml:"environmentFromArn"`
	// The value passed in by users to the physical name prop of the resource.
	//
	// - `undefined` implies that a physical name will be allocated by
	//    CloudFormation during deployment.
	// - a concrete value implies a specific physical name
	// - `PhysicalName.GENERATE_IF_NEEDED` is a marker that indicates that a physical will only be generated
	//    by the CDK if it is needed for cross-environment references. Otherwise, it will be allocated by CloudFormation.
	PhysicalName *string `field:"optional" json:"physicalName" yaml:"physicalName"`
	// The AWS region this resource belongs to.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Details for the resource from which flow logs will be captured.
	// See: [FlowLog ResourceType](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-flowlog.html#cfn-ec2-flowlog-resourcetype)
	//
	ResourceType awsec2.FlowLogResourceType `field:"required" json:"resourceType" yaml:"resourceType"`
	// The location where flow logs should be delivered.
	// See: [FlowLog LogDestinationType](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-flowlog.html#cfn-ec2-flowlog-logdestinationtype)
	//
	Destination FlowLogDestination `field:"optional" json:"destination" yaml:"destination"`
	// The fields to include in the flow log record, in the order in which they should appear.
	//
	// For a list of available fields, see {@link FlowLogField}.
	// See: [FlowLog LogFormat](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-flowlog.html#cfn-ec2-flowlog-logformat)
	//
	Format FlowLogFormat `field:"optional" json:"format" yaml:"format"`
	// The maximum interval of time during which a flow of packets is captured and aggregated into a flow log record.
	// See: [FlowLog MaxAggregationInterval](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-flowlog.html#cfn-ec2-flowlog-maxaggregationinterval)
	//
	MaxAggregationInterval FlowLogAggregationInterval `field:"optional" json:"maxAggregationInterval" yaml:"maxAggregationInterval"`
	// The type of traffic to monitor (accepted traffic, rejected traffic, or all traffic).
	// See: [FlowLog TrafficType](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-flowlog.html#cfn-ec2-flowlog-traffictype)
	//
	TrafficType awsec2.FlowLogTrafficType `field:"optional" json:"trafficType" yaml:"trafficType"`
}

