package kinesisfirehose

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awskinesisfirehose"
	"github.com/aws/constructs-go/constructs/v10"
)

type AppendDelimiterProcessor interface {
	DeliveryStreamProcessor
	Delimiter() *string
	ProcessorType() ProcessorType
	AddProcessorParameter(name *string, value *string)
	Bind(_scope constructs.IConstruct) *awskinesisfirehose.CfnDeliveryStream_ProcessorProperty
}

// The jsii proxy struct for AppendDelimiterProcessor
type jsiiProxy_AppendDelimiterProcessor struct {
	jsiiProxy_DeliveryStreamProcessor
}

func (j *jsiiProxy_AppendDelimiterProcessor) Delimiter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delimiter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppendDelimiterProcessor) ProcessorType() ProcessorType {
	var returns ProcessorType
	_jsii_.Get(
		j,
		"processorType",
		&returns,
	)
	return returns
}


func NewAppendDelimiterProcessor(options *AppendDelimiterProcessorOptions) AppendDelimiterProcessor {
	_init_.Initialize()

	if err := validateNewAppendDelimiterProcessorParameters(options); err != nil {
		panic(err)
	}
	j := jsiiProxy_AppendDelimiterProcessor{}

	_jsii_.Create(
		"cdk-extensions.kinesis_firehose.AppendDelimiterProcessor",
		[]interface{}{options},
		&j,
	)

	return &j
}

func NewAppendDelimiterProcessor_Override(a AppendDelimiterProcessor, options *AppendDelimiterProcessorOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.kinesis_firehose.AppendDelimiterProcessor",
		[]interface{}{options},
		a,
	)
}

func (a *jsiiProxy_AppendDelimiterProcessor) AddProcessorParameter(name *string, value *string) {
	if err := a.validateAddProcessorParameterParameters(name, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addProcessorParameter",
		[]interface{}{name, value},
	)
}

func (a *jsiiProxy_AppendDelimiterProcessor) Bind(_scope constructs.IConstruct) *awskinesisfirehose.CfnDeliveryStream_ProcessorProperty {
	if err := a.validateBindParameters(_scope); err != nil {
		panic(err)
	}
	var returns *awskinesisfirehose.CfnDeliveryStream_ProcessorProperty

	_jsii_.Invoke(
		a,
		"bind",
		[]interface{}{_scope},
		&returns,
	)

	return returns
}

