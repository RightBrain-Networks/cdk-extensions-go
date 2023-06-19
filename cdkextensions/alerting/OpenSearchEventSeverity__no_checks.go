//go:build no_runtime_type_checking

package alerting

// Building without runtime type checking enabled, so all the below just return nil

func (o *jsiiProxy_OpenSearchEventSeverity) validateBuildConditionParameters(path *string) error {
	return nil
}

func validateOpenSearchEventSeverity_OfParameters(standardized *string, original *string, priority *float64) error {
	return nil
}

func validateOpenSearchEventSeverity_ThresholdParameters(level OpenSearchEventSeverity) error {
	return nil
}

