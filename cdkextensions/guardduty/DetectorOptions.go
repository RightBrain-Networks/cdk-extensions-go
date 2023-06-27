package guardduty


// Optional configuration for the detector resource.
type DetectorOptions struct {
	// Specifies whether the detector is to be enabled on creation.
	// See: [Detector Enable](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-detector.html#cfn-guardduty-detector-enable)
	//
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// Collection of additional features to be configured on the detector.
	// See: [Detector Features](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-detector.html#cfn-guardduty-detector-features)
	//
	Features *[]IFeature `field:"optional" json:"features" yaml:"features"`
	// Specifies how frequently updated findings are exported.
	// See: [Detector FindingPublishingFrequency](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-detector.html#cfn-guardduty-detector-findingpublishingfrequency)
	//
	PublishingFrequency FindingPublishingFrequency `field:"optional" json:"publishingFrequency" yaml:"publishingFrequency"`
}

