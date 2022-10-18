//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package sso

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_Assignment) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (a *jsiiProxy_Assignment) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (a *jsiiProxy_Assignment) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateAssignment_IsConstructParameters(x interface{}) error {
	return nil
}

func validateAssignment_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateAssignment_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewAssignmentParameters(scope constructs.Construct, id *string, props *AssignmentProps) error {
	return nil
}

