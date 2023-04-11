package kinesisfirehose

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awskinesisfirehose"
	"github.com/aws/constructs-go/constructs/v10"
)

type CustomProcessor interface {
	DeliveryStreamProcessor
	ProcessorType() ProcessorType
	AddParameter(name *string, value *string)
	AddProcessorParameter(name *string, value *string)
	Bind(_scope constructs.IConstruct) *awskinesisfirehose.CfnDeliveryStream_ProcessorProperty
}

// The jsii proxy struct for CustomProcessor
type jsiiProxy_CustomProcessor struct {
	jsiiProxy_DeliveryStreamProcessor
}

func (j *jsiiProxy_CustomProcessor) ProcessorType() ProcessorType {
	var returns ProcessorType
	_jsii_.Get(
		j,
		"processorType",
		&returns,
	)
	return returns
}


func NewCustomProcessor(options *CustomProcessorOptions) CustomProcessor {
	_init_.Initialize()

	if err := validateNewCustomProcessorParameters(options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CustomProcessor{}

	_jsii_.Create(
		"cdk-extensions.kinesis_firehose.CustomProcessor",
		[]interface{}{options},
		&j,
	)

	return &j
}

func NewCustomProcessor_Override(c CustomProcessor, options *CustomProcessorOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.kinesis_firehose.CustomProcessor",
		[]interface{}{options},
		c,
	)
}

func (c *jsiiProxy_CustomProcessor) AddParameter(name *string, value *string) {
	if err := c.validateAddParameterParameters(name, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addParameter",
		[]interface{}{name, value},
	)
}

func (c *jsiiProxy_CustomProcessor) AddProcessorParameter(name *string, value *string) {
	if err := c.validateAddProcessorParameterParameters(name, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addProcessorParameter",
		[]interface{}{name, value},
	)
}

func (c *jsiiProxy_CustomProcessor) Bind(_scope constructs.IConstruct) *awskinesisfirehose.CfnDeliveryStream_ProcessorProperty {
	if err := c.validateBindParameters(_scope); err != nil {
		panic(err)
	}
	var returns *awskinesisfirehose.CfnDeliveryStream_ProcessorProperty

	_jsii_.Invoke(
		c,
		"bind",
		[]interface{}{_scope},
		&returns,
	)

	return returns
}

