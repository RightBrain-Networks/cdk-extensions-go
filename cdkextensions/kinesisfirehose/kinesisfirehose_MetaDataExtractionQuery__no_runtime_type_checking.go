//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package kinesisfirehose

// Building without runtime type checking enabled, so all the below just return nil

func validateMetaDataExtractionQuery_JqParameters(fields *map[string]*string) error {
	return nil
}

func validateMetaDataExtractionQuery_OfParameters(query *string) error {
	return nil
}

func (j *jsiiProxy_MetaDataExtractionQuery) validateSetQueryParameters(val *string) error {
	return nil
}

func validateNewMetaDataExtractionQueryParameters(query *string) error {
	return nil
}

