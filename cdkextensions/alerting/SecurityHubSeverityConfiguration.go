package alerting

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"
)

type SecurityHubSeverityConfiguration interface {
}

// The jsii proxy struct for SecurityHubSeverityConfiguration
type jsiiProxy_SecurityHubSeverityConfiguration struct {
	_ byte // padding
}

func NewSecurityHubSeverityConfiguration() SecurityHubSeverityConfiguration {
	_init_.Initialize()

	j := jsiiProxy_SecurityHubSeverityConfiguration{}

	_jsii_.Create(
		"cdk-extensions.alerting.SecurityHubSeverityConfiguration",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewSecurityHubSeverityConfiguration_Override(s SecurityHubSeverityConfiguration) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.alerting.SecurityHubSeverityConfiguration",
		nil, // no parameters
		s,
	)
}

func SecurityHubSeverityConfiguration_All() ISecurityHubSeverityConfiguration {
	_init_.Initialize()

	var returns ISecurityHubSeverityConfiguration

	_jsii_.StaticInvoke(
		"cdk-extensions.alerting.SecurityHubSeverityConfiguration",
		"all",
		nil, // no parameters
		&returns,
	)

	return returns
}

func SecurityHubSeverityConfiguration_Custom(levels ...SecurityHubSeverity) ISecurityHubSeverityConfiguration {
	_init_.Initialize()

	args := []interface{}{}
	for _, a := range levels {
		args = append(args, a)
	}

	var returns ISecurityHubSeverityConfiguration

	_jsii_.StaticInvoke(
		"cdk-extensions.alerting.SecurityHubSeverityConfiguration",
		"custom",
		args,
		&returns,
	)

	return returns
}

func SecurityHubSeverityConfiguration_Threshold(level SecurityHubSeverity) ISecurityHubSeverityConfiguration {
	_init_.Initialize()

	if err := validateSecurityHubSeverityConfiguration_ThresholdParameters(level); err != nil {
		panic(err)
	}
	var returns ISecurityHubSeverityConfiguration

	_jsii_.StaticInvoke(
		"cdk-extensions.alerting.SecurityHubSeverityConfiguration",
		"threshold",
		[]interface{}{level},
		&returns,
	)

	return returns
}

