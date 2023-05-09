//go:build no_runtime_type_checking

package ec2

// Building without runtime type checking enabled, so all the below just return nil

func validateIpv6CidrAssignment_CustomParameters(options *Ipv4CidrAssignmentCustomOptions) error {
	return nil
}

func validateIpv6CidrAssignment_IpamPoolParameters(options *Ipv6CidrAssignmentIpamPoolOptions) error {
	return nil
}

