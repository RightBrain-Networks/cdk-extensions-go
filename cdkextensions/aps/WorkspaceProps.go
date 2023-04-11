package aps


// Configuration for the Workspace resource.
type WorkspaceProps struct {
	// The AWS account ID this resource belongs to.
	Account *string `field:"optional" json:"account" yaml:"account"`
	// ARN to deduce region and account from.
	//
	// The ARN is parsed and the account and region are taken from the ARN.
	// This should be used for imported resources.
	//
	// Cannot be supplied together with either `account` or `region`.
	EnvironmentFromArn *string `field:"optional" json:"environmentFromArn" yaml:"environmentFromArn"`
	// The value passed in by users to the physical name prop of the resource.
	//
	// - `undefined` implies that a physical name will be allocated by
	//   CloudFormation during deployment.
	// - a concrete value implies a specific physical name
	// - `PhysicalName.GENERATE_IF_NEEDED` is a marker that indicates that a physical will only be generated
	//   by the CDK if it is needed for cross-environment references. Otherwise, it will be allocated by CloudFormation.
	PhysicalName *string `field:"optional" json:"physicalName" yaml:"physicalName"`
	// The AWS region this resource belongs to.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// The details used to configure alerting for the APS workspace.
	// See: [Workspace AlertManagerDefinition](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-aps-workspace.html#cfn-aps-workspace-alertmanagerdefinition)
	//
	Alerting *WorkspaceAlertingOptions `field:"optional" json:"alerting" yaml:"alerting"`
	// An alias that you assign to this workspace to help you identify it. It does not need to be unique.
	//
	// The alias can be as many as 100 characters and can include any type of
	// characters.
	// See: [Workspace Alias](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-aps-workspace.html#cfn-aps-workspace-alias)
	//
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// The details used for configuring logging for the APS workspace.
	// See: [Workspace LoggingConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-aps-workspace.html#cfn-aps-workspace-loggingconfiguration)
	//
	Logging *WorkspaceLoggingOptions `field:"optional" json:"logging" yaml:"logging"`
}

