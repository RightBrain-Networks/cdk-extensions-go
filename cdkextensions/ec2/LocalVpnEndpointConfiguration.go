package ec2


// Configuration object containing the vlues needed to configure the local end of a VPN connection.
type LocalVpnEndpointConfiguration struct {
	// The ID of the transit gateway that serves as the local end of the VPN connection.
	TransitGatewayId *string `field:"optional" json:"transitGatewayId" yaml:"transitGatewayId"`
	// The ID of the VPN gateway that serves as the local end of the VPN connection.
	VpnGatewayId *string `field:"optional" json:"vpnGatewayId" yaml:"vpnGatewayId"`
}

