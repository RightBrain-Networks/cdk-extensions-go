//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package aps

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RuleGroup) validateAddAlertingRuleParameters(options *AlertingRuleProps) error {
	return nil
}

func (r *jsiiProxy_RuleGroup) validateAddRecordingRuleParameters(options *RecordingRuleProps) error {
	return nil
}

func (r *jsiiProxy_RuleGroup) validateAddRuleParameters(rule IPrometheusRule) error {
	return nil
}

func (r *jsiiProxy_RuleGroup) validateBindParameters(scope constructs.IConstruct) error {
	return nil
}

func validateRuleGroup_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewRuleGroupParameters(scope constructs.IConstruct, id *string, props *RuleGroupProps) error {
	return nil
}

