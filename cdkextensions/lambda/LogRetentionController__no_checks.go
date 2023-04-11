//go:build no_runtime_type_checking

package lambda

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LogRetentionController) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (l *jsiiProxy_LogRetentionController) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (l *jsiiProxy_LogRetentionController) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateLogRetentionController_IsConstructParameters(x interface{}) error {
	return nil
}

func validateLogRetentionController_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateLogRetentionController_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewLogRetentionControllerParameters(scope constructs.IConstruct, id *string, props *LogRetentionControllerProps) error {
	return nil
}

