package aps

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"
)

// A template that can be used for formatting alerts sent by alert manager.
// See: [Notification template examples](https://prometheus.io/docs/alerting/latest/notification_examples/)
//
type AlertManagerTemplate interface {
	// The template content as a string.
	//
	// Uses the Go templating system.
	Content() *string
}

// The jsii proxy struct for AlertManagerTemplate
type jsiiProxy_AlertManagerTemplate struct {
	_ byte // padding
}

func (j *jsiiProxy_AlertManagerTemplate) Content() *string {
	var returns *string
	_jsii_.Get(
		j,
		"content",
		&returns,
	)
	return returns
}


// Loads an alert template from the local filesystem.
//
// Returns: An object representing the template that can be used when
// configuring alert manager for Amzon APS.
func AlertManagerTemplate_FromFile(path *string) AlertManagerTemplate {
	_init_.Initialize()

	if err := validateAlertManagerTemplate_FromFileParameters(path); err != nil {
		panic(err)
	}
	var returns AlertManagerTemplate

	_jsii_.StaticInvoke(
		"cdk-extensions.aps.AlertManagerTemplate",
		"fromFile",
		[]interface{}{path},
		&returns,
	)

	return returns
}

// Loads an alert manager template using a template string.
//
// Returns: An object representing the template that can be used when
// configuring alert manager for Amzon APS.
func AlertManagerTemplate_FromString(content *string) AlertManagerTemplate {
	_init_.Initialize()

	if err := validateAlertManagerTemplate_FromStringParameters(content); err != nil {
		panic(err)
	}
	var returns AlertManagerTemplate

	_jsii_.StaticInvoke(
		"cdk-extensions.aps.AlertManagerTemplate",
		"fromString",
		[]interface{}{content},
		&returns,
	)

	return returns
}

