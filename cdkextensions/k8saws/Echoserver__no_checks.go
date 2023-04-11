//go:build no_runtime_type_checking

package k8saws

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_Echoserver) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (e *jsiiProxy_Echoserver) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (e *jsiiProxy_Echoserver) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (e *jsiiProxy_Echoserver) validateRegisterDomainParameters(domain route53.Domain) error {
	return nil
}

func validateEchoserver_IsConstructParameters(x interface{}) error {
	return nil
}

func validateEchoserver_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateEchoserver_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewEchoserverParameters(scope constructs.Construct, id *string, props *EchoserverProps) error {
	return nil
}

