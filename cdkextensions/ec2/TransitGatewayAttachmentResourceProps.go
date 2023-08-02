package ec2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

// Configuration for TransitGatewayAttachmentResource resource.
type TransitGatewayAttachmentResourceProps struct {
	// The AWS account ID this resource belongs to.
	// Default: - the resource is in the same account as the stack it belongs to.
	//
	Account *string `field:"optional" json:"account" yaml:"account"`
	// ARN to deduce region and account from.
	//
	// The ARN is parsed and the account and region are taken from the ARN.
	// This should be used for imported resources.
	//
	// Cannot be supplied together with either `account` or `region`.
	// Default: - take environment from `account`, `region` parameters, or use Stack environment.
	//
	EnvironmentFromArn *string `field:"optional" json:"environmentFromArn" yaml:"environmentFromArn"`
	// The value passed in by users to the physical name prop of the resource.
	//
	// - `undefined` implies that a physical name will be allocated by
	//   CloudFormation during deployment.
	// - a concrete value implies a specific physical name
	// - `PhysicalName.GENERATE_IF_NEEDED` is a marker that indicates that a physical will only be generated
	//   by the CDK if it is needed for cross-environment references. Otherwise, it will be allocated by CloudFormation.
	// Default: - The physical name will be allocated by CloudFormation at deployment time.
	//
	PhysicalName *string `field:"optional" json:"physicalName" yaml:"physicalName"`
	// The AWS region this resource belongs to.
	// Default: - the resource is in the same region as the stack it belongs to.
	//
	Region *string `field:"optional" json:"region" yaml:"region"`
	// The transit gateway for which the attachment should be created.
	// See: [TransitGatewayVpcAttachment TransitGatewayId](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewayvpcattachment.html#cfn-ec2-transitgatewayvpcattachment-transitgatewayid)
	//
	TransitGateway ITransitGateway `field:"required" json:"transitGateway" yaml:"transitGateway"`
	// The VPC where the attachment should be created.
	// See: [TransitGatewayVpcAttachment VpcId](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewayvpcattachment.html#cfn-ec2-transitgatewayvpcattachment-vpcid)
	//
	Vpc awsec2.IVpc `field:"required" json:"vpc" yaml:"vpc"`
	// Enables appliance mode on the attachment.
	//
	// When appliance mode is enabled, all traffic flowing between attachments is
	// forwarded to an appliance in a shared VPC to be inspected and processed.
	// See: [Appliance in a shared services VPC](https://docs.aws.amazon.com/vpc/latest/tgw/transit-gateway-appliance-scenario.html)
	//
	ApplianceModeSupport *bool `field:"optional" json:"applianceModeSupport" yaml:"applianceModeSupport"`
	// Enables DNS support for the attachment.
	//
	// With DNS Support enabled public DNS names that resolve to a connected VPC
	// will be translated to private IP addresses when resolved in a connected VPC.
	// See: [TransitGatewayVpcAttachment DnsSupport](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-transitgatewayvpcattachment-options.html#cfn-ec2-transitgatewayvpcattachment-options-dnssupport)
	//
	DnsSupport *bool `field:"optional" json:"dnsSupport" yaml:"dnsSupport"`
	// Enables DNS support for the attachment.
	//
	// With DNS Support enabled public DNS names that resolve to a connected VPC
	// will be translated to private IP addresses when resolved in a connected VPC.
	// See: [IPv6 connectivity with TransitGateway](https://docs.aws.amazon.com/whitepapers/latest/ipv6-on-aws/amazon-vpc-connectivity-options-for-ipv6.html#ipv6-connectivity-with-transit-gateway)
	//
	Ipv6Support *bool `field:"optional" json:"ipv6Support" yaml:"ipv6Support"`
	// The name of the Transit Gateway Attachment.
	//
	// Used to tag the attachment with a name that will be displayed in the AWS
	// EC2 console.
	// See: [TransitGatewayVpcAttachment Tags](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewayvpcattachment.html#cfn-ec2-transitgatewayvpcattachment-tags)
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The subnets where the attachment should be created.
	//
	// Can select up to one subnet per Availability Zone.
	// See: [TransitGatewayVpcAttachment SubnetIds](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewayvpcattachment.html#cfn-ec2-transitgatewayvpcattachment-subnetids)
	//
	Subnets *awsec2.SubnetSelection `field:"optional" json:"subnets" yaml:"subnets"`
}

