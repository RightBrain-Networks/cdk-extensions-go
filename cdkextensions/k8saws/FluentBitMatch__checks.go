//go:build !no_runtime_type_checking

package k8saws

import (
	"fmt"
)

func validateFluentBitMatch_GlobParameters(pattern *string) error {
	if pattern == nil {
		return fmt.Errorf("parameter pattern is required, but nil was provided")
	}

	return nil
}

func validateFluentBitMatch_RegexParameters(pattern *string) error {
	if pattern == nil {
		return fmt.Errorf("parameter pattern is required, but nil was provided")
	}

	return nil
}

