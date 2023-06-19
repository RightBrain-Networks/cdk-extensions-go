//go:build no_runtime_type_checking

package alerting

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_InspectorFinding) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (i *jsiiProxy_InspectorFinding) validateBindParameters(_node constructs.IConstruct) error {
	return nil
}

func (i *jsiiProxy_InspectorFinding) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (i *jsiiProxy_InspectorFinding) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (i *jsiiProxy_InspectorFinding) validateRegisterIssueTriggerParameters(id *string, options *InspectorFindingEventOptions) error {
	return nil
}

func validateInspectorFinding_IsConstructParameters(x interface{}) error {
	return nil
}

func validateInspectorFinding_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateInspectorFinding_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewInspectorFindingParameters(scope constructs.IConstruct, id *string, props *InspectorFindingProps) error {
	return nil
}

