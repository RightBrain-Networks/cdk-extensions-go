//go:build no_runtime_type_checking

package glue

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_Trigger) validateAddActionParameters(action ITriggerAction) error {
	return nil
}

func (t *jsiiProxy_Trigger) validateAddPredicateParameters(predicate ITriggerPredicate) error {
	return nil
}

func (t *jsiiProxy_Trigger) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (t *jsiiProxy_Trigger) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (t *jsiiProxy_Trigger) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateTrigger_FromTriggerArnParameters(scope constructs.IConstruct, id *string, triggerArn *string) error {
	return nil
}

func validateTrigger_FromTriggerNameParameters(scope constructs.IConstruct, id *string, triggerName *string) error {
	return nil
}

func validateTrigger_IsConstructParameters(x interface{}) error {
	return nil
}

func validateTrigger_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateTrigger_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewTriggerParameters(scope constructs.Construct, id *string, props *TriggerProps) error {
	return nil
}

