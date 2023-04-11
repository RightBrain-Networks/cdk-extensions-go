package ram


// Configuration for ResourceShare resource.
type ResourceShareProps struct {
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
	// Specifies whether principals outside your organization in AWS Organizations can be associated with a resource share.
	//
	// A value of `true`
	// lets you share with individual AWS accounts that are not in your
	// organization. A value of `false` only has meaning if your account is a
	// member of an AWS Organization.
	// See: [ResourceShare.AllowExternalPrinicpals](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ram-resourceshare.html#cfn-ram-resourceshare-allowexternalprincipals)
	//
	AllowExternalPrincipals *bool `field:"optional" json:"allowExternalPrincipals" yaml:"allowExternalPrincipals"`
	// Controls whether the resource share should attempt to search for AWS accounts that are part of the same CDK application.
	//
	// Any accounts is finds
	// will be added to the resource automatically and will be able to use the
	// shared resources.
	AutoDiscoverAccounts *bool `field:"optional" json:"autoDiscoverAccounts" yaml:"autoDiscoverAccounts"`
	// Specifies the name of the resource share.
	// See: [ResourceShare.Name](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ram-resourceshare.html#cfn-ram-resourceshare-name)
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Specifies a list of one or more principals to associate with the resource share.
	// See: [ResourceShare.Prinicipals](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ram-resourceshare.html#cfn-ram-resourceshare-principals)
	//
	Principals *[]ISharedPrincipal `field:"optional" json:"principals" yaml:"principals"`
	// Specifies a list of AWS resources to share with the configured principal accounts and organizations.
	// See: [ResourceShare.ResourceArns](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ram-resourceshare.html#cfn-ram-resourceshare-resourcearns)
	//
	Resources *[]ISharable `field:"optional" json:"resources" yaml:"resources"`
}

