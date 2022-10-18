//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package kinesisfirehose

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_MetadataExtractionProcessor) validateAddProcessorParameterParameters(name *string, value *string) error {
	return nil
}

func (m *jsiiProxy_MetadataExtractionProcessor) validateBindParameters(_scope constructs.IConstruct) error {
	return nil
}

func validateNewMetadataExtractionProcessorParameters(options *MetadataExtractionProcessorOptions) error {
	return nil
}

