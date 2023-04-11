//go:build no_runtime_type_checking

package sso

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AccessControlAttribute) validateAddSourceParameters(source *string) error {
	return nil
}

func (a *jsiiProxy_AccessControlAttribute) validateBindParameters(scope constructs.IConstruct) error {
	return nil
}

func validateNewAccessControlAttributeParameters(options *AccessControlAttributeOptions) error {
	return nil
}

