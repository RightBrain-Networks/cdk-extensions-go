//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package ec2

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FlowLogFormat) validateAddFieldParameters(field FlowLogField) error {
	return nil
}

