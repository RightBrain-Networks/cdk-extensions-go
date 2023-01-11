//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

package rds

import (
	"fmt"
)

func validateDatabaseProxyEndpointAccess_OfParameters(role *string) error {
	if role == nil {
		return fmt.Errorf("parameter role is required, but nil was provided")
	}

	return nil
}

