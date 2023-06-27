//go:build no_runtime_type_checking

package guardduty

// Building without runtime type checking enabled, so all the below just return nil

func validateFindingPublishingFrequency_OfParameters(label *string) error {
	return nil
}

