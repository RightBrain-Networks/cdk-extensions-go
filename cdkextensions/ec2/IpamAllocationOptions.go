package ec2


type IpamAllocationOptions struct {
	Allocation IIpamAllocationConfiguration `field:"optional" json:"allocation" yaml:"allocation"`
	Description *string `field:"optional" json:"description" yaml:"description"`
}

