//go:build no_runtime_type_checking

package rds

// Building without runtime type checking enabled, so all the below just return nil

func validateDatabaseProxyEndpointAccess_OfParameters(role *string) error {
	return nil
}

