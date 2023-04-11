package ec2


type AddCidrToPoolOptions struct {
	Configuration IIpamPoolCidrConfiguration `field:"required" json:"configuration" yaml:"configuration"`
	AllowInline *bool `field:"optional" json:"allowInline" yaml:"allowInline"`
}

