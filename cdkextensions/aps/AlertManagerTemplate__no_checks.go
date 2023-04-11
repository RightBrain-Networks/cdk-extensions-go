//go:build no_runtime_type_checking

package aps

// Building without runtime type checking enabled, so all the below just return nil

func validateAlertManagerTemplate_FromFileParameters(path *string) error {
	return nil
}

func validateAlertManagerTemplate_FromStringParameters(content *string) error {
	return nil
}

