//go:build no_runtime_type_checking

package ec2

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TransitGatewayAttachment) validateAddRouteParameters(cidr *string, routeTable ITransitGatewayRouteTable) error {
	return nil
}

func (t *jsiiProxy_TransitGatewayAttachment) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (t *jsiiProxy_TransitGatewayAttachment) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (t *jsiiProxy_TransitGatewayAttachment) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateTransitGatewayAttachment_FromTransitGatewayAttachmentIdParameters(scope constructs.IConstruct, id *string, transitGatewayAttachmentId *string) error {
	return nil
}

func validateTransitGatewayAttachment_IsConstructParameters(x interface{}) error {
	return nil
}

func validateTransitGatewayAttachment_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateTransitGatewayAttachment_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewTransitGatewayAttachmentParameters(scope constructs.Construct, id *string, props *TransitGatewayAttachmentProps) error {
	return nil
}

