package kinesisfirehose

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awskinesisfirehose"
	"github.com/aws/constructs-go/constructs/v10"
)

type HiveJsonInputSerDe interface {
	InputFormat
	AddTimestampFormat(format *string) HiveJsonInputSerDe
	Bind(_scope constructs.IConstruct) *awskinesisfirehose.CfnDeliveryStream_InputFormatConfigurationProperty
}

// The jsii proxy struct for HiveJsonInputSerDe
type jsiiProxy_HiveJsonInputSerDe struct {
	jsiiProxy_InputFormat
}

func NewHiveJsonInputSerDe(options *HiveJsonInputSerDeOptions) HiveJsonInputSerDe {
	_init_.Initialize()

	if err := validateNewHiveJsonInputSerDeParameters(options); err != nil {
		panic(err)
	}
	j := jsiiProxy_HiveJsonInputSerDe{}

	_jsii_.Create(
		"cdk-extensions.kinesis_firehose.HiveJsonInputSerDe",
		[]interface{}{options},
		&j,
	)

	return &j
}

func NewHiveJsonInputSerDe_Override(h HiveJsonInputSerDe, options *HiveJsonInputSerDeOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.kinesis_firehose.HiveJsonInputSerDe",
		[]interface{}{options},
		h,
	)
}

func HiveJsonInputSerDe_HiveJson(options *HiveJsonInputSerDeOptions) HiveJsonInputSerDe {
	_init_.Initialize()

	if err := validateHiveJsonInputSerDe_HiveJsonParameters(options); err != nil {
		panic(err)
	}
	var returns HiveJsonInputSerDe

	_jsii_.StaticInvoke(
		"cdk-extensions.kinesis_firehose.HiveJsonInputSerDe",
		"hiveJson",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func HiveJsonInputSerDe_OpenxJson(options *OpenxJsonInputSerDeOptions) OpenxJsonInputSerDe {
	_init_.Initialize()

	if err := validateHiveJsonInputSerDe_OpenxJsonParameters(options); err != nil {
		panic(err)
	}
	var returns OpenxJsonInputSerDe

	_jsii_.StaticInvoke(
		"cdk-extensions.kinesis_firehose.HiveJsonInputSerDe",
		"openxJson",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HiveJsonInputSerDe) AddTimestampFormat(format *string) HiveJsonInputSerDe {
	if err := h.validateAddTimestampFormatParameters(format); err != nil {
		panic(err)
	}
	var returns HiveJsonInputSerDe

	_jsii_.Invoke(
		h,
		"addTimestampFormat",
		[]interface{}{format},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HiveJsonInputSerDe) Bind(_scope constructs.IConstruct) *awskinesisfirehose.CfnDeliveryStream_InputFormatConfigurationProperty {
	if err := h.validateBindParameters(_scope); err != nil {
		panic(err)
	}
	var returns *awskinesisfirehose.CfnDeliveryStream_InputFormatConfigurationProperty

	_jsii_.Invoke(
		h,
		"bind",
		[]interface{}{_scope},
		&returns,
	)

	return returns
}

