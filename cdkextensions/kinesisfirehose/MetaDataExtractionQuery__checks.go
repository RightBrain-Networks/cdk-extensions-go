//go:build !no_runtime_type_checking

package kinesisfirehose

import (
	"fmt"
)

func validateMetaDataExtractionQuery_JqParameters(fields *map[string]*string) error {
	if fields == nil {
		return fmt.Errorf("parameter fields is required, but nil was provided")
	}

	return nil
}

func validateMetaDataExtractionQuery_OfParameters(query *string) error {
	if query == nil {
		return fmt.Errorf("parameter query is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_MetaDataExtractionQuery) validateSetQueryParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewMetaDataExtractionQueryParameters(query *string) error {
	if query == nil {
		return fmt.Errorf("parameter query is required, but nil was provided")
	}

	return nil
}

