//go:build !no_runtime_type_checking

package kinesisfirehose

import (
	"fmt"
)

func (j *jsiiProxy_JsonQuery) validateAddFieldParameters(name *string, query *string) error {
	if name == nil {
		return fmt.Errorf("parameter name is required, but nil was provided")
	}

	if query == nil {
		return fmt.Errorf("parameter query is required, but nil was provided")
	}

	return nil
}

func validateJsonQuery_JqParameters(fields *map[string]*string) error {
	if fields == nil {
		return fmt.Errorf("parameter fields is required, but nil was provided")
	}

	return nil
}

func validateJsonQuery_OfParameters(query *string) error {
	if query == nil {
		return fmt.Errorf("parameter query is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_JsonQuery) validateSetQueryParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

