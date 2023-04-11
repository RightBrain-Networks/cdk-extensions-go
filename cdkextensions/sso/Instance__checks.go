//go:build !no_runtime_type_checking

package sso

import (
	"fmt"

	"github.com/aws/constructs-go/constructs/v10"
)

func validateInstance_FromArnParameters(scope constructs.IConstruct, id *string, arn *string) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if arn == nil {
		return fmt.Errorf("parameter arn is required, but nil was provided")
	}

	return nil
}

func validateInstance_FromInstanceIdParameters(scope constructs.IConstruct, id *string, instanceId *string) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if instanceId == nil {
		return fmt.Errorf("parameter instanceId is required, but nil was provided")
	}

	return nil
}

