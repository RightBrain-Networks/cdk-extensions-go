//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package k8saws

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_ExternalDns) validateAddDomainFilterParameters(domain *string) error {
	return nil
}

func (e *jsiiProxy_ExternalDns) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (e *jsiiProxy_ExternalDns) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (e *jsiiProxy_ExternalDns) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateExternalDns_IsConstructParameters(x interface{}) error {
	return nil
}

func validateExternalDns_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateExternalDns_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewExternalDnsParameters(scope constructs.Construct, id *string, props *ExternalDnsProps) error {
	return nil
}

