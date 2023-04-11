//go:build no_runtime_type_checking

package kinesisfirehose

// Building without runtime type checking enabled, so all the below just return nil

func (o *jsiiProxy_OutputFormat) validateBindParameters(scope constructs.IConstruct) error {
	return nil
}

func validateOutputFormat_OrcParameters(options *OrcOutputSerDeOptions) error {
	return nil
}

func validateOutputFormat_ParquetParameters(options *ParquetOutputSerDeOptions) error {
	return nil
}

