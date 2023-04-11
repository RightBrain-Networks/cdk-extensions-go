//go:build !no_runtime_type_checking

package route53

import (
	"fmt"
)

func (i *jsiiProxy_IDnsResolvable) validateRegisterDomainParameters(domain Domain) error {
	if domain == nil {
		return fmt.Errorf("parameter domain is required, but nil was provided")
	}

	return nil
}

