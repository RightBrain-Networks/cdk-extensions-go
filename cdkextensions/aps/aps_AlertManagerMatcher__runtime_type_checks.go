//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

package aps

import (
	"fmt"
)

func validateAlertManagerMatcher_FromComponentsParameters(label *string, operator MatchOperator, value *string) error {
	if label == nil {
		return fmt.Errorf("parameter label is required, but nil was provided")
	}

	if operator == "" {
		return fmt.Errorf("parameter operator is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateAlertManagerMatcher_FromStringParameters(expression *string) error {
	if expression == nil {
		return fmt.Errorf("parameter expression is required, but nil was provided")
	}

	return nil
}

