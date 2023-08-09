//go:build no_runtime_type_checking

package ssm

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IAutomationDocument) validateArnForAutomationDefinitionVersionParameters(version *string) error {
	return nil
}

