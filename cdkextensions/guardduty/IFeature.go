package guardduty

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Represents a feature that will be configured for a GuardDuty detector.
// See: [DetectorFeatureConfiguration](https://docs.aws.amazon.com/guardduty/latest/APIReference/API_DetectorFeatureConfiguration.html)
//
type IFeature interface {
	// Provides a data source that must be enabled for the feature to function.
	// See: [Detector DataSources](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-detector.html#cfn-guardduty-detector-datasources)
	//
	DataSource() IDataSource
	// Whether the feature should be enabled or not.
	// See: [FeatureConfigurations Status](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-guardduty-detector-featureconfigurations.html#cfn-guardduty-detector-featureconfigurations-status)
	//
	Enabled() *bool
	// Name of the feature.
	// See: [FeatureConfigurations Name](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-guardduty-detector-featureconfigurations.html#cfn-guardduty-detector-featureconfigurations-name)
	//
	Name() *string
	// Additional configuration of the feature.
	// See: [FeatureConfigurations AdditionalConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-guardduty-detector-featureconfigurations.html#cfn-guardduty-detector-featureconfigurations-additionalconfiguration)
	//
	Settings() *[]IFeatureSetting
}

// The jsii proxy for IFeature
type jsiiProxy_IFeature struct {
	_ byte // padding
}

func (j *jsiiProxy_IFeature) DataSource() IDataSource {
	var returns IDataSource
	_jsii_.Get(
		j,
		"dataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFeature) Enabled() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFeature) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFeature) Settings() *[]IFeatureSetting {
	var returns *[]IFeatureSetting
	_jsii_.Get(
		j,
		"settings",
		&returns,
	)
	return returns
}

