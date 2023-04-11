//go:build no_runtime_type_checking

package kinesisfirehose

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_InputFormat) validateBindParameters(scope constructs.IConstruct) error {
	return nil
}

func validateInputFormat_HiveJsonParameters(options *HiveJsonInputSerDeOptions) error {
	return nil
}

func validateInputFormat_OpenxJsonParameters(options *OpenxJsonInputSerDeOptions) error {
	return nil
}

