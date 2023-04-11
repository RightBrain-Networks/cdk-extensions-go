//go:build !no_runtime_type_checking

package aps

import (
	"fmt"
)

func validateAlertManagerTemplate_FromFileParameters(path *string) error {
	if path == nil {
		return fmt.Errorf("parameter path is required, but nil was provided")
	}

	return nil
}

func validateAlertManagerTemplate_FromStringParameters(content *string) error {
	if content == nil {
		return fmt.Errorf("parameter content is required, but nil was provided")
	}

	return nil
}

