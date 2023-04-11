//go:build no_runtime_type_checking

package sso

// Building without runtime type checking enabled, so all the below just return nil

func validateAssignmentTarget_AwsAccountParameters(accountId *string) error {
	return nil
}

func validateAssignmentTarget_OfParameters(targetType AssignmentTargetType, targetId *string) error {
	return nil
}

