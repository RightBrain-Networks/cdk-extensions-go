//go:build no_runtime_type_checking

package ec2

// Building without runtime type checking enabled, so all the below just return nil

func validateIpv4CidrAssignment_CustomParameters(options *Ipv4CidrAssignmentCustomOptions) error {
	return nil
}

func validateIpv4CidrAssignment_IpamPoolParameters(options *Ipv4CidrAssignmentIpamPoolOptions) error {
	return nil
}

