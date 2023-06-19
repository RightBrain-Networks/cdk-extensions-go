//go:build no_runtime_type_checking

package alerting

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IssueParserPluginBase) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (i *jsiiProxy_IssueParserPluginBase) validateBindParameters(_node constructs.IConstruct) error {
	return nil
}

func (i *jsiiProxy_IssueParserPluginBase) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (i *jsiiProxy_IssueParserPluginBase) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateIssueParserPluginBase_IsConstructParameters(x interface{}) error {
	return nil
}

func validateIssueParserPluginBase_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateIssueParserPluginBase_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewIssueParserPluginBaseParameters(scope constructs.IConstruct, id *string, props *IssueParserPluginBaseProps) error {
	return nil
}

