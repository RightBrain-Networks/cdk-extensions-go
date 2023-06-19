package stepfunctions

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"
)

type SfnFn interface {
}

// The jsii proxy struct for SfnFn
type jsiiProxy_SfnFn struct {
	_ byte // padding
}

func NewSfnFn() SfnFn {
	_init_.Initialize()

	j := jsiiProxy_SfnFn{}

	_jsii_.Create(
		"cdk-extensions.stepfunctions.SfnFn",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewSfnFn_Override(s SfnFn) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.stepfunctions.SfnFn",
		nil, // no parameters
		s,
	)
}

func SfnFn_Array(values ...interface{}) *string {
	_init_.Initialize()

	args := []interface{}{}
	for _, a := range values {
		args = append(args, a)
	}

	var returns *string

	_jsii_.StaticInvoke(
		"cdk-extensions.stepfunctions.SfnFn",
		"array",
		args,
		&returns,
	)

	return returns
}

func SfnFn_ArrayContains(array *string, lookingFor interface{}) *string {
	_init_.Initialize()

	if err := validateSfnFn_ArrayContainsParameters(array, lookingFor); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"cdk-extensions.stepfunctions.SfnFn",
		"arrayContains",
		[]interface{}{array, lookingFor},
		&returns,
	)

	return returns
}

func SfnFn_ArrayGetItem(array *string, index interface{}) *string {
	_init_.Initialize()

	if err := validateSfnFn_ArrayGetItemParameters(array, index); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"cdk-extensions.stepfunctions.SfnFn",
		"arrayGetItem",
		[]interface{}{array, index},
		&returns,
	)

	return returns
}

func SfnFn_ArrayLength(array *string) *string {
	_init_.Initialize()

	if err := validateSfnFn_ArrayLengthParameters(array); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"cdk-extensions.stepfunctions.SfnFn",
		"arrayLength",
		[]interface{}{array},
		&returns,
	)

	return returns
}

func SfnFn_ArrayPartition(array *string, size interface{}) *string {
	_init_.Initialize()

	if err := validateSfnFn_ArrayPartitionParameters(array, size); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"cdk-extensions.stepfunctions.SfnFn",
		"arrayPartition",
		[]interface{}{array, size},
		&returns,
	)

	return returns
}

func SfnFn_ArrayRange(start interface{}, stop interface{}, step interface{}) *string {
	_init_.Initialize()

	if err := validateSfnFn_ArrayRangeParameters(start, stop, step); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"cdk-extensions.stepfunctions.SfnFn",
		"arrayRange",
		[]interface{}{start, stop, step},
		&returns,
	)

	return returns
}

func SfnFn_ArrayUnique(array *string) *string {
	_init_.Initialize()

	if err := validateSfnFn_ArrayUniqueParameters(array); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"cdk-extensions.stepfunctions.SfnFn",
		"arrayUnique",
		[]interface{}{array},
		&returns,
	)

	return returns
}

func SfnFn_Base64Decode(value *string) *string {
	_init_.Initialize()

	if err := validateSfnFn_Base64DecodeParameters(value); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"cdk-extensions.stepfunctions.SfnFn",
		"base64Decode",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func SfnFn_Base64Encode(value *string) *string {
	_init_.Initialize()

	if err := validateSfnFn_Base64EncodeParameters(value); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"cdk-extensions.stepfunctions.SfnFn",
		"base64Encode",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func SfnFn_Format(template *string, values *[]*string) *string {
	_init_.Initialize()

	if err := validateSfnFn_FormatParameters(template, values); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"cdk-extensions.stepfunctions.SfnFn",
		"format",
		[]interface{}{template, values},
		&returns,
	)

	return returns
}

func SfnFn_Hash(data *string, algorithm *string) *string {
	_init_.Initialize()

	if err := validateSfnFn_HashParameters(data, algorithm); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"cdk-extensions.stepfunctions.SfnFn",
		"hash",
		[]interface{}{data, algorithm},
		&returns,
	)

	return returns
}

func SfnFn_JsonMerge(json1 *string, json2 *string) *string {
	_init_.Initialize()

	if err := validateSfnFn_JsonMergeParameters(json1, json2); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"cdk-extensions.stepfunctions.SfnFn",
		"jsonMerge",
		[]interface{}{json1, json2},
		&returns,
	)

	return returns
}

func SfnFn_JsonToString(data *string) *string {
	_init_.Initialize()

	if err := validateSfnFn_JsonToStringParameters(data); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"cdk-extensions.stepfunctions.SfnFn",
		"jsonToString",
		[]interface{}{data},
		&returns,
	)

	return returns
}

func SfnFn_MathAdd(value *string, step interface{}) *string {
	_init_.Initialize()

	if err := validateSfnFn_MathAddParameters(value, step); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"cdk-extensions.stepfunctions.SfnFn",
		"mathAdd",
		[]interface{}{value, step},
		&returns,
	)

	return returns
}

func SfnFn_MathRandom(start interface{}, end interface{}) *string {
	_init_.Initialize()

	if err := validateSfnFn_MathRandomParameters(start, end); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"cdk-extensions.stepfunctions.SfnFn",
		"mathRandom",
		[]interface{}{start, end},
		&returns,
	)

	return returns
}

func SfnFn_StringSplit(data *string, splitter *string) *string {
	_init_.Initialize()

	if err := validateSfnFn_StringSplitParameters(data, splitter); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"cdk-extensions.stepfunctions.SfnFn",
		"stringSplit",
		[]interface{}{data, splitter},
		&returns,
	)

	return returns
}

func SfnFn_StringToJson(data *string) *string {
	_init_.Initialize()

	if err := validateSfnFn_StringToJsonParameters(data); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"cdk-extensions.stepfunctions.SfnFn",
		"stringToJson",
		[]interface{}{data},
		&returns,
	)

	return returns
}

func SfnFn_Uuid() *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"cdk-extensions.stepfunctions.SfnFn",
		"uuid",
		nil, // no parameters
		&returns,
	)

	return returns
}

