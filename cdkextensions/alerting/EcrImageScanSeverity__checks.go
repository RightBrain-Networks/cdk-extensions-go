//go:build !no_runtime_type_checking

package alerting

import (
	"fmt"
)

func validateEcrImageScanSeverity_OfParameters(name *string, priority *float64, standardized *string) error {
	if name == nil {
		return fmt.Errorf("parameter name is required, but nil was provided")
	}

	if priority == nil {
		return fmt.Errorf("parameter priority is required, but nil was provided")
	}

	if standardized == nil {
		return fmt.Errorf("parameter standardized is required, but nil was provided")
	}

	return nil
}

