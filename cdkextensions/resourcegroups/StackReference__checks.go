//go:build !no_runtime_type_checking

package resourcegroups

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2"
)

func validateStackReference_FromStackParameters(stack awscdk.Stack) error {
	if stack == nil {
		return fmt.Errorf("parameter stack is required, but nil was provided")
	}

	return nil
}

func validateStackReference_FromStackIdParameters(stackId *string) error {
	if stackId == nil {
		return fmt.Errorf("parameter stackId is required, but nil was provided")
	}

	return nil
}

