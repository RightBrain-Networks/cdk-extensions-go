//go:build no_runtime_type_checking

package alerting

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EcrScanFinding) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (e *jsiiProxy_EcrScanFinding) validateBindParameters(_scope constructs.IConstruct) error {
	return nil
}

func (e *jsiiProxy_EcrScanFinding) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (e *jsiiProxy_EcrScanFinding) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (e *jsiiProxy_EcrScanFinding) validateRegisterIssueTriggerParameters(id *string, options *EcrScanFindingEventOptions) error {
	return nil
}

func validateEcrScanFinding_IsConstructParameters(x interface{}) error {
	return nil
}

func validateEcrScanFinding_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateEcrScanFinding_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewEcrScanFindingParameters(scope constructs.IConstruct, id *string, props *EcrScanFindingProps) error {
	return nil
}

