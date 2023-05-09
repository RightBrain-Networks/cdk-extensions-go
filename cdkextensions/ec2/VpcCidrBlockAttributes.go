package ec2


type VpcCidrBlockAttributes struct {
	AssociationId *string `field:"required" json:"associationId" yaml:"associationId"`
	Cidr *string `field:"required" json:"cidr" yaml:"cidr"`
}

