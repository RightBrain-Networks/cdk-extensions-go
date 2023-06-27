package guardduty

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Additional configuration of the feature.
// See: [FeatureConfigurations AdditionalConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-guardduty-detector-featureconfigurations.html#cfn-guardduty-detector-featureconfigurations-additionalconfiguration)
//
type IFeatureSetting interface {
	// Status of the additional configuration of a feature.
	// See: [FeatureAdditionalConfiguration Status](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-guardduty-detector-featureadditionalconfiguration.html#cfn-guardduty-detector-featureadditionalconfiguration-status)
	//
	Enabled() *bool
	// Name of the additional configuration of a feature.
	// See: [FeatureAdditionalConfiguration Name](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-guardduty-detector-featureadditionalconfiguration.html#cfn-guardduty-detector-featureadditionalconfiguration-name)
	//
	Name() *string
}

// The jsii proxy for IFeatureSetting
type jsiiProxy_IFeatureSetting struct {
	_ byte // padding
}

func (j *jsiiProxy_IFeatureSetting) Enabled() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFeatureSetting) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

