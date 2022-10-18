//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package kinesisfirehose

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AppendDelimiterProcessor) validateAddProcessorParameterParameters(name *string, value *string) error {
	return nil
}

func (a *jsiiProxy_AppendDelimiterProcessor) validateBindParameters(_scope constructs.IConstruct) error {
	return nil
}

func validateNewAppendDelimiterProcessorParameters(options *AppendDelimiterProcessorOptions) error {
	return nil
}

