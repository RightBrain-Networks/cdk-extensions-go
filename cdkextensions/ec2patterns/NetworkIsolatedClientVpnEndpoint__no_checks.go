//go:build no_runtime_type_checking

package ec2patterns

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_NetworkIsolatedClientVpnEndpoint) validateAddAuthorizationRuleParameters(id *string, options *AddAuthorizationRuleOptions) error {
	return nil
}

func (n *jsiiProxy_NetworkIsolatedClientVpnEndpoint) validateAddMultiSubnetRouteParameters(id *string, options *AddMultiSubnetRouteOptions) error {
	return nil
}

func (n *jsiiProxy_NetworkIsolatedClientVpnEndpoint) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (n *jsiiProxy_NetworkIsolatedClientVpnEndpoint) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (n *jsiiProxy_NetworkIsolatedClientVpnEndpoint) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (n *jsiiProxy_NetworkIsolatedClientVpnEndpoint) validateRegisterTransitGatewayParameters(transitGateway ec2.ITransitGateway) error {
	return nil
}

func validateNetworkIsolatedClientVpnEndpoint_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNetworkIsolatedClientVpnEndpoint_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNetworkIsolatedClientVpnEndpoint_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewNetworkIsolatedClientVpnEndpointParameters(scope constructs.IConstruct, id *string, props *NetworkIsolatedClientVpnEndpointProps) error {
	return nil
}

