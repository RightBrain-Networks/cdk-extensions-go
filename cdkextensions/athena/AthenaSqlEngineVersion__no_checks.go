//go:build no_runtime_type_checking

package athena

// Building without runtime type checking enabled, so all the below just return nil

func validateAthenaSqlEngineVersion_OfParameters(name *string) error {
	return nil
}

