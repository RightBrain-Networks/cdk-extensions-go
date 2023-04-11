package ec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Represents a transit gateway route table in AWS.
type ITransitGatewayRouteTable interface {
	// Adds a route to this transit gateway route table.
	AddRoute(options *TransitGatewayRouteOptions) TransitGatewayRoute
	// The ARN of the transit gateway route table.
	TransitGatewayRouteTableArn() *string
	// The ID of the transit gateway route table.
	TransitGatewayRouteTableId() *string
}

// The jsii proxy for ITransitGatewayRouteTable
type jsiiProxy_ITransitGatewayRouteTable struct {
	_ byte // padding
}

func (i *jsiiProxy_ITransitGatewayRouteTable) AddRoute(options *TransitGatewayRouteOptions) TransitGatewayRoute {
	if err := i.validateAddRouteParameters(options); err != nil {
		panic(err)
	}
	var returns TransitGatewayRoute

	_jsii_.Invoke(
		i,
		"addRoute",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_ITransitGatewayRouteTable) TransitGatewayRouteTableArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transitGatewayRouteTableArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITransitGatewayRouteTable) TransitGatewayRouteTableId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transitGatewayRouteTableId",
		&returns,
	)
	return returns
}

