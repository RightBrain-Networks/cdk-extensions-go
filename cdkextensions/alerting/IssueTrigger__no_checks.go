//go:build no_runtime_type_checking

package alerting

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IssueTrigger) validateAddOverrideParameters(handlerOverrides IssueHandlerOverride) error {
	return nil
}

func (i *jsiiProxy_IssueTrigger) validateBindParameters(stateMachine awsstepfunctions.StateMachine) error {
	return nil
}

func validateIssueTrigger_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewIssueTriggerParameters(scope constructs.IConstruct, id *string, props *IssueTriggerProps) error {
	return nil
}

