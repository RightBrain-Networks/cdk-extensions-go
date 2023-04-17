package ec2


type AddByoipIpv6PoolOptions struct {
	AdvertiseService AdvertiseService `field:"optional" json:"advertiseService" yaml:"advertiseService"`
	DefaultNetmaskLength *float64 `field:"optional" json:"defaultNetmaskLength" yaml:"defaultNetmaskLength"`
	Description *string `field:"optional" json:"description" yaml:"description"`
	Locale *string `field:"optional" json:"locale" yaml:"locale"`
	MaxNetmaskLength *float64 `field:"optional" json:"maxNetmaskLength" yaml:"maxNetmaskLength"`
	MinNetmaskLength *float64 `field:"optional" json:"minNetmaskLength" yaml:"minNetmaskLength"`
	Name *string `field:"optional" json:"name" yaml:"name"`
	PubliclyAdvertisable *bool `field:"optional" json:"publiclyAdvertisable" yaml:"publiclyAdvertisable"`
	TagRestrictions *map[string]*string `field:"optional" json:"tagRestrictions" yaml:"tagRestrictions"`
}

