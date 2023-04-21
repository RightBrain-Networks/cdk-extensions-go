package ec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/vibe-io/cdk-extensions-go/cdkextensions/ec2/internal"
)

// Represents a transit gateway in AWS.
type ITransitGateway interface {
	constructs.IConstruct
	AddRouteTable(options *TransitGatewayRouteTableOptions) TransitGatewayRouteTable
	AddVpn(id *string, options *VpnAttachmentOptions) VpnConnection
	AttachVpc(vpc awsec2.IVpc, options *VpcAttachmentOptions) TransitGatewayAttachment
	DefaultRouteTable() ITransitGatewayRouteTable
	TransitGatewayArn() *string
	TransitGatewayId() *string
}

// The jsii proxy for ITransitGateway
type jsiiProxy_ITransitGateway struct {
	internal.Type__constructsIConstruct
}

func (i *jsiiProxy_ITransitGateway) AddRouteTable(options *TransitGatewayRouteTableOptions) TransitGatewayRouteTable {
	if err := i.validateAddRouteTableParameters(options); err != nil {
		panic(err)
	}
	var returns TransitGatewayRouteTable

	_jsii_.Invoke(
		i,
		"addRouteTable",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ITransitGateway) AddVpn(id *string, options *VpnAttachmentOptions) VpnConnection {
	if err := i.validateAddVpnParameters(id, options); err != nil {
		panic(err)
	}
	var returns VpnConnection

	_jsii_.Invoke(
		i,
		"addVpn",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ITransitGateway) AttachVpc(vpc awsec2.IVpc, options *VpcAttachmentOptions) TransitGatewayAttachment {
	if err := i.validateAttachVpcParameters(vpc, options); err != nil {
		panic(err)
	}
	var returns TransitGatewayAttachment

	_jsii_.Invoke(
		i,
		"attachVpc",
		[]interface{}{vpc, options},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_ITransitGateway) DefaultRouteTable() ITransitGatewayRouteTable {
	var returns ITransitGatewayRouteTable
	_jsii_.Get(
		j,
		"defaultRouteTable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITransitGateway) TransitGatewayArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transitGatewayArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITransitGateway) TransitGatewayId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transitGatewayId",
		&returns,
	)
	return returns
}

