//go:build no_runtime_type_checking

package k8saws

// Building without runtime type checking enabled, so all the below just return nil

func validateFluentBitMatch_GlobParameters(pattern *string) error {
	return nil
}

func validateFluentBitMatch_RegexParameters(pattern *string) error {
	return nil
}

