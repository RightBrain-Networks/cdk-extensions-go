package ec2


type CidrAssignmentCidrDetails struct {
	Cidr *string `field:"required" json:"cidr" yaml:"cidr"`
	Family AddressFamily `field:"required" json:"family" yaml:"family"`
	Netmask *float64 `field:"required" json:"netmask" yaml:"netmask"`
}

