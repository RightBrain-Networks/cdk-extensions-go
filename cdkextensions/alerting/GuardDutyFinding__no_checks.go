//go:build no_runtime_type_checking

package alerting

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_GuardDutyFinding) validateAddSectionFieldParameters(id *string, key *string, path *string) error {
	return nil
}

func (g *jsiiProxy_GuardDutyFinding) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (g *jsiiProxy_GuardDutyFinding) validateBindParameters(_scope constructs.IConstruct) error {
	return nil
}

func (g *jsiiProxy_GuardDutyFinding) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (g *jsiiProxy_GuardDutyFinding) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (g *jsiiProxy_GuardDutyFinding) validateRegisterIssueTriggerParameters(id *string, options *GuardDutyFindingRuleOptions) error {
	return nil
}

func validateGuardDutyFinding_IsConstructParameters(x interface{}) error {
	return nil
}

func validateGuardDutyFinding_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateGuardDutyFinding_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewGuardDutyFindingParameters(scope constructs.IConstruct, id *string, props *GuardDutyFindingProps) error {
	return nil
}

