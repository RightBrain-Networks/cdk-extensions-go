//go:build no_runtime_type_checking

package config

// Building without runtime type checking enabled, so all the below just return nil

func validateRemediationTargetType_OfParameters(value *string) error {
	return nil
}

