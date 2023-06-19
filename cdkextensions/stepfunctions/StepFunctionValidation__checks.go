//go:build !no_runtime_type_checking

package stepfunctions

import (
	"fmt"
)

func validateStepFunctionValidation_IsIntrinsicParameters(value *string) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateStepFunctionValidation_IsJsonPathParameters(value *string) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateStepFunctionValidation_IsStatesExpressionParameters(value *string) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

