package alerting

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/vibe-io/cdk-extensions-go/cdkextensions/alerting/internal"
)

type IIssueParser interface {
	awscdk.IResource
	Bind(node constructs.IConstruct) *[]IssueTrigger
	Handler() awsstepfunctions.IStateMachine
	MatchType() *string
}

// The jsii proxy for IIssueParser
type jsiiProxy_IIssueParser struct {
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IIssueParser) Bind(node constructs.IConstruct) *[]IssueTrigger {
	if err := i.validateBindParameters(node); err != nil {
		panic(err)
	}
	var returns *[]IssueTrigger

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{node},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IIssueParser) Handler() awsstepfunctions.IStateMachine {
	var returns awsstepfunctions.IStateMachine
	_jsii_.Get(
		j,
		"handler",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IIssueParser) MatchType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"matchType",
		&returns,
	)
	return returns
}

