//go:build no_runtime_type_checking

package alerting

// Building without runtime type checking enabled, so all the below just return nil

func validateSecurityHubSeverityConfiguration_ThresholdParameters(level SecurityHubSeverity) error {
	return nil
}

