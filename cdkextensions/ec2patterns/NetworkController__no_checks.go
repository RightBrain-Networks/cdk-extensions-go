//go:build no_runtime_type_checking

package ec2patterns

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_NetworkController) validateAddHubParameters(scope constructs.IConstruct, id *string, options *AddHubOptions) error {
	return nil
}

func (n *jsiiProxy_NetworkController) validateAddSpokeParameters(scope constructs.IConstruct, id *string, options *AddNetworkOptions) error {
	return nil
}

func (n *jsiiProxy_NetworkController) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (n *jsiiProxy_NetworkController) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (n *jsiiProxy_NetworkController) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (n *jsiiProxy_NetworkController) validateRegisterAccountParameters(account *string) error {
	return nil
}

func (n *jsiiProxy_NetworkController) validateRegisterCidrParameters(scope constructs.IConstruct, id *string, cidr *string) error {
	return nil
}

func (n *jsiiProxy_NetworkController) validateRegisterRegionParameters(region *string) error {
	return nil
}

func validateNetworkController_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNetworkController_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNetworkController_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewNetworkControllerParameters(scope constructs.IConstruct, id *string, props *NetworkControllerProps) error {
	return nil
}

