package ec2


type ResolvedIpamPoolCidrConfiguration struct {
	Cidr *string `field:"optional" json:"cidr" yaml:"cidr"`
	NetmaskLength *float64 `field:"optional" json:"netmaskLength" yaml:"netmaskLength"`
}

