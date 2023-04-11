//go:build no_runtime_type_checking

package k8saws

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AwsServiceDiscoveryRegistry) validateBindParameters(scope constructs.IConstruct) error {
	return nil
}

