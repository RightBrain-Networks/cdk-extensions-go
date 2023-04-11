package ec2


// Configuration options for importing a transit gateway peering attachment.
type TransitGatewayPeeringAttachmentImportAttributes struct {
	// The ARN of this Transit Gateway Attachment.
	Arn *string `field:"optional" json:"arn" yaml:"arn"`
	// The ID of this Transit Gateway Attachment.
	AttachmentId *string `field:"optional" json:"attachmentId" yaml:"attachmentId"`
	// The time the transit gateway peering attachment was created.
	CreationTime *string `field:"optional" json:"creationTime" yaml:"creationTime"`
	// The state of the transit gateway peering attachment.
	State *string `field:"optional" json:"state" yaml:"state"`
	// The status of the transit gateway peering attachment.
	Status *string `field:"optional" json:"status" yaml:"status"`
	// The status code for the current status of the attachment.
	StatusCode *string `field:"optional" json:"statusCode" yaml:"statusCode"`
	// The status message for the current status of the attachment.
	StatusMessage *string `field:"optional" json:"statusMessage" yaml:"statusMessage"`
}

