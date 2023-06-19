//go:build !no_runtime_type_checking

package alerting

import (
	"fmt"
)

func (o *jsiiProxy_OpenSearchEventSeverity) validateBuildConditionParameters(path *string) error {
	if path == nil {
		return fmt.Errorf("parameter path is required, but nil was provided")
	}

	return nil
}

func validateOpenSearchEventSeverity_OfParameters(standardized *string, original *string, priority *float64) error {
	if standardized == nil {
		return fmt.Errorf("parameter standardized is required, but nil was provided")
	}

	if original == nil {
		return fmt.Errorf("parameter original is required, but nil was provided")
	}

	if priority == nil {
		return fmt.Errorf("parameter priority is required, but nil was provided")
	}

	return nil
}

func validateOpenSearchEventSeverity_ThresholdParameters(level OpenSearchEventSeverity) error {
	if level == nil {
		return fmt.Errorf("parameter level is required, but nil was provided")
	}

	return nil
}

