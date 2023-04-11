//go:build no_runtime_type_checking

package sso

// Building without runtime type checking enabled, so all the below just return nil

func validateUser_FromUserIdParameters(scope constructs.IConstruct, id *string, userId *string) error {
	return nil
}

