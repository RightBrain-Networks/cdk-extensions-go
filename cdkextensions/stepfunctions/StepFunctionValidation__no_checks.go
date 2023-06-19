//go:build no_runtime_type_checking

package stepfunctions

// Building without runtime type checking enabled, so all the below just return nil

func validateStepFunctionValidation_IsIntrinsicParameters(value *string) error {
	return nil
}

func validateStepFunctionValidation_IsJsonPathParameters(value *string) error {
	return nil
}

func validateStepFunctionValidation_IsStatesExpressionParameters(value *string) error {
	return nil
}

