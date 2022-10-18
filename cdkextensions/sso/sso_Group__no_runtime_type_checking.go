//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package sso

// Building without runtime type checking enabled, so all the below just return nil

func validateGroup_FromGroupIdParameters(scope constructs.IConstruct, id *string, groupId *string) error {
	return nil
}

