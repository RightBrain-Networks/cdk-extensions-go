package glue

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/vibe-io/cdk-extensions-go/cdkextensions/glue/internal"
)

// Represents a Glue Trigger in AWS.
type ITrigger interface {
	constructs.IConstruct
	// The Amazon Resource Name (ARN) of the trigger.
	TriggerArn() *string
	// The name of the trigger.
	TriggerName() *string
}

// The jsii proxy for ITrigger
type jsiiProxy_ITrigger struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ITrigger) TriggerArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"triggerArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITrigger) TriggerName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"triggerName",
		&returns,
	)
	return returns
}

