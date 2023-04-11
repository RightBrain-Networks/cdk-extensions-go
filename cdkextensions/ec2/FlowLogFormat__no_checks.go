//go:build no_runtime_type_checking

package ec2

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FlowLogFormat) validateAddFieldParameters(field FlowLogField) error {
	return nil
}

func validateFlowLogFormat_FromTemplateParameters(template *string) error {
	return nil
}

