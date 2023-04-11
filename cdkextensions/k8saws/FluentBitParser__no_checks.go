//go:build no_runtime_type_checking

package k8saws

// Building without runtime type checking enabled, so all the below just return nil

func validateFluentBitParser_JsonParameters(name *string) error {
	return nil
}

func validateFluentBitParser_LogfmtParameters(name *string) error {
	return nil
}

func validateFluentBitParser_LtsvParameters(name *string) error {
	return nil
}

func validateFluentBitParser_RegexParameters(name *string, regex *string) error {
	return nil
}

