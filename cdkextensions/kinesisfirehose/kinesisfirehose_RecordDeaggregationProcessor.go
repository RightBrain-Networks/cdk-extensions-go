package kinesisfirehose

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awskinesisfirehose"
	"github.com/aws/constructs-go/constructs/v10"
)

type RecordDeaggregationProcessor interface {
	DeliveryStreamProcessor
	Delimiter() *string
	ProcessorType() ProcessorType
	SubRecordType() SubRecordType
	AddProcessorParameter(name *string, value *string)
	Bind(_scope constructs.IConstruct) *awskinesisfirehose.CfnDeliveryStream_ProcessorProperty
}

// The jsii proxy struct for RecordDeaggregationProcessor
type jsiiProxy_RecordDeaggregationProcessor struct {
	jsiiProxy_DeliveryStreamProcessor
}

func (j *jsiiProxy_RecordDeaggregationProcessor) Delimiter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delimiter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecordDeaggregationProcessor) ProcessorType() ProcessorType {
	var returns ProcessorType
	_jsii_.Get(
		j,
		"processorType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecordDeaggregationProcessor) SubRecordType() SubRecordType {
	var returns SubRecordType
	_jsii_.Get(
		j,
		"subRecordType",
		&returns,
	)
	return returns
}


func NewRecordDeaggregationProcessor(options *RecordDeaggregationProcessorOptions) RecordDeaggregationProcessor {
	_init_.Initialize()

	if err := validateNewRecordDeaggregationProcessorParameters(options); err != nil {
		panic(err)
	}
	j := jsiiProxy_RecordDeaggregationProcessor{}

	_jsii_.Create(
		"cdk-extensions.kinesis_firehose.RecordDeaggregationProcessor",
		[]interface{}{options},
		&j,
	)

	return &j
}

func NewRecordDeaggregationProcessor_Override(r RecordDeaggregationProcessor, options *RecordDeaggregationProcessorOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.kinesis_firehose.RecordDeaggregationProcessor",
		[]interface{}{options},
		r,
	)
}

func RecordDeaggregationProcessor_Delimited(options *DelimitedDeaggregationOptions) RecordDeaggregationProcessor {
	_init_.Initialize()

	if err := validateRecordDeaggregationProcessor_DelimitedParameters(options); err != nil {
		panic(err)
	}
	var returns RecordDeaggregationProcessor

	_jsii_.StaticInvoke(
		"cdk-extensions.kinesis_firehose.RecordDeaggregationProcessor",
		"delimited",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func RecordDeaggregationProcessor_Json() RecordDeaggregationProcessor {
	_init_.Initialize()

	var returns RecordDeaggregationProcessor

	_jsii_.StaticInvoke(
		"cdk-extensions.kinesis_firehose.RecordDeaggregationProcessor",
		"json",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RecordDeaggregationProcessor) AddProcessorParameter(name *string, value *string) {
	if err := r.validateAddProcessorParameterParameters(name, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"addProcessorParameter",
		[]interface{}{name, value},
	)
}

func (r *jsiiProxy_RecordDeaggregationProcessor) Bind(_scope constructs.IConstruct) *awskinesisfirehose.CfnDeliveryStream_ProcessorProperty {
	if err := r.validateBindParameters(_scope); err != nil {
		panic(err)
	}
	var returns *awskinesisfirehose.CfnDeliveryStream_ProcessorProperty

	_jsii_.Invoke(
		r,
		"bind",
		[]interface{}{_scope},
		&returns,
	)

	return returns
}

