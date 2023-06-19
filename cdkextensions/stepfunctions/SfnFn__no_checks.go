//go:build no_runtime_type_checking

package stepfunctions

// Building without runtime type checking enabled, so all the below just return nil

func validateSfnFn_ArrayContainsParameters(array *string, lookingFor interface{}) error {
	return nil
}

func validateSfnFn_ArrayGetItemParameters(array *string, index interface{}) error {
	return nil
}

func validateSfnFn_ArrayLengthParameters(array *string) error {
	return nil
}

func validateSfnFn_ArrayPartitionParameters(array *string, size interface{}) error {
	return nil
}

func validateSfnFn_ArrayRangeParameters(start interface{}, stop interface{}, step interface{}) error {
	return nil
}

func validateSfnFn_ArrayUniqueParameters(array *string) error {
	return nil
}

func validateSfnFn_Base64DecodeParameters(value *string) error {
	return nil
}

func validateSfnFn_Base64EncodeParameters(value *string) error {
	return nil
}

func validateSfnFn_FormatParameters(template *string, values *[]*string) error {
	return nil
}

func validateSfnFn_HashParameters(data *string, algorithm *string) error {
	return nil
}

func validateSfnFn_JsonMergeParameters(json1 *string, json2 *string) error {
	return nil
}

func validateSfnFn_JsonToStringParameters(data *string) error {
	return nil
}

func validateSfnFn_MathAddParameters(value *string, step interface{}) error {
	return nil
}

func validateSfnFn_MathRandomParameters(start interface{}, end interface{}) error {
	return nil
}

func validateSfnFn_StringSplitParameters(data *string, splitter *string) error {
	return nil
}

func validateSfnFn_StringToJsonParameters(data *string) error {
	return nil
}

