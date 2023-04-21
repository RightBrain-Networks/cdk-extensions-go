//go:build no_runtime_type_checking

package ec2

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TransitGatewayAttachmentBase) validateAddRouteParameters(id *string, cidr *string, routeTable ITransitGatewayRouteTable) error {
	return nil
}

func (t *jsiiProxy_TransitGatewayAttachmentBase) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (t *jsiiProxy_TransitGatewayAttachmentBase) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (t *jsiiProxy_TransitGatewayAttachmentBase) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateTransitGatewayAttachmentBase_IsConstructParameters(x interface{}) error {
	return nil
}

func validateTransitGatewayAttachmentBase_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateTransitGatewayAttachmentBase_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewTransitGatewayAttachmentBaseParameters(scope constructs.Construct, id *string, props *awscdk.ResourceProps) error {
	return nil
}

