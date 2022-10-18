//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package kinesisfirehose

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LambdaProcessor) validateAddProcessorParameterParameters(name *string, value *string) error {
	return nil
}

func (l *jsiiProxy_LambdaProcessor) validateBindParameters(_scope constructs.IConstruct) error {
	return nil
}

func validateNewLambdaProcessorParameters(options *LambdaProcessorOptions) error {
	return nil
}

