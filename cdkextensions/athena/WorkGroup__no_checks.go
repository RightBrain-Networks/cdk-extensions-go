//go:build no_runtime_type_checking

package athena

// Building without runtime type checking enabled, so all the below just return nil

func (w *jsiiProxy_WorkGroup) validateAddNamedQueryParameters(id *string, options *AddNamedQueryOptions) error {
	return nil
}

func (w *jsiiProxy_WorkGroup) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (w *jsiiProxy_WorkGroup) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (w *jsiiProxy_WorkGroup) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateWorkGroup_FromWorkGroupArnParameters(scope constructs.IConstruct, id *string, arn *string) error {
	return nil
}

func validateWorkGroup_FromWorkGroupAttributesParameters(scope constructs.IConstruct, id *string, attrs *WorkGroupAttributes) error {
	return nil
}

func validateWorkGroup_FromWorkGroupNameParameters(scope constructs.IConstruct, id *string, name *string) error {
	return nil
}

func validateWorkGroup_IsConstructParameters(x interface{}) error {
	return nil
}

func validateWorkGroup_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateWorkGroup_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewWorkGroupParameters(scope constructs.IConstruct, id *string, props *WorkGroupProps) error {
	return nil
}

