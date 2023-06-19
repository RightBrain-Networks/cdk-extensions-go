//go:build no_runtime_type_checking

package alerting

// Building without runtime type checking enabled, so all the below just return nil

func validateEcrImageScanSeverity_OfParameters(name *string, priority *float64, standardized *string) error {
	return nil
}

