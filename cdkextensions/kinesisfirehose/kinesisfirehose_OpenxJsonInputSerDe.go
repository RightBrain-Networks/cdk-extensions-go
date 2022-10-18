package kinesisfirehose

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awskinesisfirehose"
	"github.com/aws/constructs-go/constructs/v10"
)

type OpenxJsonInputSerDe interface {
	InputFormat
	CaseInsensitive() *bool
	ConvertDotsToUnderscores() *bool
	AddColumnKeyMapping(columnName *string, jsonKey *string) OpenxJsonInputSerDe
	Bind(_scope constructs.IConstruct) *awskinesisfirehose.CfnDeliveryStream_InputFormatConfigurationProperty
}

// The jsii proxy struct for OpenxJsonInputSerDe
type jsiiProxy_OpenxJsonInputSerDe struct {
	jsiiProxy_InputFormat
}

func (j *jsiiProxy_OpenxJsonInputSerDe) CaseInsensitive() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"caseInsensitive",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpenxJsonInputSerDe) ConvertDotsToUnderscores() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"convertDotsToUnderscores",
		&returns,
	)
	return returns
}


func NewOpenxJsonInputSerDe(options *OpenxJsonInputSerDeOptions) OpenxJsonInputSerDe {
	_init_.Initialize()

	if err := validateNewOpenxJsonInputSerDeParameters(options); err != nil {
		panic(err)
	}
	j := jsiiProxy_OpenxJsonInputSerDe{}

	_jsii_.Create(
		"cdk-extensions.kinesis_firehose.OpenxJsonInputSerDe",
		[]interface{}{options},
		&j,
	)

	return &j
}

func NewOpenxJsonInputSerDe_Override(o OpenxJsonInputSerDe, options *OpenxJsonInputSerDeOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.kinesis_firehose.OpenxJsonInputSerDe",
		[]interface{}{options},
		o,
	)
}

func OpenxJsonInputSerDe_HiveJson(options *HiveJsonInputSerDeOptions) HiveJsonInputSerDe {
	_init_.Initialize()

	if err := validateOpenxJsonInputSerDe_HiveJsonParameters(options); err != nil {
		panic(err)
	}
	var returns HiveJsonInputSerDe

	_jsii_.StaticInvoke(
		"cdk-extensions.kinesis_firehose.OpenxJsonInputSerDe",
		"hiveJson",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func OpenxJsonInputSerDe_OpenxJson(options *OpenxJsonInputSerDeOptions) OpenxJsonInputSerDe {
	_init_.Initialize()

	if err := validateOpenxJsonInputSerDe_OpenxJsonParameters(options); err != nil {
		panic(err)
	}
	var returns OpenxJsonInputSerDe

	_jsii_.StaticInvoke(
		"cdk-extensions.kinesis_firehose.OpenxJsonInputSerDe",
		"openxJson",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpenxJsonInputSerDe) AddColumnKeyMapping(columnName *string, jsonKey *string) OpenxJsonInputSerDe {
	if err := o.validateAddColumnKeyMappingParameters(columnName, jsonKey); err != nil {
		panic(err)
	}
	var returns OpenxJsonInputSerDe

	_jsii_.Invoke(
		o,
		"addColumnKeyMapping",
		[]interface{}{columnName, jsonKey},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpenxJsonInputSerDe) Bind(_scope constructs.IConstruct) *awskinesisfirehose.CfnDeliveryStream_InputFormatConfigurationProperty {
	if err := o.validateBindParameters(_scope); err != nil {
		panic(err)
	}
	var returns *awskinesisfirehose.CfnDeliveryStream_InputFormatConfigurationProperty

	_jsii_.Invoke(
		o,
		"bind",
		[]interface{}{_scope},
		&returns,
	)

	return returns
}

