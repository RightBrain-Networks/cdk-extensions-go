package kinesisfirehose

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awskinesisfirehose"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/constructs-go/constructs/v10"
)

type LambdaProcessor interface {
	DeliveryStreamProcessor
	LambdaFunction() awslambda.IFunction
	ProcessorType() ProcessorType
	AddProcessorParameter(name *string, value *string)
	Bind(_scope constructs.IConstruct) *awskinesisfirehose.CfnDeliveryStream_ProcessorProperty
}

// The jsii proxy struct for LambdaProcessor
type jsiiProxy_LambdaProcessor struct {
	jsiiProxy_DeliveryStreamProcessor
}

func (j *jsiiProxy_LambdaProcessor) LambdaFunction() awslambda.IFunction {
	var returns awslambda.IFunction
	_jsii_.Get(
		j,
		"lambdaFunction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaProcessor) ProcessorType() ProcessorType {
	var returns ProcessorType
	_jsii_.Get(
		j,
		"processorType",
		&returns,
	)
	return returns
}


func NewLambdaProcessor(options *LambdaProcessorOptions) LambdaProcessor {
	_init_.Initialize()

	if err := validateNewLambdaProcessorParameters(options); err != nil {
		panic(err)
	}
	j := jsiiProxy_LambdaProcessor{}

	_jsii_.Create(
		"cdk-extensions.kinesis_firehose.LambdaProcessor",
		[]interface{}{options},
		&j,
	)

	return &j
}

func NewLambdaProcessor_Override(l LambdaProcessor, options *LambdaProcessorOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.kinesis_firehose.LambdaProcessor",
		[]interface{}{options},
		l,
	)
}

func (l *jsiiProxy_LambdaProcessor) AddProcessorParameter(name *string, value *string) {
	if err := l.validateAddProcessorParameterParameters(name, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"addProcessorParameter",
		[]interface{}{name, value},
	)
}

func (l *jsiiProxy_LambdaProcessor) Bind(_scope constructs.IConstruct) *awskinesisfirehose.CfnDeliveryStream_ProcessorProperty {
	if err := l.validateBindParameters(_scope); err != nil {
		panic(err)
	}
	var returns *awskinesisfirehose.CfnDeliveryStream_ProcessorProperty

	_jsii_.Invoke(
		l,
		"bind",
		[]interface{}{_scope},
		&returns,
	)

	return returns
}

