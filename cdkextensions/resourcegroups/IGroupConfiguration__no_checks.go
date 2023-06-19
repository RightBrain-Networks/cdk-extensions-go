//go:build no_runtime_type_checking

package resourcegroups

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IGroupConfiguration) validateBindParameters(scope constructs.IConstruct) error {
	return nil
}

