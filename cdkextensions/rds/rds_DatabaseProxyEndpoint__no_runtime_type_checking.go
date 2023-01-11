//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package rds

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DatabaseProxyEndpoint) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (d *jsiiProxy_DatabaseProxyEndpoint) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (d *jsiiProxy_DatabaseProxyEndpoint) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateDatabaseProxyEndpoint_IsConstructParameters(x interface{}) error {
	return nil
}

func validateDatabaseProxyEndpoint_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateDatabaseProxyEndpoint_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewDatabaseProxyEndpointParameters(scope constructs.Construct, id *string, props *DatabaseProxyEndpointProps) error {
	return nil
}

