package ec2


// Configuration object containing the vlues needed to configure the remote end of a VPN connection.
type RemoteVpnEndpointConfiguration struct {
	// The BGP ASN of the customer gateway which is configured with the details of the remote endpoint device.
	CustomerGatewayAsn *float64 `field:"required" json:"customerGatewayAsn" yaml:"customerGatewayAsn"`
	// The ID of the customer gateway which is configured with the details of the remote endpoint device.
	CustomerGatewayId *string `field:"required" json:"customerGatewayId" yaml:"customerGatewayId"`
	// The IP address of the customer gateway which is configured with the details of the remote endpoint device.
	CustomerGatewayIp *string `field:"required" json:"customerGatewayIp" yaml:"customerGatewayIp"`
}

