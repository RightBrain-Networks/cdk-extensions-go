//go:build no_runtime_type_checking

package k8saws

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SecretsManagerReference) validateAddFieldMappingParameters(field *SecretFieldReference) error {
	return nil
}

func (s *jsiiProxy_SecretsManagerReference) validateBindParameters(_scope constructs.IConstruct) error {
	return nil
}

func validateNewSecretsManagerReferenceParameters(secret awssecretsmanager.ISecret, options *SecretsManagerReferenceOptions) error {
	return nil
}

