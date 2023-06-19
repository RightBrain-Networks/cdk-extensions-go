//go:build no_runtime_type_checking

package alerting

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IssuePluginBase) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (i *jsiiProxy_IssuePluginBase) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (i *jsiiProxy_IssuePluginBase) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateIssuePluginBase_IsConstructParameters(x interface{}) error {
	return nil
}

func validateIssuePluginBase_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateIssuePluginBase_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewIssuePluginBaseParameters(scope constructs.IConstruct, id *string, props *IssuePluginBaseProps) error {
	return nil
}

