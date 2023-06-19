//go:build !no_runtime_type_checking

package stepfunctions

import (
	"fmt"
)

func validateSfnFn_ArrayContainsParameters(array *string, lookingFor interface{}) error {
	if array == nil {
		return fmt.Errorf("parameter array is required, but nil was provided")
	}

	if lookingFor == nil {
		return fmt.Errorf("parameter lookingFor is required, but nil was provided")
	}

	return nil
}

func validateSfnFn_ArrayGetItemParameters(array *string, index interface{}) error {
	if array == nil {
		return fmt.Errorf("parameter array is required, but nil was provided")
	}

	if index == nil {
		return fmt.Errorf("parameter index is required, but nil was provided")
	}

	return nil
}

func validateSfnFn_ArrayLengthParameters(array *string) error {
	if array == nil {
		return fmt.Errorf("parameter array is required, but nil was provided")
	}

	return nil
}

func validateSfnFn_ArrayPartitionParameters(array *string, size interface{}) error {
	if array == nil {
		return fmt.Errorf("parameter array is required, but nil was provided")
	}

	if size == nil {
		return fmt.Errorf("parameter size is required, but nil was provided")
	}

	return nil
}

func validateSfnFn_ArrayRangeParameters(start interface{}, stop interface{}, step interface{}) error {
	if start == nil {
		return fmt.Errorf("parameter start is required, but nil was provided")
	}

	if stop == nil {
		return fmt.Errorf("parameter stop is required, but nil was provided")
	}

	if step == nil {
		return fmt.Errorf("parameter step is required, but nil was provided")
	}

	return nil
}

func validateSfnFn_ArrayUniqueParameters(array *string) error {
	if array == nil {
		return fmt.Errorf("parameter array is required, but nil was provided")
	}

	return nil
}

func validateSfnFn_Base64DecodeParameters(value *string) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateSfnFn_Base64EncodeParameters(value *string) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateSfnFn_FormatParameters(template *string, values *[]*string) error {
	if template == nil {
		return fmt.Errorf("parameter template is required, but nil was provided")
	}

	if values == nil {
		return fmt.Errorf("parameter values is required, but nil was provided")
	}

	return nil
}

func validateSfnFn_HashParameters(data *string, algorithm *string) error {
	if data == nil {
		return fmt.Errorf("parameter data is required, but nil was provided")
	}

	if algorithm == nil {
		return fmt.Errorf("parameter algorithm is required, but nil was provided")
	}

	return nil
}

func validateSfnFn_JsonMergeParameters(json1 *string, json2 *string) error {
	if json1 == nil {
		return fmt.Errorf("parameter json1 is required, but nil was provided")
	}

	if json2 == nil {
		return fmt.Errorf("parameter json2 is required, but nil was provided")
	}

	return nil
}

func validateSfnFn_JsonToStringParameters(data *string) error {
	if data == nil {
		return fmt.Errorf("parameter data is required, but nil was provided")
	}

	return nil
}

func validateSfnFn_MathAddParameters(value *string, step interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	if step == nil {
		return fmt.Errorf("parameter step is required, but nil was provided")
	}

	return nil
}

func validateSfnFn_MathRandomParameters(start interface{}, end interface{}) error {
	if start == nil {
		return fmt.Errorf("parameter start is required, but nil was provided")
	}

	if end == nil {
		return fmt.Errorf("parameter end is required, but nil was provided")
	}

	return nil
}

func validateSfnFn_StringSplitParameters(data *string, splitter *string) error {
	if data == nil {
		return fmt.Errorf("parameter data is required, but nil was provided")
	}

	if splitter == nil {
		return fmt.Errorf("parameter splitter is required, but nil was provided")
	}

	return nil
}

func validateSfnFn_StringToJsonParameters(data *string) error {
	if data == nil {
		return fmt.Errorf("parameter data is required, but nil was provided")
	}

	return nil
}

