//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package asserts

// Building without runtime type checking enabled, so all the below just return nil

func (j *jsiiProxy_JoinedJson) validateTestParameters(actual interface{}) error {
	return nil
}

func validateJoinedJson_IsMatcherParameters(x interface{}) error {
	return nil
}

func validateNewJoinedJsonParameters(pattern interface{}) error {
	return nil
}

