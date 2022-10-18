//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package kinesisfirehose

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CustomProcessor) validateAddParameterParameters(name *string, value *string) error {
	return nil
}

func (c *jsiiProxy_CustomProcessor) validateAddProcessorParameterParameters(name *string, value *string) error {
	return nil
}

func (c *jsiiProxy_CustomProcessor) validateBindParameters(_scope constructs.IConstruct) error {
	return nil
}

func validateNewCustomProcessorParameters(options *CustomProcessorOptions) error {
	return nil
}

