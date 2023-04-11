package ec2


// Configuration for TransitGateway resource.
type TransitGatewayProps struct {
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
	AmazonSideAsn *float64 `field:"optional" json:"amazonSideAsn" yaml:"amazonSideAsn"`
	AutoAcceptSharedAttachments *bool `field:"optional" json:"autoAcceptSharedAttachments" yaml:"autoAcceptSharedAttachments"`
	CidrBlocks *[]*string `field:"optional" json:"cidrBlocks" yaml:"cidrBlocks"`
	DefaultRouteTableAssociation *bool `field:"optional" json:"defaultRouteTableAssociation" yaml:"defaultRouteTableAssociation"`
	DefaultRouteTableId *string `field:"optional" json:"defaultRouteTableId" yaml:"defaultRouteTableId"`
	DefaultRouteTablePropagation *bool `field:"optional" json:"defaultRouteTablePropagation" yaml:"defaultRouteTablePropagation"`
	Description *string `field:"optional" json:"description" yaml:"description"`
	DnsSupport *bool `field:"optional" json:"dnsSupport" yaml:"dnsSupport"`
	MulticastSupport *bool `field:"optional" json:"multicastSupport" yaml:"multicastSupport"`
	Name *string `field:"optional" json:"name" yaml:"name"`
	VpnEcmpSupport *bool `field:"optional" json:"vpnEcmpSupport" yaml:"vpnEcmpSupport"`
}

