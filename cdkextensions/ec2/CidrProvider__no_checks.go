//go:build no_runtime_type_checking

package ec2

// Building without runtime type checking enabled, so all the below just return nil

func validateCidrProvider_CidrParameters(cidr *string) error {
	return nil
}

func validateCidrProvider_IpamPoolParameters(pool IIpamPool, netmask *float64) error {
	return nil
}

