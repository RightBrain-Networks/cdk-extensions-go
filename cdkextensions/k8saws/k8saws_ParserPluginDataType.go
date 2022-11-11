package k8saws

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"
)

// Represents the various types of data that can be mapped in Fluent Bit using a parser plugin.
type ParserPluginDataType interface {
	// The name of the data type.
	Name() *string
}

// The jsii proxy struct for ParserPluginDataType
type jsiiProxy_ParserPluginDataType struct {
	_ byte // padding
}

func (j *jsiiProxy_ParserPluginDataType) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}


// An escape hatch method that allow specifying arbitrary custom data types.
//
// Returns: An object representing the data type.
func ParserPluginDataType_Of(name *string) ParserPluginDataType {
	_init_.Initialize()

	if err := validateParserPluginDataType_OfParameters(name); err != nil {
		panic(err)
	}
	var returns ParserPluginDataType

	_jsii_.StaticInvoke(
		"cdk-extensions.k8s_aws.ParserPluginDataType",
		"of",
		[]interface{}{name},
		&returns,
	)

	return returns
}

func ParserPluginDataType_BOOL() ParserPluginDataType {
	_init_.Initialize()
	var returns ParserPluginDataType
	_jsii_.StaticGet(
		"cdk-extensions.k8s_aws.ParserPluginDataType",
		"BOOL",
		&returns,
	)
	return returns
}

func ParserPluginDataType_FLOAT() ParserPluginDataType {
	_init_.Initialize()
	var returns ParserPluginDataType
	_jsii_.StaticGet(
		"cdk-extensions.k8s_aws.ParserPluginDataType",
		"FLOAT",
		&returns,
	)
	return returns
}

func ParserPluginDataType_HEX() ParserPluginDataType {
	_init_.Initialize()
	var returns ParserPluginDataType
	_jsii_.StaticGet(
		"cdk-extensions.k8s_aws.ParserPluginDataType",
		"HEX",
		&returns,
	)
	return returns
}

func ParserPluginDataType_INTEGER() ParserPluginDataType {
	_init_.Initialize()
	var returns ParserPluginDataType
	_jsii_.StaticGet(
		"cdk-extensions.k8s_aws.ParserPluginDataType",
		"INTEGER",
		&returns,
	)
	return returns
}

func ParserPluginDataType_LOGFMT() ParserPluginDataType {
	_init_.Initialize()
	var returns ParserPluginDataType
	_jsii_.StaticGet(
		"cdk-extensions.k8s_aws.ParserPluginDataType",
		"LOGFMT",
		&returns,
	)
	return returns
}

func ParserPluginDataType_LTSV() ParserPluginDataType {
	_init_.Initialize()
	var returns ParserPluginDataType
	_jsii_.StaticGet(
		"cdk-extensions.k8s_aws.ParserPluginDataType",
		"LTSV",
		&returns,
	)
	return returns
}

func ParserPluginDataType_REGEX() ParserPluginDataType {
	_init_.Initialize()
	var returns ParserPluginDataType
	_jsii_.StaticGet(
		"cdk-extensions.k8s_aws.ParserPluginDataType",
		"REGEX",
		&returns,
	)
	return returns
}

func ParserPluginDataType_STRING() ParserPluginDataType {
	_init_.Initialize()
	var returns ParserPluginDataType
	_jsii_.StaticGet(
		"cdk-extensions.k8s_aws.ParserPluginDataType",
		"STRING",
		&returns,
	)
	return returns
}

