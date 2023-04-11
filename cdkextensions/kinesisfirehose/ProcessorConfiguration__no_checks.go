//go:build no_runtime_type_checking

package kinesisfirehose

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_ProcessorConfiguration) validateBindParameters(_scope constructs.IConstruct) error {
	return nil
}

func validateNewProcessorConfigurationParameters(options *ProcessorConfigurationOptions) error {
	return nil
}

