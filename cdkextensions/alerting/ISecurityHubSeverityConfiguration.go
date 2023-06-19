package alerting

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type ISecurityHubSeverityConfiguration interface {
	Levels() *[]SecurityHubSeverity
}

// The jsii proxy for ISecurityHubSeverityConfiguration
type jsiiProxy_ISecurityHubSeverityConfiguration struct {
	_ byte // padding
}

func (j *jsiiProxy_ISecurityHubSeverityConfiguration) Levels() *[]SecurityHubSeverity {
	var returns *[]SecurityHubSeverity
	_jsii_.Get(
		j,
		"levels",
		&returns,
	)
	return returns
}

