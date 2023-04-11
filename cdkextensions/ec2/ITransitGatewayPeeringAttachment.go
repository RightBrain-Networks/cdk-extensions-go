package ec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Represents a transit gateway route table in AWS.
type ITransitGatewayPeeringAttachment interface {
	ITransitGatewayAttachment
	// The time the transit gateway peering attachment was created.
	TransitGatewayAttachmentCreationTime() *string
	// The state of the transit gateway peering attachment.
	TransitGatewayAttachmentState() *string
	// The status of the transit gateway peering attachment.
	TransitGatewayAttachmentStatus() *string
	// The status code for the current status of the attachment.
	TransitGatewayAttachmentStatusCode() *string
	// The status message for the current status of the attachment.
	TransitGatewayAttachmentStatusMessage() *string
}

// The jsii proxy for ITransitGatewayPeeringAttachment
type jsiiProxy_ITransitGatewayPeeringAttachment struct {
	jsiiProxy_ITransitGatewayAttachment
}

func (j *jsiiProxy_ITransitGatewayPeeringAttachment) TransitGatewayAttachmentCreationTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transitGatewayAttachmentCreationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITransitGatewayPeeringAttachment) TransitGatewayAttachmentState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transitGatewayAttachmentState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITransitGatewayPeeringAttachment) TransitGatewayAttachmentStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transitGatewayAttachmentStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITransitGatewayPeeringAttachment) TransitGatewayAttachmentStatusCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transitGatewayAttachmentStatusCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITransitGatewayPeeringAttachment) TransitGatewayAttachmentStatusMessage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transitGatewayAttachmentStatusMessage",
		&returns,
	)
	return returns
}

