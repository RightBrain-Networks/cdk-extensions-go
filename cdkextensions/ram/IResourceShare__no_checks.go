//go:build no_runtime_type_checking

package ram

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IResourceShare) validateAddPrincipalParameters(principal ISharedPrincipal) error {
	return nil
}

func (i *jsiiProxy_IResourceShare) validateAddResourceParameters(resource ISharable) error {
	return nil
}

