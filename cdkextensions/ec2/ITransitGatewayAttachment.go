package ec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Represents a Transit Gateway Attachment in AWS.
type ITransitGatewayAttachment interface {
	// Adds a route that directs traffic to this transit gateway attachment.
	//
	// Returns: The TransitGatewayRoute that was added.
	AddRoute(id *string, cidr *string, routeTable ITransitGatewayRouteTable) ITransitGatewayRoute
	// The ARN of the transit gateway attachment.
	TransitGatewayAttachmentArn() *string
	// The ID of the transit gateway attachment.
	TransitGatewayAttachmentId() *string
}

// The jsii proxy for ITransitGatewayAttachment
type jsiiProxy_ITransitGatewayAttachment struct {
	_ byte // padding
}

func (i *jsiiProxy_ITransitGatewayAttachment) AddRoute(id *string, cidr *string, routeTable ITransitGatewayRouteTable) ITransitGatewayRoute {
	if err := i.validateAddRouteParameters(id, cidr, routeTable); err != nil {
		panic(err)
	}
	var returns ITransitGatewayRoute

	_jsii_.Invoke(
		i,
		"addRoute",
		[]interface{}{id, cidr, routeTable},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_ITransitGatewayAttachment) TransitGatewayAttachmentArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transitGatewayAttachmentArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITransitGatewayAttachment) TransitGatewayAttachmentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transitGatewayAttachmentId",
		&returns,
	)
	return returns
}

