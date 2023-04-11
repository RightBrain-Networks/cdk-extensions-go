//go:build no_runtime_type_checking

package kinesisfirehose

// Building without runtime type checking enabled, so all the below just return nil

func (b *jsiiProxy_BufferingConfiguration) validateBindParameters(_scope constructs.IConstruct) error {
	return nil
}

func validateNewBufferingConfigurationParameters(options *BufferingConfigurationOptions) error {
	return nil
}

