//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

package k8saws

import (
	"fmt"
)

func validateModifyCondition_AKeyMatchesParameters(regex *string) error {
	if regex == nil {
		return fmt.Errorf("parameter regex is required, but nil was provided")
	}

	return nil
}

func validateModifyCondition_KeyDoesNotExistsParameters(key *string) error {
	if key == nil {
		return fmt.Errorf("parameter key is required, but nil was provided")
	}

	return nil
}

func validateModifyCondition_KeyExistsParameters(key *string) error {
	if key == nil {
		return fmt.Errorf("parameter key is required, but nil was provided")
	}

	return nil
}

func validateModifyCondition_KeyValueDoesNotEqualParameters(key *string, value *string) error {
	if key == nil {
		return fmt.Errorf("parameter key is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateModifyCondition_KeyValueDoesNotMatchParameters(key *string, value *string) error {
	if key == nil {
		return fmt.Errorf("parameter key is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateModifyCondition_KeyValueEqualsParameters(key *string, value *string) error {
	if key == nil {
		return fmt.Errorf("parameter key is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateModifyCondition_KeyValueMatchesParameters(key *string, value *string) error {
	if key == nil {
		return fmt.Errorf("parameter key is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateModifyCondition_MatchingKeysDoNotHaveMatchingValuesParameters(key *string, value *string) error {
	if key == nil {
		return fmt.Errorf("parameter key is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateModifyCondition_MatchingKeysHaveMatchingValuesParameters(key *string, value *string) error {
	if key == nil {
		return fmt.Errorf("parameter key is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateModifyCondition_NoKeyMatchesParameters(regex *string) error {
	if regex == nil {
		return fmt.Errorf("parameter regex is required, but nil was provided")
	}

	return nil
}

func validateModifyCondition_OfParameters(condition *string, args *[]*string) error {
	if condition == nil {
		return fmt.Errorf("parameter condition is required, but nil was provided")
	}

	if args == nil {
		return fmt.Errorf("parameter args is required, but nil was provided")
	}

	return nil
}

