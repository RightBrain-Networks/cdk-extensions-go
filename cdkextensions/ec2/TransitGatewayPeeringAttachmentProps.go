package ec2


// Configuration for TransitGatewayPeeringAttachment resource.
type TransitGatewayPeeringAttachmentProps struct {
	// The name of the transit gateway peering attachment.
	//
	// Used to tag the attachment with a name that will be displayed in the AWS
	// EC2 console.
	// See: [TransitGatewayPeeringAttachment Tags](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewaypeeringattachment.html#cfn-ec2-transitgatewaypeeringattachment-tags)
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The account that contains the transit gateway being peered with.
	// See: [TransitGatewayPeeringAttachment PeerAccountId](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewaypeeringattachment.html#cfn-ec2-transitgatewaypeeringattachment-peeraccountid)
	//
	PeerAccountId *string `field:"optional" json:"peerAccountId" yaml:"peerAccountId"`
	// The region that contains the transit gateway being peered with.
	// See: [TransitGatewayPeeringAttachment PeerRegion](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewaypeeringattachment.html#cfn-ec2-transitgatewaypeeringattachment-peerregion)
	//
	PeerRegion *string `field:"optional" json:"peerRegion" yaml:"peerRegion"`
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
	//   CloudFormation during deployment.
	// - a concrete value implies a specific physical name
	// - `PhysicalName.GENERATE_IF_NEEDED` is a marker that indicates that a physical will only be generated
	//   by the CDK if it is needed for cross-environment references. Otherwise, it will be allocated by CloudFormation.
	PhysicalName *string `field:"optional" json:"physicalName" yaml:"physicalName"`
	// The AWS region this resource belongs to.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// The local side of the transit gateway peering connection.
	// See: [TransitGatewayPeeringAttachment TransitGatewayId](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewaypeeringattachment.html#cfn-ec2-transitgatewaypeeringattachment-transitgatewayid)
	//
	LocalTransitGateway ITransitGateway `field:"required" json:"localTransitGateway" yaml:"localTransitGateway"`
	// The remote transit gateway being peered with.
	// See: [TransitGatewayPeeringAttachment PeerTransitGatewayId](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewaypeeringattachment.html#cfn-ec2-transitgatewaypeeringattachment-peertransitgatewayid)
	//
	PeerTransitGateway ITransitGateway `field:"required" json:"peerTransitGateway" yaml:"peerTransitGateway"`
}

