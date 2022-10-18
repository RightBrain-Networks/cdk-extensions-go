//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package ec2

// Building without runtime type checking enabled, so all the below just return nil

func validateNewFlowLogFieldParameters(name *string, type_ FlowLogDataType) error {
	return nil
}

