package glue


// Configuration for the Glue SecurityConfiguration resource.
type SecurityConfigurationProps struct {
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
	//    CloudFormation during deployment.
	// - a concrete value implies a specific physical name
	// - `PhysicalName.GENERATE_IF_NEEDED` is a marker that indicates that a physical will only be generated
	//    by the CDK if it is needed for cross-environment references. Otherwise, it will be allocated by CloudFormation.
	PhysicalName *string `field:"optional" json:"physicalName" yaml:"physicalName"`
	// The AWS region this resource belongs to.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Cloudwatch Encryption Settings.
	// See: [AWS::Glue::SecurityConfiguration EncryptionConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-securityconfiguration-encryptionconfiguration.html#cfn-glue-securityconfiguration-encryptionconfiguration-cloudwatchencryption)
	//
	CloudWatchEncryption *CloudWatchEncryption `field:"optional" json:"cloudWatchEncryption" yaml:"cloudWatchEncryption"`
	// The encryption configuration for job bookmarks.
	// See: [AWS::Glue::SecurityConfiguration EncryptionConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-securityconfiguration-encryptionconfiguration.html#cfn-glue-securityconfiguration-encryptionconfiguration-jobbookmarksencryption)
	//
	JobBookmarksEncryption *JobBookmarksEncryption `field:"optional" json:"jobBookmarksEncryption" yaml:"jobBookmarksEncryption"`
	// Name for the Security Configuration.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The encyption configuration for Amazon Simple Storage Service (Amazon S3) data.
	// See: [AWS::Glue::SecurityConfiguration EncryptionConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-securityconfiguration-encryptionconfiguration.html#cfn-glue-securityconfiguration-encryptionconfiguration-s3encryptions)
	//
	S3Encryption *S3Encryption `field:"optional" json:"s3Encryption" yaml:"s3Encryption"`
}

