package ec2


type Ipv4CidrAssignmentIpamPoolOptions struct {
	Pool IIpv4IpamPool `field:"required" json:"pool" yaml:"pool"`
	AllocationId *string `field:"optional" json:"allocationId" yaml:"allocationId"`
	Netmask *float64 `field:"optional" json:"netmask" yaml:"netmask"`
}

