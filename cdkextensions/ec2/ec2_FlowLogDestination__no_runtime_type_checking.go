//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package ec2

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FlowLogDestination) validateBindParameters(scope constructs.IConstruct) error {
	return nil
}

func validateFlowLogDestination_ToS3Parameters(options *FlowLogS3Options) error {
	return nil
}

