//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

package ec2

import (
	"fmt"
)

func validateNewFlowLogFieldParameters(name *string, type_ FlowLogDataType) error {
	if name == nil {
		return fmt.Errorf("parameter name is required, but nil was provided")
	}

	if type_ == "" {
		return fmt.Errorf("parameter type_ is required, but nil was provided")
	}

	return nil
}

