//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package ram

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_ResourceShare) validateAddPrincipalParameters(principal ISharedPrincipal) error {
	return nil
}

func (r *jsiiProxy_ResourceShare) validateAddResourceParameters(resource ISharedResource) error {
	return nil
}

func (r *jsiiProxy_ResourceShare) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (r *jsiiProxy_ResourceShare) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (r *jsiiProxy_ResourceShare) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateResourceShare_IsConstructParameters(x interface{}) error {
	return nil
}

func validateResourceShare_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateResourceShare_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewResourceShareParameters(scope constructs.Construct, id *string, props *ResourceShareProps) error {
	return nil
}

