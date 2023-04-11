package ec2


type ResolvedIpamAllocationConfiguration struct {
	Cidr *string `field:"optional" json:"cidr" yaml:"cidr"`
	NetmaskLength *float64 `field:"optional" json:"netmaskLength" yaml:"netmaskLength"`
}

