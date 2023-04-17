//go:build no_runtime_type_checking

package ec2

// Building without runtime type checking enabled, so all the below just return nil

func validateAdvertiseService_OfParameters(name *string) error {
	return nil
}

