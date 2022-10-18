//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package kinesisfirehose

// Building without runtime type checking enabled, so all the below just return nil

func (j *jsiiProxy_JsonQuery) validateAddFieldParameters(name *string, query *string) error {
	return nil
}

func validateJsonQuery_JqParameters(fields *map[string]*string) error {
	return nil
}

func validateJsonQuery_OfParameters(query *string) error {
	return nil
}

func (j *jsiiProxy_JsonQuery) validateSetQueryParameters(val *string) error {
	return nil
}

