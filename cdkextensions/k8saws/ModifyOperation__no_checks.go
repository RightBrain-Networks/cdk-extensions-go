//go:build no_runtime_type_checking

package k8saws

// Building without runtime type checking enabled, so all the below just return nil

func validateModifyOperation_AddParameters(key *string, value *string) error {
	return nil
}

func validateModifyOperation_CopyParameters(originalKey *string, newKey *string) error {
	return nil
}

func validateModifyOperation_HardCopyParameters(originalKey *string, newKey *string) error {
	return nil
}

func validateModifyOperation_HardRenameParameters(originalKey *string, renamedKey *string) error {
	return nil
}

func validateModifyOperation_MoveToEndParameters(key *string) error {
	return nil
}

func validateModifyOperation_MoveToStartParameters(key *string) error {
	return nil
}

func validateModifyOperation_OfParameters(operation *string, args *[]*string) error {
	return nil
}

func validateModifyOperation_RemoveParameters(key *string) error {
	return nil
}

func validateModifyOperation_RemoveRegexParameters(regex *string) error {
	return nil
}

func validateModifyOperation_RemoveWildcardParameters(key *string) error {
	return nil
}

func validateModifyOperation_RenameParameters(originalKey *string, renamedKey *string) error {
	return nil
}

func validateModifyOperation_SetParameters(key *string, value *string) error {
	return nil
}

