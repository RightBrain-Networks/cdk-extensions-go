//go:build no_runtime_type_checking

package k8saws

// Building without runtime type checking enabled, so all the below just return nil

func validateFluentBitLogStreamOutput_FromLogStreamParameters(logStream awslogs.ILogStream) error {
	return nil
}

func validateFluentBitLogStreamOutput_FromNameParameters(name *string) error {
	return nil
}

func validateFluentBitLogStreamOutput_FromPrefixParameters(prefix *string) error {
	return nil
}

