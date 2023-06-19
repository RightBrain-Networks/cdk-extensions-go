//go:build !no_runtime_type_checking

package alerting

import (
	"fmt"
)

func validateSecurityHubSeverity_OfParameters(name *string, lowerBound *float64, upperBound *float64, standardized *string) error {
	if name == nil {
		return fmt.Errorf("parameter name is required, but nil was provided")
	}

	if lowerBound == nil {
		return fmt.Errorf("parameter lowerBound is required, but nil was provided")
	}

	if upperBound == nil {
		return fmt.Errorf("parameter upperBound is required, but nil was provided")
	}

	if standardized == nil {
		return fmt.Errorf("parameter standardized is required, but nil was provided")
	}

	return nil
}

