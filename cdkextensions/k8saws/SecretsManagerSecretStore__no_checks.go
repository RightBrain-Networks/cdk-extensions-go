//go:build no_runtime_type_checking

package k8saws

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SecretsManagerSecretStore) validateAddSecretParameters(id *string, secret awssecretsmanager.ISecret, options *ExternalSecretOptions) error {
	return nil
}

func (s *jsiiProxy_SecretsManagerSecretStore) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (s *jsiiProxy_SecretsManagerSecretStore) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (s *jsiiProxy_SecretsManagerSecretStore) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateSecretsManagerSecretStore_IsConstructParameters(x interface{}) error {
	return nil
}

func validateSecretsManagerSecretStore_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateSecretsManagerSecretStore_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewSecretsManagerSecretStoreParameters(scope constructs.Construct, id *string, props *SecretsManagerSecretStoreProps) error {
	return nil
}

