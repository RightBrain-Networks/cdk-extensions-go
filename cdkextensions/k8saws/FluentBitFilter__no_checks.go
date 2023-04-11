//go:build no_runtime_type_checking

package k8saws

// Building without runtime type checking enabled, so all the below just return nil

func validateFluentBitFilter_AppendFieldsParameters(match FluentBitMatch, records *[]*AppendedRecord) error {
	return nil
}

func validateFluentBitFilter_BlacklistFieldsParameters(match FluentBitMatch) error {
	return nil
}

func validateFluentBitFilter_GrepParameters(match FluentBitMatch, pattern *FluentBitGrepRegex) error {
	return nil
}

func validateFluentBitFilter_KubernetesParameters(match FluentBitMatch) error {
	return nil
}

func validateFluentBitFilter_LiftParameters(match FluentBitMatch, nestedUnder *string) error {
	return nil
}

func validateFluentBitFilter_ModifyParameters(match FluentBitMatch) error {
	return nil
}

func validateFluentBitFilter_NestParameters(match FluentBitMatch, nestUnder *string) error {
	return nil
}

func validateFluentBitFilter_ParserParameters(match FluentBitMatch, key *string) error {
	return nil
}

func validateFluentBitFilter_RewriteTagParameters(match FluentBitMatch, rules *[]*RewriteTagRule) error {
	return nil
}

func validateFluentBitFilter_ThrottleParameters(match FluentBitMatch, interval awscdk.Duration, rate *float64, window *float64) error {
	return nil
}

func validateFluentBitFilter_WhitelistFieldsParameters(match FluentBitMatch) error {
	return nil
}

