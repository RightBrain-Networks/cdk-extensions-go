package kinesisfirehose

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awskinesisfirehose"
	"github.com/aws/constructs-go/constructs/v10"
)

type DeliveryStreamProcessor interface {
	ProcessorType() ProcessorType
	AddProcessorParameter(name *string, value *string)
	Bind(_scope constructs.IConstruct) *awskinesisfirehose.CfnDeliveryStream_ProcessorProperty
}

// The jsii proxy struct for DeliveryStreamProcessor
type jsiiProxy_DeliveryStreamProcessor struct {
	_ byte // padding
}

func (j *jsiiProxy_DeliveryStreamProcessor) ProcessorType() ProcessorType {
	var returns ProcessorType
	_jsii_.Get(
		j,
		"processorType",
		&returns,
	)
	return returns
}


func NewDeliveryStreamProcessor(options *DeliveryStreamProcessorOptions) DeliveryStreamProcessor {
	_init_.Initialize()

	if err := validateNewDeliveryStreamProcessorParameters(options); err != nil {
		panic(err)
	}
	j := jsiiProxy_DeliveryStreamProcessor{}

	_jsii_.Create(
		"cdk-extensions.kinesis_firehose.DeliveryStreamProcessor",
		[]interface{}{options},
		&j,
	)

	return &j
}

func NewDeliveryStreamProcessor_Override(d DeliveryStreamProcessor, options *DeliveryStreamProcessorOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.kinesis_firehose.DeliveryStreamProcessor",
		[]interface{}{options},
		d,
	)
}

func (d *jsiiProxy_DeliveryStreamProcessor) AddProcessorParameter(name *string, value *string) {
	if err := d.validateAddProcessorParameterParameters(name, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addProcessorParameter",
		[]interface{}{name, value},
	)
}

func (d *jsiiProxy_DeliveryStreamProcessor) Bind(_scope constructs.IConstruct) *awskinesisfirehose.CfnDeliveryStream_ProcessorProperty {
	if err := d.validateBindParameters(_scope); err != nil {
		panic(err)
	}
	var returns *awskinesisfirehose.CfnDeliveryStream_ProcessorProperty

	_jsii_.Invoke(
		d,
		"bind",
		[]interface{}{_scope},
		&returns,
	)

	return returns
}

