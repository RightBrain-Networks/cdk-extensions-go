//go:build no_runtime_type_checking

package networkmanager

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TransitGatewayRegistration) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (t *jsiiProxy_TransitGatewayRegistration) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (t *jsiiProxy_TransitGatewayRegistration) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateTransitGatewayRegistration_IsConstructParameters(x interface{}) error {
	return nil
}

func validateTransitGatewayRegistration_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateTransitGatewayRegistration_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewTransitGatewayRegistrationParameters(scope constructs.IConstruct, id *string, props *TransitGatewayRegistrationProps) error {
	return nil
}

