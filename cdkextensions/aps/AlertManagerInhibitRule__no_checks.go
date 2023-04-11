//go:build no_runtime_type_checking

package aps

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AlertManagerInhibitRule) validateAddEqualLabelParameters(label *string) error {
	return nil
}

func (a *jsiiProxy_AlertManagerInhibitRule) validateAddSourceMatcherParameters(matcher AlertManagerMatcher) error {
	return nil
}

func (a *jsiiProxy_AlertManagerInhibitRule) validateAddTargetMatcherParameters(matcher AlertManagerMatcher) error {
	return nil
}

func (a *jsiiProxy_AlertManagerInhibitRule) validateBindParameters(_scope constructs.IConstruct) error {
	return nil
}

func validateAlertManagerInhibitRule_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewAlertManagerInhibitRuleParameters(scope AlertManagerConfiguration, id *string, options AlertManagerInhibitRuleProps) error {
	return nil
}

