//go:build no_runtime_type_checking

package alerting

// Building without runtime type checking enabled, so all the below just return nil

func (o *jsiiProxy_OpenSearchEvent) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (o *jsiiProxy_OpenSearchEvent) validateBindParameters(_node constructs.IConstruct) error {
	return nil
}

func (o *jsiiProxy_OpenSearchEvent) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (o *jsiiProxy_OpenSearchEvent) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (o *jsiiProxy_OpenSearchEvent) validateRegisterIssueTriggerParameters(id *string, options *OpenSearchEventRuleOptions) error {
	return nil
}

func validateOpenSearchEvent_IsConstructParameters(x interface{}) error {
	return nil
}

func validateOpenSearchEvent_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateOpenSearchEvent_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewOpenSearchEventParameters(scope constructs.IConstruct, id *string, props *OpenSearchEventProps) error {
	return nil
}

