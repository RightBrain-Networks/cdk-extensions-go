//go:build no_runtime_type_checking

package resourcegroups

// Building without runtime type checking enabled, so all the below just return nil

func validateStackReference_FromStackParameters(stack awscdk.Stack) error {
	return nil
}

func validateStackReference_FromStackIdParameters(stackId *string) error {
	return nil
}

