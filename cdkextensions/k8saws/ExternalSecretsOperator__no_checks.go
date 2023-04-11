//go:build no_runtime_type_checking

package k8saws

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_ExternalSecretsOperator) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (e *jsiiProxy_ExternalSecretsOperator) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (e *jsiiProxy_ExternalSecretsOperator) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (e *jsiiProxy_ExternalSecretsOperator) validateRegisterSecretsManagerSecretParameters(id *string, secret awssecretsmanager.ISecret, options *NamespacedExternalSecretOptions) error {
	return nil
}

func (e *jsiiProxy_ExternalSecretsOperator) validateRegisterSsmParameterSecretParameters(id *string, parameter awsssm.IParameter, options *NamespacedExternalSecretOptions) error {
	return nil
}

func validateExternalSecretsOperator_IsConstructParameters(x interface{}) error {
	return nil
}

func validateExternalSecretsOperator_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateExternalSecretsOperator_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewExternalSecretsOperatorParameters(scope constructs.Construct, id *string, props *ExternalSecretsOperatorProps) error {
	return nil
}

