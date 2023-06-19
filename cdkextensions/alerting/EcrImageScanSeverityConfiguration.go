package alerting

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"
)

type EcrImageScanSeverityConfiguration interface {
}

// The jsii proxy struct for EcrImageScanSeverityConfiguration
type jsiiProxy_EcrImageScanSeverityConfiguration struct {
	_ byte // padding
}

func NewEcrImageScanSeverityConfiguration() EcrImageScanSeverityConfiguration {
	_init_.Initialize()

	j := jsiiProxy_EcrImageScanSeverityConfiguration{}

	_jsii_.Create(
		"cdk-extensions.alerting.EcrImageScanSeverityConfiguration",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewEcrImageScanSeverityConfiguration_Override(e EcrImageScanSeverityConfiguration) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.alerting.EcrImageScanSeverityConfiguration",
		nil, // no parameters
		e,
	)
}

func EcrImageScanSeverityConfiguration_All() IEcrImageScanSeverityConfiguration {
	_init_.Initialize()

	var returns IEcrImageScanSeverityConfiguration

	_jsii_.StaticInvoke(
		"cdk-extensions.alerting.EcrImageScanSeverityConfiguration",
		"all",
		nil, // no parameters
		&returns,
	)

	return returns
}

func EcrImageScanSeverityConfiguration_Custom(levels ...EcrImageScanSeverity) IEcrImageScanSeverityConfiguration {
	_init_.Initialize()

	args := []interface{}{}
	for _, a := range levels {
		args = append(args, a)
	}

	var returns IEcrImageScanSeverityConfiguration

	_jsii_.StaticInvoke(
		"cdk-extensions.alerting.EcrImageScanSeverityConfiguration",
		"custom",
		args,
		&returns,
	)

	return returns
}

func EcrImageScanSeverityConfiguration_Threshold(level EcrImageScanSeverity) IEcrImageScanSeverityConfiguration {
	_init_.Initialize()

	if err := validateEcrImageScanSeverityConfiguration_ThresholdParameters(level); err != nil {
		panic(err)
	}
	var returns IEcrImageScanSeverityConfiguration

	_jsii_.StaticInvoke(
		"cdk-extensions.alerting.EcrImageScanSeverityConfiguration",
		"threshold",
		[]interface{}{level},
		&returns,
	)

	return returns
}

