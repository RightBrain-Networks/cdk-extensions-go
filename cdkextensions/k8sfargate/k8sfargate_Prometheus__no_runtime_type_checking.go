//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package k8sfargate

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_Prometheus) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (p *jsiiProxy_Prometheus) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (p *jsiiProxy_Prometheus) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validatePrometheus_IsConstructParameters(x interface{}) error {
	return nil
}

func validatePrometheus_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validatePrometheus_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewPrometheusParameters(scope constructs.IConstruct, id *string, props *PrometheusProps) error {
	return nil
}

