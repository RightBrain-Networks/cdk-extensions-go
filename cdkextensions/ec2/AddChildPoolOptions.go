package ec2


type AddChildPoolOptions struct {
	AddressConfiguration AddressConfiguration `field:"optional" json:"addressConfiguration" yaml:"addressConfiguration"`
	AutoImport *bool `field:"optional" json:"autoImport" yaml:"autoImport"`
	DefaultNetmaskLength *float64 `field:"optional" json:"defaultNetmaskLength" yaml:"defaultNetmaskLength"`
	Description *string `field:"optional" json:"description" yaml:"description"`
	Locale *string `field:"optional" json:"locale" yaml:"locale"`
	MaxNetmaskLength *float64 `field:"optional" json:"maxNetmaskLength" yaml:"maxNetmaskLength"`
	MinNetmaskLength *float64 `field:"optional" json:"minNetmaskLength" yaml:"minNetmaskLength"`
	Name *string `field:"optional" json:"name" yaml:"name"`
	ProvisionedCidrs *[]*string `field:"optional" json:"provisionedCidrs" yaml:"provisionedCidrs"`
	TagRestrictions *map[string]*string `field:"optional" json:"tagRestrictions" yaml:"tagRestrictions"`
}

