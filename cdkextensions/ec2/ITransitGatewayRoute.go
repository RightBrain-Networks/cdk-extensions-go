package ec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Represents a Transit Gateway Route in AWS.
type ITransitGatewayRoute interface {
	// The ID of the Transit Gateway Route.
	TransitGatewayRouteId() *string
}

// The jsii proxy for ITransitGatewayRoute
type jsiiProxy_ITransitGatewayRoute struct {
	_ byte // padding
}

func (j *jsiiProxy_ITransitGatewayRoute) TransitGatewayRouteId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transitGatewayRouteId",
		&returns,
	)
	return returns
}

