package kinesisfirehose

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"
)

type ProcessorType interface {
	// The name of the processor to apply to the delivery stream.
	Name() *string
}

// The jsii proxy struct for ProcessorType
type jsiiProxy_ProcessorType struct {
	_ byte // padding
}

func (j *jsiiProxy_ProcessorType) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}


func ProcessorType_Of(name *string) ProcessorType {
	_init_.Initialize()

	if err := validateProcessorType_OfParameters(name); err != nil {
		panic(err)
	}
	var returns ProcessorType

	_jsii_.StaticInvoke(
		"cdk-extensions.kinesis_firehose.ProcessorType",
		"of",
		[]interface{}{name},
		&returns,
	)

	return returns
}

func ProcessorType_APPEND_DELIMITER_TO_RECORD() ProcessorType {
	_init_.Initialize()
	var returns ProcessorType
	_jsii_.StaticGet(
		"cdk-extensions.kinesis_firehose.ProcessorType",
		"APPEND_DELIMITER_TO_RECORD",
		&returns,
	)
	return returns
}

func ProcessorType_LAMBDA() ProcessorType {
	_init_.Initialize()
	var returns ProcessorType
	_jsii_.StaticGet(
		"cdk-extensions.kinesis_firehose.ProcessorType",
		"LAMBDA",
		&returns,
	)
	return returns
}

func ProcessorType_METADATA_EXTRACTION() ProcessorType {
	_init_.Initialize()
	var returns ProcessorType
	_jsii_.StaticGet(
		"cdk-extensions.kinesis_firehose.ProcessorType",
		"METADATA_EXTRACTION",
		&returns,
	)
	return returns
}

func ProcessorType_RECORD_DEAGGREGATION() ProcessorType {
	_init_.Initialize()
	var returns ProcessorType
	_jsii_.StaticGet(
		"cdk-extensions.kinesis_firehose.ProcessorType",
		"RECORD_DEAGGREGATION",
		&returns,
	)
	return returns
}

