package guardduty

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"
)

// Specifies how frequently updated findings are exported.
// See: [Detector FindingPublishingFrequency](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-detector.html#cfn-guardduty-detector-findingpublishingfrequency)
//
type FindingPublishingFrequency interface {
	// The label for the publishing frequency value as would be expected by CloudFormation.
	Label() *string
}

// The jsii proxy struct for FindingPublishingFrequency
type jsiiProxy_FindingPublishingFrequency struct {
	_ byte // padding
}

func (j *jsiiProxy_FindingPublishingFrequency) Label() *string {
	var returns *string
	_jsii_.Get(
		j,
		"label",
		&returns,
	)
	return returns
}


// Escape hatch for specifying custom publishing frequency values.
//
// This is primarily intended in case additional options are added and
// support for those values has not yet been officially added.
//
// Where possible it is recommended that the existing provided values be
// used.
//
// Returns: An object that can be used for setting the publishing frequency
// for a detector resource.
func FindingPublishingFrequency_Of(label *string) FindingPublishingFrequency {
	_init_.Initialize()

	if err := validateFindingPublishingFrequency_OfParameters(label); err != nil {
		panic(err)
	}
	var returns FindingPublishingFrequency

	_jsii_.StaticInvoke(
		"cdk-extensions.guardduty.FindingPublishingFrequency",
		"of",
		[]interface{}{label},
		&returns,
	)

	return returns
}

func FindingPublishingFrequency_FIFTEEN_MINUTES() FindingPublishingFrequency {
	_init_.Initialize()
	var returns FindingPublishingFrequency
	_jsii_.StaticGet(
		"cdk-extensions.guardduty.FindingPublishingFrequency",
		"FIFTEEN_MINUTES",
		&returns,
	)
	return returns
}

func FindingPublishingFrequency_ONE_HOUR() FindingPublishingFrequency {
	_init_.Initialize()
	var returns FindingPublishingFrequency
	_jsii_.StaticGet(
		"cdk-extensions.guardduty.FindingPublishingFrequency",
		"ONE_HOUR",
		&returns,
	)
	return returns
}

func FindingPublishingFrequency_SIX_HOURS() FindingPublishingFrequency {
	_init_.Initialize()
	var returns FindingPublishingFrequency
	_jsii_.StaticGet(
		"cdk-extensions.guardduty.FindingPublishingFrequency",
		"SIX_HOURS",
		&returns,
	)
	return returns
}

