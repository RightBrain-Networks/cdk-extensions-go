package guardduty


// Configuration shared by all GuardDuty detector features.
type FeatureOptions struct {
	// Whether the feature should be enabled or not.
	// See: [FeatureConfigurations Status](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-guardduty-detector-featureconfigurations.html#cfn-guardduty-detector-featureconfigurations-status)
	//
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
}

