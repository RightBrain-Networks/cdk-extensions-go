//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

package sso

import (
	"fmt"
)

func validateAssignmentTarget_AwsAccountParameters(accountId *string) error {
	if accountId == nil {
		return fmt.Errorf("parameter accountId is required, but nil was provided")
	}

	return nil
}

func validateAssignmentTarget_OfParameters(targetType AssignmentTargetType, targetId *string) error {
	if targetType == nil {
		return fmt.Errorf("parameter targetType is required, but nil was provided")
	}

	if targetId == nil {
		return fmt.Errorf("parameter targetId is required, but nil was provided")
	}

	return nil
}

