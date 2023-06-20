//go:build no_runtime_type_checking

package stepfunctions

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_StringReplace) validateNextParameters(state awsstepfunctions.IChainable) error {
	return nil
}

func validateStringReplace_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewStringReplaceParameters(scope constructs.IConstruct, id *string, props *StringReplaceProps) error {
	return nil
}

