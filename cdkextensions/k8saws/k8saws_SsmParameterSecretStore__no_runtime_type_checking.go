//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package k8saws

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SsmParameterSecretStore) validateAddSecretParameters(id *string, parameter awsssm.IParameter, options *ExternalSecretOptions) error {
	return nil
}

func (s *jsiiProxy_SsmParameterSecretStore) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (s *jsiiProxy_SsmParameterSecretStore) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (s *jsiiProxy_SsmParameterSecretStore) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateSsmParameterSecretStore_IsConstructParameters(x interface{}) error {
	return nil
}

func validateSsmParameterSecretStore_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateSsmParameterSecretStore_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewSsmParameterSecretStoreParameters(scope constructs.Construct, id *string, props *SsmParameterSecretStoreProps) error {
	return nil
}

