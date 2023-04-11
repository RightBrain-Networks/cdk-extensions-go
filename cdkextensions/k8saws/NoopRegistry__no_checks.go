//go:build no_runtime_type_checking

package k8saws

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_NoopRegistry) validateBindParameters(_scope constructs.IConstruct) error {
	return nil
}

