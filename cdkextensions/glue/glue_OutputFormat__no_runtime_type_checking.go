//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package glue

// Building without runtime type checking enabled, so all the below just return nil

func validateNewOutputFormatParameters(className *string) error {
	return nil
}

