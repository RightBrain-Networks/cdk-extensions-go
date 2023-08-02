package ec2patterns

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscertificatemanager"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/vibe-io/cdk-extensions-go/cdkextensions/ec2"
)

type NetworkIsolatedClientVpnEndpointProps struct {
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
	ServerCertificate awscertificatemanager.ICertificate `field:"required" json:"serverCertificate" yaml:"serverCertificate"`
	SubnetCidr ec2.IIpv4CidrAssignment `field:"required" json:"subnetCidr" yaml:"subnetCidr"`
	Vpc awsec2.IVpc `field:"required" json:"vpc" yaml:"vpc"`
	AuthorizeAllUsersToVpcCidr *bool `field:"optional" json:"authorizeAllUsersToVpcCidr" yaml:"authorizeAllUsersToVpcCidr"`
	ClientCertificate awscertificatemanager.ICertificate `field:"optional" json:"clientCertificate" yaml:"clientCertificate"`
	ClientConnectionHandler awsec2.IClientVpnConnectionHandler `field:"optional" json:"clientConnectionHandler" yaml:"clientConnectionHandler"`
	ClientLoginBanner *string `field:"optional" json:"clientLoginBanner" yaml:"clientLoginBanner"`
	Description *string `field:"optional" json:"description" yaml:"description"`
	DnsServers *[]*string `field:"optional" json:"dnsServers" yaml:"dnsServers"`
	Logging *bool `field:"optional" json:"logging" yaml:"logging"`
	LogGroup awslogs.ILogGroup `field:"optional" json:"logGroup" yaml:"logGroup"`
	LogStream awslogs.ILogStream `field:"optional" json:"logStream" yaml:"logStream"`
	MaxAzs *float64 `field:"optional" json:"maxAzs" yaml:"maxAzs"`
	Port awsec2.VpnPort `field:"optional" json:"port" yaml:"port"`
	SecurityGroups *[]awsec2.ISecurityGroup `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	SelfServicePortal *bool `field:"optional" json:"selfServicePortal" yaml:"selfServicePortal"`
	SplitTunnel *bool `field:"optional" json:"splitTunnel" yaml:"splitTunnel"`
	TransitGateway ec2.ITransitGateway `field:"optional" json:"transitGateway" yaml:"transitGateway"`
	TransportProtocol awsec2.TransportProtocol `field:"optional" json:"transportProtocol" yaml:"transportProtocol"`
	UserBasedAuthentication awsec2.ClientVpnUserBasedAuthentication `field:"optional" json:"userBasedAuthentication" yaml:"userBasedAuthentication"`
	VpnCidr ec2.IIpv4CidrAssignment `field:"optional" json:"vpnCidr" yaml:"vpnCidr"`
}

