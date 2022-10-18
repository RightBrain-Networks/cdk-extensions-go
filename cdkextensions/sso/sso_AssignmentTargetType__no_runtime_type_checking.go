//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package sso

// Building without runtime type checking enabled, so all the below just return nil

func validateAssignmentTargetType_OfParameters(name *string) error {
	return nil
}

