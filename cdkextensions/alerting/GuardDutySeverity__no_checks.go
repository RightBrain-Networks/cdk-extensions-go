//go:build no_runtime_type_checking

package alerting

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_GuardDutySeverity) validateBuildConditionParameters(path *string) error {
	return nil
}

func validateGuardDutySeverity_OfParameters(standardized *string) error {
	return nil
}

func validateGuardDutySeverity_ThresholdParameters(level GuardDutySeverity) error {
	return nil
}

