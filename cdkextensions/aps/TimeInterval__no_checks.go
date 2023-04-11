//go:build no_runtime_type_checking

package aps

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TimeInterval) validateAddIntervalParameters(interval TimeIntervalEntry) error {
	return nil
}

func (t *jsiiProxy_TimeInterval) validateBindParameters(scope constructs.IConstruct) error {
	return nil
}

func validateTimeInterval_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewTimeIntervalParameters(scope AlertManagerConfiguration, id *string, options *TimeIntervalProps) error {
	return nil
}

