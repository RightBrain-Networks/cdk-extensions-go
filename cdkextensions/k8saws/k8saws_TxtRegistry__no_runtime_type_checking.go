//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package k8saws

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TxtRegistry) validateBindParameters(scope constructs.IConstruct) error {
	return nil
}

func validateNewTxtRegistryParameters(options *TxtRegistryOptions) error {
	return nil
}

