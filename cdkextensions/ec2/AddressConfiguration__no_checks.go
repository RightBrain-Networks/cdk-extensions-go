//go:build no_runtime_type_checking

package ec2

// Building without runtime type checking enabled, so all the below just return nil

func validateAddressConfiguration_Ipv4Parameters(options *Ipv4ConfigurationOptions) error {
	return nil
}

func validateAddressConfiguration_Ipv6Parameters(options *Ipv6ConfigurationOptions) error {
	return nil
}

func validateAddressConfiguration_OfParameters(props *AddressConfigurationProps) error {
	return nil
}

