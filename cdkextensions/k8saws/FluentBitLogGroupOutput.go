package k8saws

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
)

// Represents valid log group output configuration options to be used by Fluent Bit when writing to CloudWatch Logs.
type FluentBitLogGroupOutput interface {
	// Flag that determines whether or not a log group should be automatically created.
	AutoCreate() *bool
	// A log group resource object to use as the destination.
	LogGroup() awslogs.ILogGroup
	// The name for the log group that should be used for output records.
	LogGroupName() *string
}

// The jsii proxy struct for FluentBitLogGroupOutput
type jsiiProxy_FluentBitLogGroupOutput struct {
	_ byte // padding
}

func (j *jsiiProxy_FluentBitLogGroupOutput) AutoCreate() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"autoCreate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitLogGroupOutput) LogGroup() awslogs.ILogGroup {
	var returns awslogs.ILogGroup
	_jsii_.Get(
		j,
		"logGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitLogGroupOutput) LogGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logGroupName",
		&returns,
	)
	return returns
}


// Sets a flag saying that a log group should be created automatically.
//
// Depending on the configuration of the plugin, this flag will either cause
// permissions to be granted for Fluent Bit to create the log group itself or
// the plugin CDK resource will create a Log Group and use that as the
// destination.
//
// Returns: A FluentBitLogGroupOutput object representing the configured log
// group destination.
func FluentBitLogGroupOutput_Create() FluentBitLogGroupOutput {
	_init_.Initialize()

	var returns FluentBitLogGroupOutput

	_jsii_.StaticInvoke(
		"cdk-extensions.k8s_aws.FluentBitLogGroupOutput",
		"create",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Sets the destination log group to a LogGroup CDK resource.
//
// Returns: A FluentBitLogGroupOutput object representing the configured log
// group destination.
func FluentBitLogGroupOutput_FromLogGroup(logGroup awslogs.ILogGroup) FluentBitLogGroupOutput {
	_init_.Initialize()

	if err := validateFluentBitLogGroupOutput_FromLogGroupParameters(logGroup); err != nil {
		panic(err)
	}
	var returns FluentBitLogGroupOutput

	_jsii_.StaticInvoke(
		"cdk-extensions.k8s_aws.FluentBitLogGroupOutput",
		"fromLogGroup",
		[]interface{}{logGroup},
		&returns,
	)

	return returns
}

// Sets the destination for logs to the named log group.
//
// Returns: A FluentBitLogGroupOutput object representing the configured log
// group destination.
func FluentBitLogGroupOutput_FromName(name *string, create *bool) FluentBitLogGroupOutput {
	_init_.Initialize()

	if err := validateFluentBitLogGroupOutput_FromNameParameters(name); err != nil {
		panic(err)
	}
	var returns FluentBitLogGroupOutput

	_jsii_.StaticInvoke(
		"cdk-extensions.k8s_aws.FluentBitLogGroupOutput",
		"fromName",
		[]interface{}{name, create},
		&returns,
	)

	return returns
}

