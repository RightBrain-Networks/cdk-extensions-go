//go:build !no_runtime_type_checking

package kinesisfirehose

import (
	"fmt"
)

func validateTableVersion_FixedParameters(version *float64) error {
	if version == nil {
		return fmt.Errorf("parameter version is required, but nil was provided")
	}

	return nil
}

