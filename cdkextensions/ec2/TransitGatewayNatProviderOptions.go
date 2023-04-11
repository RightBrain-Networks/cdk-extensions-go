package ec2


type TransitGatewayNatProviderOptions struct {
	TransitGateway ITransitGateway `field:"required" json:"transitGateway" yaml:"transitGateway"`
}

