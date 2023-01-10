package glue

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/vibe-io/cdk-extensions-go/cdkextensions/glue/internal"
)

// Represnets a Glue Job in AWS.
type IJob interface {
	constructs.IConstruct
	// The Amazon Resource Name (ARN) of the job.
	JobArn() *string
	// The name of the job.
	JobName() *string
}

// The jsii proxy for IJob
type jsiiProxy_IJob struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IJob) JobArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IJob) JobName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobName",
		&returns,
	)
	return returns
}

