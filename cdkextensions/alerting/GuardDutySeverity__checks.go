//go:build !no_runtime_type_checking

package alerting

import (
	"fmt"
)

func (g *jsiiProxy_GuardDutySeverity) validateBuildConditionParameters(path *string) error {
	if path == nil {
		return fmt.Errorf("parameter path is required, but nil was provided")
	}

	return nil
}

func validateGuardDutySeverity_OfParameters(standardized *string) error {
	if standardized == nil {
		return fmt.Errorf("parameter standardized is required, but nil was provided")
	}

	return nil
}

func validateGuardDutySeverity_ThresholdParameters(level GuardDutySeverity) error {
	if level == nil {
		return fmt.Errorf("parameter level is required, but nil was provided")
	}

	return nil
}

