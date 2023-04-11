package ec2


// Optional configuration for TransitGatewayPeeringAttachment resource.
type TransitGatewayPeeringAttachmentOptions struct {
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
}

