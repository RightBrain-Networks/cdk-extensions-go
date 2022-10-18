package kinesisfirehose

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awskinesisfirehose"
	"github.com/aws/constructs-go/constructs/v10"
)

type MetadataExtractionProcessor interface {
	DeliveryStreamProcessor
	Engine() JsonParsingEngine
	ProcessorType() ProcessorType
	Query() MetaDataExtractionQuery
	AddProcessorParameter(name *string, value *string)
	Bind(_scope constructs.IConstruct) *awskinesisfirehose.CfnDeliveryStream_ProcessorProperty
}

// The jsii proxy struct for MetadataExtractionProcessor
type jsiiProxy_MetadataExtractionProcessor struct {
	jsiiProxy_DeliveryStreamProcessor
}

func (j *jsiiProxy_MetadataExtractionProcessor) Engine() JsonParsingEngine {
	var returns JsonParsingEngine
	_jsii_.Get(
		j,
		"engine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MetadataExtractionProcessor) ProcessorType() ProcessorType {
	var returns ProcessorType
	_jsii_.Get(
		j,
		"processorType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MetadataExtractionProcessor) Query() MetaDataExtractionQuery {
	var returns MetaDataExtractionQuery
	_jsii_.Get(
		j,
		"query",
		&returns,
	)
	return returns
}


func NewMetadataExtractionProcessor(options *MetadataExtractionProcessorOptions) MetadataExtractionProcessor {
	_init_.Initialize()

	if err := validateNewMetadataExtractionProcessorParameters(options); err != nil {
		panic(err)
	}
	j := jsiiProxy_MetadataExtractionProcessor{}

	_jsii_.Create(
		"cdk-extensions.kinesis_firehose.MetadataExtractionProcessor",
		[]interface{}{options},
		&j,
	)

	return &j
}

func NewMetadataExtractionProcessor_Override(m MetadataExtractionProcessor, options *MetadataExtractionProcessorOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.kinesis_firehose.MetadataExtractionProcessor",
		[]interface{}{options},
		m,
	)
}

func (m *jsiiProxy_MetadataExtractionProcessor) AddProcessorParameter(name *string, value *string) {
	if err := m.validateAddProcessorParameterParameters(name, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addProcessorParameter",
		[]interface{}{name, value},
	)
}

func (m *jsiiProxy_MetadataExtractionProcessor) Bind(_scope constructs.IConstruct) *awskinesisfirehose.CfnDeliveryStream_ProcessorProperty {
	if err := m.validateBindParameters(_scope); err != nil {
		panic(err)
	}
	var returns *awskinesisfirehose.CfnDeliveryStream_ProcessorProperty

	_jsii_.Invoke(
		m,
		"bind",
		[]interface{}{_scope},
		&returns,
	)

	return returns
}

