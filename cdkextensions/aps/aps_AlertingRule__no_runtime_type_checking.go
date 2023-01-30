//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package aps

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AlertingRule) validateAddAnnotationParameters(label *string, template *string) error {
	return nil
}

func (a *jsiiProxy_AlertingRule) validateAddLabelParameters(label *string, template *string) error {
	return nil
}

func (a *jsiiProxy_AlertingRule) validateBindParameters(_scope constructs.IConstruct) error {
	return nil
}

func validateNewAlertingRuleParameters(props *AlertingRuleProps) error {
	return nil
}

