//go:build no_runtime_type_checking

package k8saws

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_ExternalSecret) validateAddSecretParameters(secret ISecretReference) error {
	return nil
}

func (e *jsiiProxy_ExternalSecret) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (e *jsiiProxy_ExternalSecret) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (e *jsiiProxy_ExternalSecret) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateExternalSecret_IsConstructParameters(x interface{}) error {
	return nil
}

func validateExternalSecret_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateExternalSecret_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewExternalSecretParameters(scope constructs.Construct, id *string, props *ExternalSecretProps) error {
	return nil
}

