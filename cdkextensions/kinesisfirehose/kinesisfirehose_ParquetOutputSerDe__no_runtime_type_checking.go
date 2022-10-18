//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package kinesisfirehose

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_ParquetOutputSerDe) validateBindParameters(_scope constructs.IConstruct) error {
	return nil
}

func validateParquetOutputSerDe_OrcParameters(options *OrcOutputSerDeOptions) error {
	return nil
}

func validateParquetOutputSerDe_ParquetParameters(options *ParquetOutputSerDeOptions) error {
	return nil
}

func validateNewParquetOutputSerDeParameters(options *ParquetOutputSerDeOptions) error {
	return nil
}

