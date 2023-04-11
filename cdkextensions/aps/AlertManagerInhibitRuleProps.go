package aps

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"
)

// Configuration options for the alert manager inhibit rule.
type AlertManagerInhibitRuleProps interface {
	// Labels that must have an equal value in the source and target alert for the inhibition to take effect.
	EqualLabels() *[]*string
	// A list of matchers for which one or more alerts have to exist for the inhibition to take effect.
	SourceMatchers() *[]AlertManagerMatcher
	// A list of matchers that have to be fulfilled by the target alerts to be muted.
	TargetMatchers() *[]AlertManagerMatcher
}

// The jsii proxy struct for AlertManagerInhibitRuleProps
type jsiiProxy_AlertManagerInhibitRuleProps struct {
	_ byte // padding
}

func (j *jsiiProxy_AlertManagerInhibitRuleProps) EqualLabels() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"equalLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertManagerInhibitRuleProps) SourceMatchers() *[]AlertManagerMatcher {
	var returns *[]AlertManagerMatcher
	_jsii_.Get(
		j,
		"sourceMatchers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertManagerInhibitRuleProps) TargetMatchers() *[]AlertManagerMatcher {
	var returns *[]AlertManagerMatcher
	_jsii_.Get(
		j,
		"targetMatchers",
		&returns,
	)
	return returns
}


func NewAlertManagerInhibitRuleProps() AlertManagerInhibitRuleProps {
	_init_.Initialize()

	j := jsiiProxy_AlertManagerInhibitRuleProps{}

	_jsii_.Create(
		"cdk-extensions.aps.AlertManagerInhibitRuleProps",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewAlertManagerInhibitRuleProps_Override(a AlertManagerInhibitRuleProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.aps.AlertManagerInhibitRuleProps",
		nil, // no parameters
		a,
	)
}

