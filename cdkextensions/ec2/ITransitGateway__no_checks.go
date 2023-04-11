//go:build no_runtime_type_checking

package ec2

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_ITransitGateway) validateAddRouteTableParameters(options *TransitGatewayRouteTableOptions) error {
	return nil
}

func (i *jsiiProxy_ITransitGateway) validateAddVpnParameters(id *string, options *VpnAttachmentOptions) error {
	return nil
}

func (i *jsiiProxy_ITransitGateway) validateAttachVpcParameters(vpc awsec2.IVpc, options *VpcAttachmentOptions) error {
	return nil
}

