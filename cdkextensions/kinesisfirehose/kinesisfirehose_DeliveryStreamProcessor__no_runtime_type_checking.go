//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package kinesisfirehose

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DeliveryStreamProcessor) validateAddProcessorParameterParameters(name *string, value *string) error {
	return nil
}

func (d *jsiiProxy_DeliveryStreamProcessor) validateBindParameters(_scope constructs.IConstruct) error {
	return nil
}

func validateNewDeliveryStreamProcessorParameters(options *DeliveryStreamProcessorOptions) error {
	return nil
}

