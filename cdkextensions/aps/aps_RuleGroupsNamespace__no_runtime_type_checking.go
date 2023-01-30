//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package aps

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RuleGroupsNamespace) validateAddRuleGroupParameters(id *string, options *RuleGroupProps) error {
	return nil
}

func (r *jsiiProxy_RuleGroupsNamespace) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (r *jsiiProxy_RuleGroupsNamespace) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (r *jsiiProxy_RuleGroupsNamespace) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateRuleGroupsNamespace_FromRuleGroupsNamespaceArnParameters(scope constructs.IConstruct, id *string, arn *string) error {
	return nil
}

func validateRuleGroupsNamespace_IsConstructParameters(x interface{}) error {
	return nil
}

func validateRuleGroupsNamespace_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateRuleGroupsNamespace_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewRuleGroupsNamespaceParameters(scope constructs.IConstruct, id *string, props *RuleGroupsNamespaceProps) error {
	return nil
}

