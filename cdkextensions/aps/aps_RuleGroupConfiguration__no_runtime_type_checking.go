//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package aps

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RuleGroupConfiguration) validateAddRuleGroupParameters(id *string, options *RuleGroupProps) error {
	return nil
}

func (r *jsiiProxy_RuleGroupConfiguration) validateBindParameters(scope constructs.IConstruct) error {
	return nil
}

func validateRuleGroupConfiguration_FromContentParameters(scope constructs.IConstruct, id *string, content *string) error {
	return nil
}

func validateRuleGroupConfiguration_FromRulesFileParameters(scope constructs.IConstruct, id *string, path *string) error {
	return nil
}

func validateRuleGroupConfiguration_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewRuleGroupConfigurationParameters(scope constructs.IConstruct, id *string, _props *RuleGroupConfigurationProps) error {
	return nil
}

