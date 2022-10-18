//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package sso

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_ReferencedManagedPolicy) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (r *jsiiProxy_ReferencedManagedPolicy) validateAttachToGroupParameters(group awsiam.IGroup) error {
	return nil
}

func (r *jsiiProxy_ReferencedManagedPolicy) validateAttachToRoleParameters(role awsiam.IRole) error {
	return nil
}

func (r *jsiiProxy_ReferencedManagedPolicy) validateAttachToUserParameters(user awsiam.IUser) error {
	return nil
}

func (r *jsiiProxy_ReferencedManagedPolicy) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (r *jsiiProxy_ReferencedManagedPolicy) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateReferencedManagedPolicy_FromAwsManagedPolicyNameParameters(managedPolicyName *string) error {
	return nil
}

func validateReferencedManagedPolicy_FromManagedPolicyArnParameters(scope constructs.Construct, id *string, managedPolicyArn *string) error {
	return nil
}

func validateReferencedManagedPolicy_FromManagedPolicyNameParameters(scope constructs.Construct, id *string, managedPolicyName *string) error {
	return nil
}

func validateReferencedManagedPolicy_IsConstructParameters(x interface{}) error {
	return nil
}

func validateReferencedManagedPolicy_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateReferencedManagedPolicy_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateReferencedManagedPolicy_OfParameters(options *ReferenceOptions) error {
	return nil
}

func validateNewReferencedManagedPolicyParameters(scope constructs.Construct, id *string, props *ReferencedManagedPolicyProps) error {
	return nil
}

