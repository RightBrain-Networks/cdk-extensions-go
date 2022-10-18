//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package sso

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_ReferencedPermissionsBoundary) validateBindParameters(_scope constructs.IConstruct) error {
	return nil
}

func validateReferencedPermissionsBoundary_FromManagedPolicyParameters(policy awsiam.IManagedPolicy) error {
	return nil
}

func validateReferencedPermissionsBoundary_FromReferenceParameters(options *ReferenceOptions) error {
	return nil
}

func validateNewReferencedPermissionsBoundaryParameters(options *ReferenceOptions) error {
	return nil
}

