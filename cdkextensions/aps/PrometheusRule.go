package aps

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"
)

type PrometheusRule interface {
}

// The jsii proxy struct for PrometheusRule
type jsiiProxy_PrometheusRule struct {
	_ byte // padding
}

func NewPrometheusRule() PrometheusRule {
	_init_.Initialize()

	j := jsiiProxy_PrometheusRule{}

	_jsii_.Create(
		"cdk-extensions.aps.PrometheusRule",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewPrometheusRule_Override(p PrometheusRule) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.aps.PrometheusRule",
		nil, // no parameters
		p,
	)
}

func PrometheusRule_AlertingRule(options *AlertingRuleProps) AlertingRule {
	_init_.Initialize()

	if err := validatePrometheusRule_AlertingRuleParameters(options); err != nil {
		panic(err)
	}
	var returns AlertingRule

	_jsii_.StaticInvoke(
		"cdk-extensions.aps.PrometheusRule",
		"alertingRule",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func PrometheusRule_RecordingRule(options *RecordingRuleProps) RecordingRule {
	_init_.Initialize()

	if err := validatePrometheusRule_RecordingRuleParameters(options); err != nil {
		panic(err)
	}
	var returns RecordingRule

	_jsii_.StaticInvoke(
		"cdk-extensions.aps.PrometheusRule",
		"recordingRule",
		[]interface{}{options},
		&returns,
	)

	return returns
}

