package kinesisfirehose

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskinesisfirehose"
	"github.com/aws/constructs-go/constructs/v10"
)

type BufferingConfiguration interface {
	Interval() awscdk.Duration
	SizeInMb() *float64
	Bind(_scope constructs.IConstruct) *awskinesisfirehose.CfnDeliveryStream_BufferingHintsProperty
}

// The jsii proxy struct for BufferingConfiguration
type jsiiProxy_BufferingConfiguration struct {
	_ byte // padding
}

func (j *jsiiProxy_BufferingConfiguration) Interval() awscdk.Duration {
	var returns awscdk.Duration
	_jsii_.Get(
		j,
		"interval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BufferingConfiguration) SizeInMb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sizeInMb",
		&returns,
	)
	return returns
}


func NewBufferingConfiguration(options *BufferingConfigurationOptions) BufferingConfiguration {
	_init_.Initialize()

	if err := validateNewBufferingConfigurationParameters(options); err != nil {
		panic(err)
	}
	j := jsiiProxy_BufferingConfiguration{}

	_jsii_.Create(
		"cdk-extensions.kinesis_firehose.BufferingConfiguration",
		[]interface{}{options},
		&j,
	)

	return &j
}

func NewBufferingConfiguration_Override(b BufferingConfiguration, options *BufferingConfigurationOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.kinesis_firehose.BufferingConfiguration",
		[]interface{}{options},
		b,
	)
}

func (b *jsiiProxy_BufferingConfiguration) Bind(_scope constructs.IConstruct) *awskinesisfirehose.CfnDeliveryStream_BufferingHintsProperty {
	if err := b.validateBindParameters(_scope); err != nil {
		panic(err)
	}
	var returns *awskinesisfirehose.CfnDeliveryStream_BufferingHintsProperty

	_jsii_.Invoke(
		b,
		"bind",
		[]interface{}{_scope},
		&returns,
	)

	return returns
}

