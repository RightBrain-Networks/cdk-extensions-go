//go:build !no_runtime_type_checking

package alerting

import (
	"fmt"
)

func validateEcrImageScanSeverityConfiguration_ThresholdParameters(level EcrImageScanSeverity) error {
	if level == nil {
		return fmt.Errorf("parameter level is required, but nil was provided")
	}

	return nil
}

