//go:build no_runtime_type_checking

package ec2

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CustomerGateway) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (c *jsiiProxy_CustomerGateway) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (c *jsiiProxy_CustomerGateway) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateCustomerGateway_FromCustomerGatewayAttributesParameters(scope constructs.IConstruct, id *string, attributes *CustomerGatewayAttributes) error {
	return nil
}

func validateCustomerGateway_FromCustomerGatewayIdParameters(scope constructs.IConstruct, id *string, customerGatewayId *string) error {
	return nil
}

func validateCustomerGateway_IsConstructParameters(x interface{}) error {
	return nil
}

func validateCustomerGateway_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCustomerGateway_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewCustomerGatewayParameters(scope constructs.Construct, id *string, props *CustomerGatewayProps) error {
	return nil
}

