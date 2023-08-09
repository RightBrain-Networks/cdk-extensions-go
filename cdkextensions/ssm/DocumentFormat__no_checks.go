//go:build no_runtime_type_checking

package ssm

// Building without runtime type checking enabled, so all the below just return nil

func validateDocumentFormat_OfParameters(value *string) error {
	return nil
}

