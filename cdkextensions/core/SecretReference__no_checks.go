//go:build no_runtime_type_checking

package core

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SecretReference) validateValueForScopeParameters(scope constructs.IConstruct) error {
	return nil
}

func validateSecretReference_StringParameters(scope constructs.IConstruct, key *string, value *string) error {
	return nil
}

