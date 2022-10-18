//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package kinesisfirehose

// Building without runtime type checking enabled, so all the below just return nil

func validateTableVersion_FixedParameters(version *float64) error {
	return nil
}

