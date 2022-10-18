//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package sso

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PermissionsBoundary) validateBindParameters(scope constructs.IConstruct) error {
	return nil
}

func validatePermissionsBoundary_FromManagedPolicyParameters(policy awsiam.IManagedPolicy) error {
	return nil
}

func validatePermissionsBoundary_FromReferenceParameters(options *ReferenceOptions) error {
	return nil
}

