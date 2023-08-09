//go:build no_runtime_type_checking

package securityhub

// Building without runtime type checking enabled, so all the below just return nil

func (h *jsiiProxy_Hub) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (h *jsiiProxy_Hub) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (h *jsiiProxy_Hub) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateHub_FromHubArnParameters(scope constructs.IConstruct, id *string, arn *string) error {
	return nil
}

func validateHub_FromHubAttributesParameters(scope constructs.IConstruct, id *string, attrs *HubAttributes) error {
	return nil
}

func validateHub_FromHubNameParameters(scope constructs.IConstruct, id *string, name *string) error {
	return nil
}

func validateHub_IsConstructParameters(x interface{}) error {
	return nil
}

func validateHub_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateHub_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewHubParameters(scope constructs.IConstruct, id *string, props *HubProps) error {
	return nil
}

