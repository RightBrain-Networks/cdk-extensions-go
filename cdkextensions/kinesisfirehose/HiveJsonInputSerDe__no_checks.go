//go:build no_runtime_type_checking

package kinesisfirehose

// Building without runtime type checking enabled, so all the below just return nil

func (h *jsiiProxy_HiveJsonInputSerDe) validateAddTimestampFormatParameters(format *string) error {
	return nil
}

func (h *jsiiProxy_HiveJsonInputSerDe) validateBindParameters(_scope constructs.IConstruct) error {
	return nil
}

func validateHiveJsonInputSerDe_HiveJsonParameters(options *HiveJsonInputSerDeOptions) error {
	return nil
}

func validateHiveJsonInputSerDe_OpenxJsonParameters(options *OpenxJsonInputSerDeOptions) error {
	return nil
}

func validateNewHiveJsonInputSerDeParameters(options *HiveJsonInputSerDeOptions) error {
	return nil
}

