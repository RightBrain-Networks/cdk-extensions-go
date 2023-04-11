//go:build no_runtime_type_checking

package networkmanager

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_GlobalNetwork) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (g *jsiiProxy_GlobalNetwork) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (g *jsiiProxy_GlobalNetwork) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (g *jsiiProxy_GlobalNetwork) validateRegisterTransitGatewayParameters(id *string, options *RegisterTransitGatewayProps) error {
	return nil
}

func validateGlobalNetwork_IsConstructParameters(x interface{}) error {
	return nil
}

func validateGlobalNetwork_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateGlobalNetwork_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewGlobalNetworkParameters(scope constructs.IConstruct, id *string, props *GlobalNetworkProps) error {
	return nil
}

