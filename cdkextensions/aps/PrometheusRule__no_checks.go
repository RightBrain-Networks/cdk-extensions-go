//go:build no_runtime_type_checking

package aps

// Building without runtime type checking enabled, so all the below just return nil

func validatePrometheusRule_AlertingRuleParameters(options *AlertingRuleProps) error {
	return nil
}

func validatePrometheusRule_RecordingRuleParameters(options *RecordingRuleProps) error {
	return nil
}

