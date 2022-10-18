//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package kinesisfirehose

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RecordDeaggregationProcessor) validateAddProcessorParameterParameters(name *string, value *string) error {
	return nil
}

func (r *jsiiProxy_RecordDeaggregationProcessor) validateBindParameters(_scope constructs.IConstruct) error {
	return nil
}

func validateRecordDeaggregationProcessor_DelimitedParameters(options *DelimitedDeaggregationOptions) error {
	return nil
}

func validateNewRecordDeaggregationProcessorParameters(options *RecordDeaggregationProcessorOptions) error {
	return nil
}

