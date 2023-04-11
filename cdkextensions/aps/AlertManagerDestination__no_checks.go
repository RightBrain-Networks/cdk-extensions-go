//go:build no_runtime_type_checking

package aps

// Building without runtime type checking enabled, so all the below just return nil

func validateAlertManagerDestination_SnsTopicParameters(topic awssns.ITopic, options *AlertManagerSnsDestinationOptions) error {
	return nil
}

