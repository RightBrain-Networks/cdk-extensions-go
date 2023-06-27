package guardduty


// Configuration for the EKS runtime monitoring GuardDuty detector feature.
type EksRuntimeMonitoringOptions struct {
	// Whether the feature should be enabled or not.
	// See: [FeatureConfigurations Status](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-guardduty-detector-featureconfigurations.html#cfn-guardduty-detector-featureconfigurations-status)
	//
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// Determines whether GuardDuty is allowed to manage addons for the EKS cluster.
	//
	// If both runtime monitoring and addon management are enabled GuardDuty can
	// automatically manage the GuardDuty security agent addon that is used to
	// send findings to GuardDuty.
	//
	// If runtime monitoring is enabled and addon management is disabled you must
	// manage the GuardDuty security agent manually.
	// See: [Configuring EKS Runtime Monitoring](https://docs.aws.amazon.com/guardduty/latest/ug/eks-protection-configuration.html)
	//
	AddonManagement *bool `field:"optional" json:"addonManagement" yaml:"addonManagement"`
}

