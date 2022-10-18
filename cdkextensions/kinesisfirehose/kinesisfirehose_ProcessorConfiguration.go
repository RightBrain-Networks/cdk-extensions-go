package kinesisfirehose

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/constructs-go/constructs/v10"
)

type ProcessorConfiguration interface {
	Enabled() *bool
	Processors() *[]DeliveryStreamProcessor
	Bind(_scope constructs.IConstruct) *ProcessorConfigurationResult
}

// The jsii proxy struct for ProcessorConfiguration
type jsiiProxy_ProcessorConfiguration struct {
	_ byte // padding
}

func (j *jsiiProxy_ProcessorConfiguration) Enabled() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcessorConfiguration) Processors() *[]DeliveryStreamProcessor {
	var returns *[]DeliveryStreamProcessor
	_jsii_.Get(
		j,
		"processors",
		&returns,
	)
	return returns
}


func NewProcessorConfiguration(options *ProcessorConfigurationOptions) ProcessorConfiguration {
	_init_.Initialize()

	if err := validateNewProcessorConfigurationParameters(options); err != nil {
		panic(err)
	}
	j := jsiiProxy_ProcessorConfiguration{}

	_jsii_.Create(
		"cdk-extensions.kinesis_firehose.ProcessorConfiguration",
		[]interface{}{options},
		&j,
	)

	return &j
}

func NewProcessorConfiguration_Override(p ProcessorConfiguration, options *ProcessorConfigurationOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.kinesis_firehose.ProcessorConfiguration",
		[]interface{}{options},
		p,
	)
}

func (p *jsiiProxy_ProcessorConfiguration) Bind(_scope constructs.IConstruct) *ProcessorConfigurationResult {
	if err := p.validateBindParameters(_scope); err != nil {
		panic(err)
	}
	var returns *ProcessorConfigurationResult

	_jsii_.Invoke(
		p,
		"bind",
		[]interface{}{_scope},
		&returns,
	)

	return returns
}

