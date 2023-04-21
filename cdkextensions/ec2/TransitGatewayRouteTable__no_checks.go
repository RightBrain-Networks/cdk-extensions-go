//go:build no_runtime_type_checking

package ec2

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TransitGatewayRouteTable) validateAddRouteParameters(id *string, options *TransitGatewayRouteOptions) error {
	return nil
}

func (t *jsiiProxy_TransitGatewayRouteTable) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (t *jsiiProxy_TransitGatewayRouteTable) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (t *jsiiProxy_TransitGatewayRouteTable) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateTransitGatewayRouteTable_FromTransitGatewayRouteTableIdParameters(scope constructs.IConstruct, id *string, transitGatewayRouteTableId *string) error {
	return nil
}

func validateTransitGatewayRouteTable_IsConstructParameters(x interface{}) error {
	return nil
}

func validateTransitGatewayRouteTable_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateTransitGatewayRouteTable_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewTransitGatewayRouteTableParameters(scope constructs.Construct, id *string, props *TransitGatewayRouteTableProps) error {
	return nil
}

