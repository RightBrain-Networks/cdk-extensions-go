package kinesisfirehose

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awskinesisfirehose"
	"github.com/aws/constructs-go/constructs/v10"
)

type OutputFormat interface {
	Bind(scope constructs.IConstruct) *awskinesisfirehose.CfnDeliveryStream_OutputFormatConfigurationProperty
}

// The jsii proxy struct for OutputFormat
type jsiiProxy_OutputFormat struct {
	_ byte // padding
}

func NewOutputFormat_Override(o OutputFormat) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.kinesis_firehose.OutputFormat",
		nil, // no parameters
		o,
	)
}

func OutputFormat_Orc(options *OrcOutputSerDeOptions) OrcOutputSerDe {
	_init_.Initialize()

	if err := validateOutputFormat_OrcParameters(options); err != nil {
		panic(err)
	}
	var returns OrcOutputSerDe

	_jsii_.StaticInvoke(
		"cdk-extensions.kinesis_firehose.OutputFormat",
		"orc",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func OutputFormat_Parquet(options *ParquetOutputSerDeOptions) ParquetOutputSerDe {
	_init_.Initialize()

	if err := validateOutputFormat_ParquetParameters(options); err != nil {
		panic(err)
	}
	var returns ParquetOutputSerDe

	_jsii_.StaticInvoke(
		"cdk-extensions.kinesis_firehose.OutputFormat",
		"parquet",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OutputFormat) Bind(scope constructs.IConstruct) *awskinesisfirehose.CfnDeliveryStream_OutputFormatConfigurationProperty {
	if err := o.validateBindParameters(scope); err != nil {
		panic(err)
	}
	var returns *awskinesisfirehose.CfnDeliveryStream_OutputFormatConfigurationProperty

	_jsii_.Invoke(
		o,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

