//go:build no_runtime_type_checking

package ec2

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TransitGatewayPeeringAttachment) validateAddRouteParameters(id *string, cidr *string, routeTable ITransitGatewayRouteTable) error {
	return nil
}

func (t *jsiiProxy_TransitGatewayPeeringAttachment) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (t *jsiiProxy_TransitGatewayPeeringAttachment) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (t *jsiiProxy_TransitGatewayPeeringAttachment) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateTransitGatewayPeeringAttachment_FromTransitGatewayPeeringAttachmentArnParameters(scope constructs.IConstruct, id *string, arn *string) error {
	return nil
}

func validateTransitGatewayPeeringAttachment_FromTransitGatewayPeeringAttachmentAttributesParameters(scope constructs.IConstruct, id *string, attrs *TransitGatewayPeeringAttachmentImportAttributes) error {
	return nil
}

func validateTransitGatewayPeeringAttachment_FromTransitGatewayPeeringAttachmentIdParameters(scope constructs.IConstruct, id *string, attachmentId *string) error {
	return nil
}

func validateTransitGatewayPeeringAttachment_IsConstructParameters(x interface{}) error {
	return nil
}

func validateTransitGatewayPeeringAttachment_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateTransitGatewayPeeringAttachment_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewTransitGatewayPeeringAttachmentParameters(scope constructs.Construct, id *string, props *TransitGatewayPeeringAttachmentProps) error {
	return nil
}

