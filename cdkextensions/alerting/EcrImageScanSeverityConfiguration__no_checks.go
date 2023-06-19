//go:build no_runtime_type_checking

package alerting

// Building without runtime type checking enabled, so all the below just return nil

func validateEcrImageScanSeverityConfiguration_ThresholdParameters(level EcrImageScanSeverity) error {
	return nil
}

