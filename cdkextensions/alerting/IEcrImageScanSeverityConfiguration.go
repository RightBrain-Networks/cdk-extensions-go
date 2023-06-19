package alerting

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type IEcrImageScanSeverityConfiguration interface {
	Levels() *[]EcrImageScanSeverity
}

// The jsii proxy for IEcrImageScanSeverityConfiguration
type jsiiProxy_IEcrImageScanSeverityConfiguration struct {
	_ byte // padding
}

func (j *jsiiProxy_IEcrImageScanSeverityConfiguration) Levels() *[]EcrImageScanSeverity {
	var returns *[]EcrImageScanSeverity
	_jsii_.Get(
		j,
		"levels",
		&returns,
	)
	return returns
}

