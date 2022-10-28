//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package k8saws

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AwsSecretStore) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (a *jsiiProxy_AwsSecretStore) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (a *jsiiProxy_AwsSecretStore) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateAwsSecretStore_IsConstructParameters(x interface{}) error {
	return nil
}

func validateAwsSecretStore_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateAwsSecretStore_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewAwsSecretStoreParameters(scope constructs.Construct, id *string, props *AwsSecretStoreProps) error {
	return nil
}

