//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package ekspatterns

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AwsIntegratedFargateCluster) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (a *jsiiProxy_AwsIntegratedFargateCluster) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (a *jsiiProxy_AwsIntegratedFargateCluster) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (a *jsiiProxy_AwsIntegratedFargateCluster) validateRegisterSecretsManagerSecretParameters(id *string, secret awssecretsmanager.ISecret, options *k8saws.NamespacedExternalSecretOptions) error {
	return nil
}

func (a *jsiiProxy_AwsIntegratedFargateCluster) validateRegisterSsmParameterSecretParameters(id *string, parameter awsssm.IParameter, options *k8saws.NamespacedExternalSecretOptions) error {
	return nil
}

func validateAwsIntegratedFargateCluster_IsConstructParameters(x interface{}) error {
	return nil
}

func validateAwsIntegratedFargateCluster_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateAwsIntegratedFargateCluster_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewAwsIntegratedFargateClusterParameters(scope constructs.Construct, id *string, props *AwsIntegratedFargateClusterProps) error {
	return nil
}

