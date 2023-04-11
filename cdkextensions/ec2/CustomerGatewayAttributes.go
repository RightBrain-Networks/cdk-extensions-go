package ec2


// Attributes used to import an existing customer gateway.
type CustomerGatewayAttributes struct {
	// The ID of the existing customer gateway being imported.
	CustomerGatewayId *string `field:"required" json:"customerGatewayId" yaml:"customerGatewayId"`
	// For devices that support BGP, the customer gateway's BGP ASN.
	BgpAsn *float64 `field:"optional" json:"bgpAsn" yaml:"bgpAsn"`
	// The Internet-routable IP address for the customer gateway's outside interface.
	//
	// The address must be static.
	IpAddress *string `field:"optional" json:"ipAddress" yaml:"ipAddress"`
}

