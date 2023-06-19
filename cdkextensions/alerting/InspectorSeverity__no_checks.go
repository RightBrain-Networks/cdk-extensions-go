//go:build no_runtime_type_checking

package alerting

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_InspectorSeverity) validateBuildConditionParameters(path *string) error {
	return nil
}

func validateInspectorSeverity_OfParameters(standardized *string, original *string, priority *float64) error {
	return nil
}

func validateInspectorSeverity_ThresholdParameters(level InspectorSeverity) error {
	return nil
}

