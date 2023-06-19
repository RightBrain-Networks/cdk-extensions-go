package alerting

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"
)

type SecurityHubSeverity interface {
	LowerBound() *float64
	Name() *string
	Standardized() *string
	UpperBound() *float64
}

// The jsii proxy struct for SecurityHubSeverity
type jsiiProxy_SecurityHubSeverity struct {
	_ byte // padding
}

func (j *jsiiProxy_SecurityHubSeverity) LowerBound() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lowerBound",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityHubSeverity) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityHubSeverity) Standardized() *string {
	var returns *string
	_jsii_.Get(
		j,
		"standardized",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityHubSeverity) UpperBound() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"upperBound",
		&returns,
	)
	return returns
}


func SecurityHubSeverity_Of(name *string, lowerBound *float64, upperBound *float64, standardized *string) SecurityHubSeverity {
	_init_.Initialize()

	if err := validateSecurityHubSeverity_OfParameters(name, lowerBound, upperBound, standardized); err != nil {
		panic(err)
	}
	var returns SecurityHubSeverity

	_jsii_.StaticInvoke(
		"cdk-extensions.alerting.SecurityHubSeverity",
		"of",
		[]interface{}{name, lowerBound, upperBound, standardized},
		&returns,
	)

	return returns
}

func SecurityHubSeverity_CRITICAL() SecurityHubSeverity {
	_init_.Initialize()
	var returns SecurityHubSeverity
	_jsii_.StaticGet(
		"cdk-extensions.alerting.SecurityHubSeverity",
		"CRITICAL",
		&returns,
	)
	return returns
}

func SecurityHubSeverity_HIGH() SecurityHubSeverity {
	_init_.Initialize()
	var returns SecurityHubSeverity
	_jsii_.StaticGet(
		"cdk-extensions.alerting.SecurityHubSeverity",
		"HIGH",
		&returns,
	)
	return returns
}

func SecurityHubSeverity_INFORMATIONAL() SecurityHubSeverity {
	_init_.Initialize()
	var returns SecurityHubSeverity
	_jsii_.StaticGet(
		"cdk-extensions.alerting.SecurityHubSeverity",
		"INFORMATIONAL",
		&returns,
	)
	return returns
}

func SecurityHubSeverity_LOW() SecurityHubSeverity {
	_init_.Initialize()
	var returns SecurityHubSeverity
	_jsii_.StaticGet(
		"cdk-extensions.alerting.SecurityHubSeverity",
		"LOW",
		&returns,
	)
	return returns
}

func SecurityHubSeverity_MEDIUM() SecurityHubSeverity {
	_init_.Initialize()
	var returns SecurityHubSeverity
	_jsii_.StaticGet(
		"cdk-extensions.alerting.SecurityHubSeverity",
		"MEDIUM",
		&returns,
	)
	return returns
}

