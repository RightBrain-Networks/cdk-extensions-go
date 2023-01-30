//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package aps

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AlertManagerReceiver) validateAddDestinationParameters(destination IAlertManagerDestination) error {
	return nil
}

func (a *jsiiProxy_AlertManagerReceiver) validateAddSnsTopicParameters(topic awssns.ITopic, options *AlertManagerSnsDestinationOptions) error {
	return nil
}

func (a *jsiiProxy_AlertManagerReceiver) validateBindParameters(scope constructs.IConstruct) error {
	return nil
}

func validateAlertManagerReceiver_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewAlertManagerReceiverParameters(scope AlertManagerConfiguration, id *string, props *AlertManagerReceiverProps) error {
	return nil
}

