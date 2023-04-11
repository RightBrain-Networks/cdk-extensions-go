//go:build no_runtime_type_checking

package ec2

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TransitGatewayAttachmentResource) validateAddRouteParameters(cidr *string, routeTable ITransitGatewayRouteTable) error {
	return nil
}

func (t *jsiiProxy_TransitGatewayAttachmentResource) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (t *jsiiProxy_TransitGatewayAttachmentResource) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (t *jsiiProxy_TransitGatewayAttachmentResource) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateTransitGatewayAttachmentResource_IsConstructParameters(x interface{}) error {
	return nil
}

func validateTransitGatewayAttachmentResource_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateTransitGatewayAttachmentResource_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewTransitGatewayAttachmentResourceParameters(scope constructs.Construct, id *string, props *TransitGatewayAttachmentResourceProps) error {
	return nil
}

