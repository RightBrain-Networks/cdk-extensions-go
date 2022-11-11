package k8saws

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awskinesis"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsopensearchservice"
	"github.com/vibe-io/cdk-extensions-go/cdkextensions/kinesisfirehose"
)

// Common options that allow configuration of destinations where Fluent Bit should send records after processing.
type FluentBitOutput interface {
}

// The jsii proxy struct for FluentBitOutput
type jsiiProxy_FluentBitOutput struct {
	_ byte // padding
}

func NewFluentBitOutput() FluentBitOutput {
	_init_.Initialize()

	j := jsiiProxy_FluentBitOutput{}

	_jsii_.Create(
		"cdk-extensions.k8s_aws.FluentBitOutput",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewFluentBitOutput_Override(f FluentBitOutput) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.k8s_aws.FluentBitOutput",
		nil, // no parameters
		f,
	)
}

// Sends matched records to a CloudWatch Logs log group.
//
// Returns: An output filter object that can be applied to the Fluent Bit
// configuration.
func FluentBitOutput_CloudwatchLogs(match FluentBitMatch, logGroup awslogs.ILogGroup) IFluentBitOutputPlugin {
	_init_.Initialize()

	if err := validateFluentBitOutput_CloudwatchLogsParameters(match, logGroup); err != nil {
		panic(err)
	}
	var returns IFluentBitOutputPlugin

	_jsii_.StaticInvoke(
		"cdk-extensions.k8s_aws.FluentBitOutput",
		"cloudwatchLogs",
		[]interface{}{match, logGroup},
		&returns,
	)

	return returns
}

// Sends matched records to a Kinesis data stream.
//
// Returns: An output filter object that can be applied to the Fluent Bit
// configuration.
func FluentBitOutput_Kinesis(match FluentBitMatch, stream awskinesis.IStream) IFluentBitOutputPlugin {
	_init_.Initialize()

	if err := validateFluentBitOutput_KinesisParameters(match, stream); err != nil {
		panic(err)
	}
	var returns IFluentBitOutputPlugin

	_jsii_.StaticInvoke(
		"cdk-extensions.k8s_aws.FluentBitOutput",
		"kinesis",
		[]interface{}{match, stream},
		&returns,
	)

	return returns
}

// Sends matched records to a Kinesis Firehose delivery stream.
//
// Returns: An output filter object that can be applied to the Fluent Bit
// configuration.
func FluentBitOutput_KinesisFirehose(match FluentBitMatch, deliveryStream kinesisfirehose.IDeliveryStream) IFluentBitOutputPlugin {
	_init_.Initialize()

	if err := validateFluentBitOutput_KinesisFirehoseParameters(match, deliveryStream); err != nil {
		panic(err)
	}
	var returns IFluentBitOutputPlugin

	_jsii_.StaticInvoke(
		"cdk-extensions.k8s_aws.FluentBitOutput",
		"kinesisFirehose",
		[]interface{}{match, deliveryStream},
		&returns,
	)

	return returns
}

// Sends matched records to an OpenSearch domain.
//
// Returns: An output filter object that can be applied to the Fluent Bit
// configuration.
func FluentBitOutput_Opensearch(match FluentBitMatch, domain awsopensearchservice.IDomain) IFluentBitOutputPlugin {
	_init_.Initialize()

	if err := validateFluentBitOutput_OpensearchParameters(match, domain); err != nil {
		panic(err)
	}
	var returns IFluentBitOutputPlugin

	_jsii_.StaticInvoke(
		"cdk-extensions.k8s_aws.FluentBitOutput",
		"opensearch",
		[]interface{}{match, domain},
		&returns,
	)

	return returns
}

