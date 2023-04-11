//go:build !no_runtime_type_checking

package k8saws

import (
	"fmt"
)

func validateModifyOperation_AddParameters(key *string, value *string) error {
	if key == nil {
		return fmt.Errorf("parameter key is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateModifyOperation_CopyParameters(originalKey *string, newKey *string) error {
	if originalKey == nil {
		return fmt.Errorf("parameter originalKey is required, but nil was provided")
	}

	if newKey == nil {
		return fmt.Errorf("parameter newKey is required, but nil was provided")
	}

	return nil
}

func validateModifyOperation_HardCopyParameters(originalKey *string, newKey *string) error {
	if originalKey == nil {
		return fmt.Errorf("parameter originalKey is required, but nil was provided")
	}

	if newKey == nil {
		return fmt.Errorf("parameter newKey is required, but nil was provided")
	}

	return nil
}

func validateModifyOperation_HardRenameParameters(originalKey *string, renamedKey *string) error {
	if originalKey == nil {
		return fmt.Errorf("parameter originalKey is required, but nil was provided")
	}

	if renamedKey == nil {
		return fmt.Errorf("parameter renamedKey is required, but nil was provided")
	}

	return nil
}

func validateModifyOperation_MoveToEndParameters(key *string) error {
	if key == nil {
		return fmt.Errorf("parameter key is required, but nil was provided")
	}

	return nil
}

func validateModifyOperation_MoveToStartParameters(key *string) error {
	if key == nil {
		return fmt.Errorf("parameter key is required, but nil was provided")
	}

	return nil
}

func validateModifyOperation_OfParameters(operation *string, args *[]*string) error {
	if operation == nil {
		return fmt.Errorf("parameter operation is required, but nil was provided")
	}

	if args == nil {
		return fmt.Errorf("parameter args is required, but nil was provided")
	}

	return nil
}

func validateModifyOperation_RemoveParameters(key *string) error {
	if key == nil {
		return fmt.Errorf("parameter key is required, but nil was provided")
	}

	return nil
}

func validateModifyOperation_RemoveRegexParameters(regex *string) error {
	if regex == nil {
		return fmt.Errorf("parameter regex is required, but nil was provided")
	}

	return nil
}

func validateModifyOperation_RemoveWildcardParameters(key *string) error {
	if key == nil {
		return fmt.Errorf("parameter key is required, but nil was provided")
	}

	return nil
}

func validateModifyOperation_RenameParameters(originalKey *string, renamedKey *string) error {
	if originalKey == nil {
		return fmt.Errorf("parameter originalKey is required, but nil was provided")
	}

	if renamedKey == nil {
		return fmt.Errorf("parameter renamedKey is required, but nil was provided")
	}

	return nil
}

func validateModifyOperation_SetParameters(key *string, value *string) error {
	if key == nil {
		return fmt.Errorf("parameter key is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

