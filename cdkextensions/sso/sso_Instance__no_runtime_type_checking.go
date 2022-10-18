//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package sso

// Building without runtime type checking enabled, so all the below just return nil

func validateInstance_FromArnParameters(scope constructs.IConstruct, id *string, arn *string) error {
	return nil
}

func validateInstance_FromInstanceIdParameters(scope constructs.IConstruct, id *string, instanceId *string) error {
	return nil
}

