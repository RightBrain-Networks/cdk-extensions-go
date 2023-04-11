//go:build !no_runtime_type_checking

package ec2

import (
	"fmt"
)

func validateCidrProvider_CidrParameters(cidr *string) error {
	if cidr == nil {
		return fmt.Errorf("parameter cidr is required, but nil was provided")
	}

	return nil
}

func validateCidrProvider_IpamPoolParameters(pool IIpamPool, netmask *float64) error {
	if pool == nil {
		return fmt.Errorf("parameter pool is required, but nil was provided")
	}

	if netmask == nil {
		return fmt.Errorf("parameter netmask is required, but nil was provided")
	}

	return nil
}

