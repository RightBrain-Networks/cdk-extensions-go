package k8saws

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
)

// Represents valid log stream output configuration options to be used by Fluent Bit when writing to CloudWatch Logs.
type FluentBitLogStreamOutput interface {
	// The name of the log stream where records should be created.
	LogStreamName() *string
	// The prefix for log streams that will be created on a per-pod basis.
	LogStreamPrefix() *string
}

// The jsii proxy struct for FluentBitLogStreamOutput
type jsiiProxy_FluentBitLogStreamOutput struct {
	_ byte // padding
}

func (j *jsiiProxy_FluentBitLogStreamOutput) LogStreamName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logStreamName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FluentBitLogStreamOutput) LogStreamPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logStreamPrefix",
		&returns,
	)
	return returns
}


// Sets output to be a log stream resource object.
//
// Returns: A FluentBitLogStreamOutput object representing the configured
// log stream destination.
func FluentBitLogStreamOutput_FromLogStream(logStream awslogs.ILogStream) FluentBitLogStreamOutput {
	_init_.Initialize()

	if err := validateFluentBitLogStreamOutput_FromLogStreamParameters(logStream); err != nil {
		panic(err)
	}
	var returns FluentBitLogStreamOutput

	_jsii_.StaticInvoke(
		"cdk-extensions.k8s_aws.FluentBitLogStreamOutput",
		"fromLogStream",
		[]interface{}{logStream},
		&returns,
	)

	return returns
}

// Sets output to a named log stream.
//
// If a log stream with the given name doesn't exist in the configured log
// group a log stream with the given name will be created.
//
// Returns: A FluentBitLogStreamOutput object representing the configured
// log stream destination.
func FluentBitLogStreamOutput_FromName(name *string) FluentBitLogStreamOutput {
	_init_.Initialize()

	if err := validateFluentBitLogStreamOutput_FromNameParameters(name); err != nil {
		panic(err)
	}
	var returns FluentBitLogStreamOutput

	_jsii_.StaticInvoke(
		"cdk-extensions.k8s_aws.FluentBitLogStreamOutput",
		"fromName",
		[]interface{}{name},
		&returns,
	)

	return returns
}

// Sets output to a prefixed log stream.
//
// Log streams will be created on a per-pod basis with the name oof the log
// streams starting with the provided prefix.
//
// Returns: A FluentBitLogStreamOutput object representing the configured
// log stream destination.
func FluentBitLogStreamOutput_FromPrefix(prefix *string) FluentBitLogStreamOutput {
	_init_.Initialize()

	if err := validateFluentBitLogStreamOutput_FromPrefixParameters(prefix); err != nil {
		panic(err)
	}
	var returns FluentBitLogStreamOutput

	_jsii_.StaticInvoke(
		"cdk-extensions.k8s_aws.FluentBitLogStreamOutput",
		"fromPrefix",
		[]interface{}{prefix},
		&returns,
	)

	return returns
}

