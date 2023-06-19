package alerting

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"
)

type IssueHandlerOverride interface {
	Handler() IIssueHandler
	Overrides() *map[string]interface{}
}

// The jsii proxy struct for IssueHandlerOverride
type jsiiProxy_IssueHandlerOverride struct {
	_ byte // padding
}

func (j *jsiiProxy_IssueHandlerOverride) Handler() IIssueHandler {
	var returns IIssueHandler
	_jsii_.Get(
		j,
		"handler",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IssueHandlerOverride) Overrides() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"overrides",
		&returns,
	)
	return returns
}


func NewIssueHandlerOverride(handler IIssueHandler, overrides *map[string]interface{}) IssueHandlerOverride {
	_init_.Initialize()

	if err := validateNewIssueHandlerOverrideParameters(handler, overrides); err != nil {
		panic(err)
	}
	j := jsiiProxy_IssueHandlerOverride{}

	_jsii_.Create(
		"cdk-extensions.alerting.IssueHandlerOverride",
		[]interface{}{handler, overrides},
		&j,
	)

	return &j
}

func NewIssueHandlerOverride_Override(i IssueHandlerOverride, handler IIssueHandler, overrides *map[string]interface{}) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.alerting.IssueHandlerOverride",
		[]interface{}{handler, overrides},
		i,
	)
}

