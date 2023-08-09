//go:build no_runtime_type_checking

package ssm

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AutomationDocument) validateAddRequirementParameters(requirement *DocumentRequirement) error {
	return nil
}

func (a *jsiiProxy_AutomationDocument) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (a *jsiiProxy_AutomationDocument) validateArnForAutomationDefinitionVersionParameters(version *string) error {
	return nil
}

func (a *jsiiProxy_AutomationDocument) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (a *jsiiProxy_AutomationDocument) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateAutomationDocument_FromManagedParameters(scope constructs.IConstruct, id *string, managedDocumentName *string) error {
	return nil
}

func validateAutomationDocument_IsConstructParameters(x interface{}) error {
	return nil
}

func validateAutomationDocument_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateAutomationDocument_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

