package ec2


// Configuration for the IPAM resource.
type IpamProps struct {
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
	// The description for the IPAM.
	// See: [IPAM Description](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ipam.html#cfn-ec2-ipam-description)
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The operating Regions for an IPAM.
	//
	// Operating Regions are AWS Regions where
	// the IPAM is allowed to manage IP address CIDRs. IPAM only discovers and
	// monitors resources in the AWS Regions you select as operating Regions.
	// See: [Create an IPAM](https://docs.aws.amazon.com/vpc/latest/ipam/create-ipam.html)
	//
	Regions *[]*string `field:"optional" json:"regions" yaml:"regions"`
}

