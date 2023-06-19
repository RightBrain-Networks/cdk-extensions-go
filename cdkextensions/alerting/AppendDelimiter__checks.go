//go:build !no_runtime_type_checking

package alerting

import (
	"fmt"
)

func validateAppendDelimiter_OfParameters(delimiter *string) error {
	if delimiter == nil {
		return fmt.Errorf("parameter delimiter is required, but nil was provided")
	}

	return nil
}

