package ec2patterns

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscertificatemanager"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/vibe-io/cdk-extensions-go/cdkextensions/ec2"
)

type AddIsolatedClientVpnEndpointOptions struct {
	ServerCertificate awscertificatemanager.ICertificate `field:"required" json:"serverCertificate" yaml:"serverCertificate"`
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
	SelfServicePortal *bool `field:"optional" json:"selfServicePortal" yaml:"selfServicePortal"`
	SplitTunnel *bool `field:"optional" json:"splitTunnel" yaml:"splitTunnel"`
	SubnetCidr ec2.IIpv4CidrAssignment `field:"optional" json:"subnetCidr" yaml:"subnetCidr"`
	TransportProtocol awsec2.TransportProtocol `field:"optional" json:"transportProtocol" yaml:"transportProtocol"`
	UserBasedAuthentication awsec2.ClientVpnUserBasedAuthentication `field:"optional" json:"userBasedAuthentication" yaml:"userBasedAuthentication"`
	VpnCidr ec2.IIpv4CidrAssignment `field:"optional" json:"vpnCidr" yaml:"vpnCidr"`
}

