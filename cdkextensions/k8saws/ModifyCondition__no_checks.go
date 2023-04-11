//go:build no_runtime_type_checking

package k8saws

// Building without runtime type checking enabled, so all the below just return nil

func validateModifyCondition_AKeyMatchesParameters(regex *string) error {
	return nil
}

func validateModifyCondition_KeyDoesNotExistsParameters(key *string) error {
	return nil
}

func validateModifyCondition_KeyExistsParameters(key *string) error {
	return nil
}

func validateModifyCondition_KeyValueDoesNotEqualParameters(key *string, value *string) error {
	return nil
}

func validateModifyCondition_KeyValueDoesNotMatchParameters(key *string, value *string) error {
	return nil
}

func validateModifyCondition_KeyValueEqualsParameters(key *string, value *string) error {
	return nil
}

func validateModifyCondition_KeyValueMatchesParameters(key *string, value *string) error {
	return nil
}

func validateModifyCondition_MatchingKeysDoNotHaveMatchingValuesParameters(key *string, value *string) error {
	return nil
}

func validateModifyCondition_MatchingKeysHaveMatchingValuesParameters(key *string, value *string) error {
	return nil
}

func validateModifyCondition_NoKeyMatchesParameters(regex *string) error {
	return nil
}

func validateModifyCondition_OfParameters(condition *string, args *[]*string) error {
	return nil
}

