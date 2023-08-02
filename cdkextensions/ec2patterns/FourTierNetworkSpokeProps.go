package ec2patterns

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/vibe-io/cdk-extensions-go/cdkextensions/ec2"
)

type FourTierNetworkSpokeProps struct {
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
	Hub FourTierNetworkHub `field:"required" json:"hub" yaml:"hub"`
	AddressManager IpAddressManager `field:"optional" json:"addressManager" yaml:"addressManager"`
	AvailabilityZones *[]*string `field:"optional" json:"availabilityZones" yaml:"availabilityZones"`
	Cidr ec2.IIpv4CidrAssignment `field:"optional" json:"cidr" yaml:"cidr"`
	ClientVpnPool ec2.IIpamPool `field:"optional" json:"clientVpnPool" yaml:"clientVpnPool"`
	DefaultInstanceTenancy awsec2.DefaultInstanceTenancy `field:"optional" json:"defaultInstanceTenancy" yaml:"defaultInstanceTenancy"`
	EnableDnsHostnames *bool `field:"optional" json:"enableDnsHostnames" yaml:"enableDnsHostnames"`
	EnableDnsSupport *bool `field:"optional" json:"enableDnsSupport" yaml:"enableDnsSupport"`
	FlowLogs *map[string]*awsec2.FlowLogOptions `field:"optional" json:"flowLogs" yaml:"flowLogs"`
	GatewayEndpoints *map[string]*awsec2.GatewayVpcEndpointOptions `field:"optional" json:"gatewayEndpoints" yaml:"gatewayEndpoints"`
	MaxAzs *float64 `field:"optional" json:"maxAzs" yaml:"maxAzs"`
	VpcName *string `field:"optional" json:"vpcName" yaml:"vpcName"`
	VpnConnections *map[string]*awsec2.VpnConnectionOptions `field:"optional" json:"vpnConnections" yaml:"vpnConnections"`
	VpnGateway *bool `field:"optional" json:"vpnGateway" yaml:"vpnGateway"`
	VpnGatewayAsn *float64 `field:"optional" json:"vpnGatewayAsn" yaml:"vpnGatewayAsn"`
	VpnRoutePropagation *[]*awsec2.SubnetSelection `field:"optional" json:"vpnRoutePropagation" yaml:"vpnRoutePropagation"`
}

