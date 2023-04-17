package ec2


// Optional configuration for the IPAM scope resource.
type PrivateIpamScopeOptions struct {
	// The description of the scope.
	// See: [IPAMScope Description](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ipamscope.html#cfn-ec2-ipamscope-description)
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
}

