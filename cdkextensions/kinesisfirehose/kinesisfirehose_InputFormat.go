package kinesisfirehose

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awskinesisfirehose"
	"github.com/aws/constructs-go/constructs/v10"
)

type InputFormat interface {
	Bind(scope constructs.IConstruct) *awskinesisfirehose.CfnDeliveryStream_InputFormatConfigurationProperty
}

// The jsii proxy struct for InputFormat
type jsiiProxy_InputFormat struct {
	_ byte // padding
}

func NewInputFormat_Override(i InputFormat) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.kinesis_firehose.InputFormat",
		nil, // no parameters
		i,
	)
}

func InputFormat_HiveJson(options *HiveJsonInputSerDeOptions) HiveJsonInputSerDe {
	_init_.Initialize()

	if err := validateInputFormat_HiveJsonParameters(options); err != nil {
		panic(err)
	}
	var returns HiveJsonInputSerDe

	_jsii_.StaticInvoke(
		"cdk-extensions.kinesis_firehose.InputFormat",
		"hiveJson",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func InputFormat_OpenxJson(options *OpenxJsonInputSerDeOptions) OpenxJsonInputSerDe {
	_init_.Initialize()

	if err := validateInputFormat_OpenxJsonParameters(options); err != nil {
		panic(err)
	}
	var returns OpenxJsonInputSerDe

	_jsii_.StaticInvoke(
		"cdk-extensions.kinesis_firehose.InputFormat",
		"openxJson",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InputFormat) Bind(scope constructs.IConstruct) *awskinesisfirehose.CfnDeliveryStream_InputFormatConfigurationProperty {
	if err := i.validateBindParameters(scope); err != nil {
		panic(err)
	}
	var returns *awskinesisfirehose.CfnDeliveryStream_InputFormatConfigurationProperty

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

