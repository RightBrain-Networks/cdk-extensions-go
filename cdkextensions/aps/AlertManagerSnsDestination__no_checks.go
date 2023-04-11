//go:build no_runtime_type_checking

package aps

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AlertManagerSnsDestination) validateAddAttributeParameters(key *string, value *string) error {
	return nil
}

func (a *jsiiProxy_AlertManagerSnsDestination) validateBindParameters(scope constructs.IConstruct) error {
	return nil
}

func validateNewAlertManagerSnsDestinationParameters(options *AlertManagerSnsDestinationProps) error {
	return nil
}

