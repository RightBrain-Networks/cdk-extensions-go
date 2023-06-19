package alerting

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"
)

type EcrImageScanSeverity interface {
	Name() *string
	Priority() *float64
	Standardized() *string
}

// The jsii proxy struct for EcrImageScanSeverity
type jsiiProxy_EcrImageScanSeverity struct {
	_ byte // padding
}

func (j *jsiiProxy_EcrImageScanSeverity) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcrImageScanSeverity) Priority() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcrImageScanSeverity) Standardized() *string {
	var returns *string
	_jsii_.Get(
		j,
		"standardized",
		&returns,
	)
	return returns
}


func EcrImageScanSeverity_Of(name *string, priority *float64, standardized *string) EcrImageScanSeverity {
	_init_.Initialize()

	if err := validateEcrImageScanSeverity_OfParameters(name, priority, standardized); err != nil {
		panic(err)
	}
	var returns EcrImageScanSeverity

	_jsii_.StaticInvoke(
		"cdk-extensions.alerting.EcrImageScanSeverity",
		"of",
		[]interface{}{name, priority, standardized},
		&returns,
	)

	return returns
}

func EcrImageScanSeverity_CRITICAL() EcrImageScanSeverity {
	_init_.Initialize()
	var returns EcrImageScanSeverity
	_jsii_.StaticGet(
		"cdk-extensions.alerting.EcrImageScanSeverity",
		"CRITICAL",
		&returns,
	)
	return returns
}

func EcrImageScanSeverity_HIGH() EcrImageScanSeverity {
	_init_.Initialize()
	var returns EcrImageScanSeverity
	_jsii_.StaticGet(
		"cdk-extensions.alerting.EcrImageScanSeverity",
		"HIGH",
		&returns,
	)
	return returns
}

func EcrImageScanSeverity_LOW() EcrImageScanSeverity {
	_init_.Initialize()
	var returns EcrImageScanSeverity
	_jsii_.StaticGet(
		"cdk-extensions.alerting.EcrImageScanSeverity",
		"LOW",
		&returns,
	)
	return returns
}

func EcrImageScanSeverity_MEDIUM() EcrImageScanSeverity {
	_init_.Initialize()
	var returns EcrImageScanSeverity
	_jsii_.StaticGet(
		"cdk-extensions.alerting.EcrImageScanSeverity",
		"MEDIUM",
		&returns,
	)
	return returns
}

func EcrImageScanSeverity_UNDEFINED() EcrImageScanSeverity {
	_init_.Initialize()
	var returns EcrImageScanSeverity
	_jsii_.StaticGet(
		"cdk-extensions.alerting.EcrImageScanSeverity",
		"UNDEFINED",
		&returns,
	)
	return returns
}

