//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package core

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_State) validateGetParameters(key *string) error {
	return nil
}

func (s *jsiiProxy_State) validateSetParameters(key *string, value interface{}) error {
	return nil
}

func validateState_OfParameters(scope constructs.IConstruct) error {
	return nil
}

