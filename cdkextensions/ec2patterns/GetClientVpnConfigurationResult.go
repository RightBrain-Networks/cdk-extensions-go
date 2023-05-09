package ec2patterns


type GetClientVpnConfigurationResult struct {
	Cidr *string `field:"required" json:"cidr" yaml:"cidr"`
}

