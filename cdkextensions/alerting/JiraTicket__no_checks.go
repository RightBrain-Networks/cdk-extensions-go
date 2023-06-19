//go:build no_runtime_type_checking

package alerting

// Building without runtime type checking enabled, so all the below just return nil

func (j *jsiiProxy_JiraTicket) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (j *jsiiProxy_JiraTicket) validateBuildEventOverridesParameters(options *JiraTicketOverrideOptions) error {
	return nil
}

func (j *jsiiProxy_JiraTicket) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (j *jsiiProxy_JiraTicket) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateJiraTicket_IsConstructParameters(x interface{}) error {
	return nil
}

func validateJiraTicket_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateJiraTicket_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewJiraTicketParameters(scope constructs.IConstruct, id *string, props *JiraTicketProps) error {
	return nil
}

