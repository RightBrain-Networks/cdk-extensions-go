//go:build no_runtime_type_checking

package k8saws

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SsmParameterReference) validateAddFieldMappingParameters(field *SecretFieldReference) error {
	return nil
}

func (s *jsiiProxy_SsmParameterReference) validateBindParameters(_scope constructs.IConstruct) error {
	return nil
}

func validateNewSsmParameterReferenceParameters(parameter awsssm.IParameter, options *SsmParameterReferenceOptions) error {
	return nil
}

