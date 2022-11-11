//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package k8saws

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AdotCollector) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (a *jsiiProxy_AdotCollector) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (a *jsiiProxy_AdotCollector) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateAdotCollector_IsConstructParameters(x interface{}) error {
	return nil
}

func validateAdotCollector_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateAdotCollector_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewAdotCollectorParameters(scope constructs.Construct, id *string, props *AdotCollectorProps) error {
	return nil
}

