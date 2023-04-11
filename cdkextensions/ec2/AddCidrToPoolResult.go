package ec2


type AddCidrToPoolResult struct {
	Inline *bool `field:"required" json:"inline" yaml:"inline"`
	Cidr IIpamPoolCidr `field:"optional" json:"cidr" yaml:"cidr"`
}

