//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

package k8saws

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
)

func validateFluentBitLogStreamOutput_FromLogStreamParameters(logStream awslogs.ILogStream) error {
	if logStream == nil {
		return fmt.Errorf("parameter logStream is required, but nil was provided")
	}

	return nil
}

func validateFluentBitLogStreamOutput_FromNameParameters(name *string) error {
	if name == nil {
		return fmt.Errorf("parameter name is required, but nil was provided")
	}

	return nil
}

func validateFluentBitLogStreamOutput_FromPrefixParameters(prefix *string) error {
	if prefix == nil {
		return fmt.Errorf("parameter prefix is required, but nil was provided")
	}

	return nil
}

