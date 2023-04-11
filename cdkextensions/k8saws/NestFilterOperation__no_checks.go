//go:build no_runtime_type_checking

package k8saws

// Building without runtime type checking enabled, so all the below just return nil

func validateNestFilterOperation_LiftParameters(options *LiftOptions) error {
	return nil
}

func validateNestFilterOperation_NestParameters(options *NestOptions) error {
	return nil
}

