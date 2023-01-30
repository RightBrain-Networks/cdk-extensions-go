//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package aps

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AlertManagerRoute) validateAddActiveTimeIntervalParameters(interval TimeInterval) error {
	return nil
}

func (a *jsiiProxy_AlertManagerRoute) validateAddChildParameters(id *string, options *AlertManagerRouteProps) error {
	return nil
}

func (a *jsiiProxy_AlertManagerRoute) validateAddGroupByLabelParameters(label *string) error {
	return nil
}

func (a *jsiiProxy_AlertManagerRoute) validateAddMatcherParameters(matcher AlertManagerMatcher) error {
	return nil
}

func (a *jsiiProxy_AlertManagerRoute) validateAddMuteTimeIntervalParameters(interval TimeInterval) error {
	return nil
}

func (a *jsiiProxy_AlertManagerRoute) validateBindParameters(scope constructs.IConstruct) error {
	return nil
}

func validateAlertManagerRoute_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewAlertManagerRouteParameters(scope interface{}, id *string, options *AlertManagerRouteProps) error {
	return nil
}

