//go:build !no_runtime_type_checking

package ram

import (
	"fmt"
)

func (i *jsiiProxy_IResourceShare) validateAddPrincipalParameters(principal ISharedPrincipal) error {
	if principal == nil {
		return fmt.Errorf("parameter principal is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_IResourceShare) validateAddResourceParameters(resource ISharable) error {
	if resource == nil {
		return fmt.Errorf("parameter resource is required, but nil was provided")
	}

	return nil
}

