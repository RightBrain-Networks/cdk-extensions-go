package guardduty

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Represents a GuardDuty detector in AWS.
type IDetector interface {
	// The Amazon Resource Name (ARN) of the GuardDuty detector.
	DetectorArn() *string
	// The ID generated by AWS for the GuardDuty detector.
	DetectorId() *string
}

// The jsii proxy for IDetector
type jsiiProxy_IDetector struct {
	_ byte // padding
}

func (j *jsiiProxy_IDetector) DetectorArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"detectorArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDetector) DetectorId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"detectorId",
		&returns,
	)
	return returns
}

