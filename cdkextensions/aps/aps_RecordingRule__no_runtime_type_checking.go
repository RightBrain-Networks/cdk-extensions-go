//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package aps

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RecordingRule) validateAddLabelParameters(label *string, value *string) error {
	return nil
}

func (r *jsiiProxy_RecordingRule) validateBindParameters(_scope constructs.IConstruct) error {
	return nil
}

func validateNewRecordingRuleParameters(props *RecordingRuleProps) error {
	return nil
}

