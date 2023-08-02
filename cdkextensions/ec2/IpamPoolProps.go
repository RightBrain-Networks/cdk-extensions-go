package ec2


type IpamPoolProps struct {
	// The AWS account ID this resource belongs to.
	// Default: - the resource is in the same account as the stack it belongs to.
	//
	Account *string `field:"optional" json:"account" yaml:"account"`
	// ARN to deduce region and account from.
	//
	// The ARN is parsed and the account and region are taken from the ARN.
	// This should be used for imported resources.
	//
	// Cannot be supplied together with either `account` or `region`.
	// Default: - take environment from `account`, `region` parameters, or use Stack environment.
	//
	EnvironmentFromArn *string `field:"optional" json:"environmentFromArn" yaml:"environmentFromArn"`
	// The value passed in by users to the physical name prop of the resource.
	//
	// - `undefined` implies that a physical name will be allocated by
	//   CloudFormation during deployment.
	// - a concrete value implies a specific physical name
	// - `PhysicalName.GENERATE_IF_NEEDED` is a marker that indicates that a physical will only be generated
	//   by the CDK if it is needed for cross-environment references. Otherwise, it will be allocated by CloudFormation.
	// Default: - The physical name will be allocated by CloudFormation at deployment time.
	//
	PhysicalName *string `field:"optional" json:"physicalName" yaml:"physicalName"`
	// The AWS region this resource belongs to.
	// Default: - the resource is in the same region as the stack it belongs to.
	//
	Region *string `field:"optional" json:"region" yaml:"region"`
	AddressConfiguration AddressConfiguration `field:"optional" json:"addressConfiguration" yaml:"addressConfiguration"`
	AutoImport *bool `field:"optional" json:"autoImport" yaml:"autoImport"`
	Description *string `field:"optional" json:"description" yaml:"description"`
	Locale *string `field:"optional" json:"locale" yaml:"locale"`
	Name *string `field:"optional" json:"name" yaml:"name"`
	ParentPool IIpamPool `field:"optional" json:"parentPool" yaml:"parentPool"`
	ProvisionedCidrs *[]*string `field:"optional" json:"provisionedCidrs" yaml:"provisionedCidrs"`
	PublicIpSource PublicIpSource `field:"optional" json:"publicIpSource" yaml:"publicIpSource"`
	TagRestrictions *map[string]*string `field:"optional" json:"tagRestrictions" yaml:"tagRestrictions"`
	IpamScope IIpamScope `field:"required" json:"ipamScope" yaml:"ipamScope"`
}

