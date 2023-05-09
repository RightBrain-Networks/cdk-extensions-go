package ec2


type CidrAssignmentIpamDetails struct {
	Family AddressFamily `field:"required" json:"family" yaml:"family"`
	Netmask *float64 `field:"required" json:"netmask" yaml:"netmask"`
	AmazonAllocated *bool `field:"optional" json:"amazonAllocated" yaml:"amazonAllocated"`
	IpamPool IIpamPool `field:"optional" json:"ipamPool" yaml:"ipamPool"`
}

