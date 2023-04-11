package ec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Represents a customer gateway in AWS.
type ICustomerGateway interface {
	// The BGP ASN of the customer gateway.
	CustomerGatewayAsn() *float64
	// The ID of the customer gateway.
	CustomerGatewayId() *string
	// The IP address of the customer gateway.
	CustomerGatewayIp() *string
}

// The jsii proxy for ICustomerGateway
type jsiiProxy_ICustomerGateway struct {
	_ byte // padding
}

func (j *jsiiProxy_ICustomerGateway) CustomerGatewayAsn() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"customerGatewayAsn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICustomerGateway) CustomerGatewayId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customerGatewayId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICustomerGateway) CustomerGatewayIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customerGatewayIp",
		&returns,
	)
	return returns
}

