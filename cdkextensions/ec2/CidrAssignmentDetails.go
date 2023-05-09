package ec2


type CidrAssignmentDetails struct {
	CidrDetails *CidrAssignmentCidrDetails `field:"optional" json:"cidrDetails" yaml:"cidrDetails"`
	IpamDetails *CidrAssignmentIpamDetails `field:"optional" json:"ipamDetails" yaml:"ipamDetails"`
}

