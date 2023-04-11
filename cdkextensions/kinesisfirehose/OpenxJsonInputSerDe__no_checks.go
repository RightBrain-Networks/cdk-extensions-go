//go:build no_runtime_type_checking

package kinesisfirehose

// Building without runtime type checking enabled, so all the below just return nil

func (o *jsiiProxy_OpenxJsonInputSerDe) validateAddColumnKeyMappingParameters(columnName *string, jsonKey *string) error {
	return nil
}

func (o *jsiiProxy_OpenxJsonInputSerDe) validateBindParameters(_scope constructs.IConstruct) error {
	return nil
}

func validateOpenxJsonInputSerDe_HiveJsonParameters(options *HiveJsonInputSerDeOptions) error {
	return nil
}

func validateOpenxJsonInputSerDe_OpenxJsonParameters(options *OpenxJsonInputSerDeOptions) error {
	return nil
}

func validateNewOpenxJsonInputSerDeParameters(options *OpenxJsonInputSerDeOptions) error {
	return nil
}

