package securityhub

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type IHub interface {
	HubArn() *string
	HubName() *string
}

// The jsii proxy for IHub
type jsiiProxy_IHub struct {
	_ byte // padding
}

func (j *jsiiProxy_IHub) HubArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hubArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IHub) HubName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hubName",
		&returns,
	)
	return returns
}

