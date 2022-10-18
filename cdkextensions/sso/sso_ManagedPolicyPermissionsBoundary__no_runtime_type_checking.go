//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package sso

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_ManagedPolicyPermissionsBoundary) validateBindParameters(_scope constructs.IConstruct) error {
	return nil
}

func validateManagedPolicyPermissionsBoundary_FromManagedPolicyParameters(policy awsiam.IManagedPolicy) error {
	return nil
}

func validateManagedPolicyPermissionsBoundary_FromReferenceParameters(options *ReferenceOptions) error {
	return nil
}

func validateNewManagedPolicyPermissionsBoundaryParameters(policy awsiam.IManagedPolicy) error {
	return nil
}

