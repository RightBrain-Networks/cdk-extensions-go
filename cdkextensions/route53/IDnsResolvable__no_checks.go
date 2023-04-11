//go:build no_runtime_type_checking

package route53

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IDnsResolvable) validateRegisterDomainParameters(domain Domain) error {
	return nil
}

