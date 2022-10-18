//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package kinesisfirehose

// Building without runtime type checking enabled, so all the below just return nil

func (o *jsiiProxy_OrcOutputSerDe) validateAddBloomFilterColumnParameters(column *string) error {
	return nil
}

func (o *jsiiProxy_OrcOutputSerDe) validateBindParameters(_scope constructs.IConstruct) error {
	return nil
}

func validateOrcOutputSerDe_OrcParameters(options *OrcOutputSerDeOptions) error {
	return nil
}

func validateOrcOutputSerDe_ParquetParameters(options *ParquetOutputSerDeOptions) error {
	return nil
}

func validateNewOrcOutputSerDeParameters(options *OrcOutputSerDeOptions) error {
	return nil
}

