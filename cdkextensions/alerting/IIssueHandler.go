package alerting

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
	"github.com/vibe-io/cdk-extensions-go/cdkextensions/alerting/internal"
)

type IIssueHandler interface {
	awscdk.IResource
	Handler() awsstepfunctions.IStateMachine
	Name() *string
}

// The jsii proxy for IIssueHandler
type jsiiProxy_IIssueHandler struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IIssueHandler) Handler() awsstepfunctions.IStateMachine {
	var returns awsstepfunctions.IStateMachine
	_jsii_.Get(
		j,
		"handler",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IIssueHandler) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

