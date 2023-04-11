//go:build no_runtime_type_checking

package ec2

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TransitGateway) validateAddCidrBlockParameters(cidr *string) error {
	return nil
}

func (t *jsiiProxy_TransitGateway) validateAddRouteTableParameters(options *TransitGatewayRouteTableOptions) error {
	return nil
}

func (t *jsiiProxy_TransitGateway) validateAddVpnParameters(id *string, options *VpnAttachmentOptions) error {
	return nil
}

func (t *jsiiProxy_TransitGateway) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (t *jsiiProxy_TransitGateway) validateAttachPeerParameters(peer ITransitGateway, options *TransitGatewayPeeringAttachmentOptions) error {
	return nil
}

func (t *jsiiProxy_TransitGateway) validateAttachVpcParameters(vpc awsec2.IVpc, options *VpcAttachmentOptions) error {
	return nil
}

func (t *jsiiProxy_TransitGateway) validateEnableSharingParameters(options *SharingOptions) error {
	return nil
}

func (t *jsiiProxy_TransitGateway) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (t *jsiiProxy_TransitGateway) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateTransitGateway_FromTransitGatewayIdParameters(scope constructs.IConstruct, id *string, transitGatewayId *string) error {
	return nil
}

func validateTransitGateway_IsConstructParameters(x interface{}) error {
	return nil
}

func validateTransitGateway_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateTransitGateway_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewTransitGatewayParameters(scope constructs.Construct, id *string, props *TransitGatewayProps) error {
	return nil
}

