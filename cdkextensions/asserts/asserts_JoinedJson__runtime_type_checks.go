//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

package asserts

import (
	"fmt"
)

func (j *jsiiProxy_JoinedJson) validateTestParameters(actual interface{}) error {
	if actual == nil {
		return fmt.Errorf("parameter actual is required, but nil was provided")
	}

	return nil
}

func validateJoinedJson_IsMatcherParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateNewJoinedJsonParameters(pattern interface{}) error {
	if pattern == nil {
		return fmt.Errorf("parameter pattern is required, but nil was provided")
	}

	return nil
}

