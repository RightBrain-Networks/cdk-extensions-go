//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package ram

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SharedPrincipal) validateBindParameters(_scope constructs.IConstruct) error {
	return nil
}

func validateSharedPrincipal_FromAccountIdParameters(account *string) error {
	return nil
}

func validateSharedPrincipal_FromConstructParameters(construct constructs.IConstruct) error {
	return nil
}

func validateSharedPrincipal_FromOrganizationalUnitArnParameters(arn *string) error {
	return nil
}

func validateSharedPrincipal_FromOrganizationArnParameters(arn *string) error {
	return nil
}

func validateSharedPrincipal_FromRoleParameters(role awsiam.IRole) error {
	return nil
}

func validateSharedPrincipal_FromStageParameters(stage awscdk.Stage) error {
	return nil
}

func validateSharedPrincipal_FromUserParameters(user awsiam.IUser) error {
	return nil
}

