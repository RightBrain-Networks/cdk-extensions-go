//go:build no_runtime_type_checking

package alerting

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IssueManager) validateAddEventRulesParameters(stateMachine awsstepfunctions.StateMachine) error {
	return nil
}

func (i *jsiiProxy_IssueManager) validateAddHandlerParameters(handler IIssueHandler) error {
	return nil
}

func (i *jsiiProxy_IssueManager) validateAddIssueParserParameters(parser IIssueParser) error {
	return nil
}

func (i *jsiiProxy_IssueManager) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (i *jsiiProxy_IssueManager) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (i *jsiiProxy_IssueManager) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateIssueManager_IsConstructParameters(x interface{}) error {
	return nil
}

func validateIssueManager_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateIssueManager_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewIssueManagerParameters(scope constructs.IConstruct, id *string, props *IssueManagerProps) error {
	return nil
}

