package config

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type IRemediationConfiguration interface {
	RemediationConfigurationArn() *string
	RemediationConfigurationName() *string
}

// The jsii proxy for IRemediationConfiguration
type jsiiProxy_IRemediationConfiguration struct {
	_ byte // padding
}

func (j *jsiiProxy_IRemediationConfiguration) RemediationConfigurationArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"remediationConfigurationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRemediationConfiguration) RemediationConfigurationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"remediationConfigurationName",
		&returns,
	)
	return returns
}

