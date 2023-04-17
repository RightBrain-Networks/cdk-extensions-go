package ec2


type AddAwsProvidedIpv6PoolOptions struct {
	Locale *string `field:"required" json:"locale" yaml:"locale"`
	DefaultNetmaskLength *float64 `field:"optional" json:"defaultNetmaskLength" yaml:"defaultNetmaskLength"`
	Description *string `field:"optional" json:"description" yaml:"description"`
	MaxNetmaskLength *float64 `field:"optional" json:"maxNetmaskLength" yaml:"maxNetmaskLength"`
	MinNetmaskLength *float64 `field:"optional" json:"minNetmaskLength" yaml:"minNetmaskLength"`
	Name *string `field:"optional" json:"name" yaml:"name"`
	Netmask *float64 `field:"optional" json:"netmask" yaml:"netmask"`
	TagRestrictions *map[string]*string `field:"optional" json:"tagRestrictions" yaml:"tagRestrictions"`
}

