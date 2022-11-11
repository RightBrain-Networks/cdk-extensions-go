//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

package k8saws

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
)

func validateFluentBitLogGroupOutput_FromLogGroupParameters(logGroup awslogs.ILogGroup) error {
	if logGroup == nil {
		return fmt.Errorf("parameter logGroup is required, but nil was provided")
	}

	return nil
}

func validateFluentBitLogGroupOutput_FromNameParameters(name *string) error {
	if name == nil {
		return fmt.Errorf("parameter name is required, but nil was provided")
	}

	return nil
}

