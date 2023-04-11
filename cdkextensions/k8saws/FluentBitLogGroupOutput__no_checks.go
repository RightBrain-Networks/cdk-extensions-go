//go:build no_runtime_type_checking

package k8saws

// Building without runtime type checking enabled, so all the below just return nil

func validateFluentBitLogGroupOutput_FromLogGroupParameters(logGroup awslogs.ILogGroup) error {
	return nil
}

func validateFluentBitLogGroupOutput_FromNameParameters(name *string) error {
	return nil
}

