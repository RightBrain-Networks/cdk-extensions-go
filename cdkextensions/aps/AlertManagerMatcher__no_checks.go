//go:build no_runtime_type_checking

package aps

// Building without runtime type checking enabled, so all the below just return nil

func validateAlertManagerMatcher_FromComponentsParameters(label *string, operator MatchOperator, value *string) error {
	return nil
}

func validateAlertManagerMatcher_FromStringParameters(expression *string) error {
	return nil
}

