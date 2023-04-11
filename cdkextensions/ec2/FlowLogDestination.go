package ec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents a resource that can act as a deliver endpoint for captured flow logs.
type FlowLogDestination interface {
	ILogDestination
	// Returns a configuration object with all the fields and resources needed to configure a flow log to write to the destination.
	Bind(scope constructs.IConstruct) *FlowLogDestinationConfig
}

// The jsii proxy struct for FlowLogDestination
type jsiiProxy_FlowLogDestination struct {
	jsiiProxy_ILogDestination
}

func NewFlowLogDestination_Override(f FlowLogDestination) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.ec2.FlowLogDestination",
		nil, // no parameters
		f,
	)
}

// Represents a CloudWatch log group that will serve as the endpoint where flow logs should be delivered.
//
// Returns: A configuration object containing details on how to set up
// logging to the log group.
// See: [Publish flow logs to CloudWatch Logs](https://docs.aws.amazon.com/vpc/latest/userguide/flow-logs-cwl.html)
//
func FlowLogDestination_ToCloudWatchLogs(logGroup awslogs.ILogGroup, role awsiam.IRole) FlowLogDestination {
	_init_.Initialize()

	var returns FlowLogDestination

	_jsii_.StaticInvoke(
		"cdk-extensions.ec2.FlowLogDestination",
		"toCloudWatchLogs",
		[]interface{}{logGroup, role},
		&returns,
	)

	return returns
}

// Represents a CloudWatch log group that will serve as the endpoint where flow logs should be delivered.
//
// Returns: A configuration object containing details on how to set up
// logging to the bucket.
// See: [Publish flow logs to Amazon S3](https://docs.aws.amazon.com/vpc/latest/userguide/flow-logs-s3.html)
//
func FlowLogDestination_ToS3(bucket awss3.IBucket, options *FlowLogS3Options) FlowLogDestination {
	_init_.Initialize()

	if err := validateFlowLogDestination_ToS3Parameters(options); err != nil {
		panic(err)
	}
	var returns FlowLogDestination

	_jsii_.StaticInvoke(
		"cdk-extensions.ec2.FlowLogDestination",
		"toS3",
		[]interface{}{bucket, options},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FlowLogDestination) Bind(scope constructs.IConstruct) *FlowLogDestinationConfig {
	if err := f.validateBindParameters(scope); err != nil {
		panic(err)
	}
	var returns *FlowLogDestinationConfig

	_jsii_.Invoke(
		f,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

