package ec2


type IpamPoolOptions struct {
	AddressConfiguration AddressConfiguration `field:"optional" json:"addressConfiguration" yaml:"addressConfiguration"`
	AutoImport *bool `field:"optional" json:"autoImport" yaml:"autoImport"`
	Description *string `field:"optional" json:"description" yaml:"description"`
	Locale *string `field:"optional" json:"locale" yaml:"locale"`
	Name *string `field:"optional" json:"name" yaml:"name"`
	ParentPool IIpamPool `field:"optional" json:"parentPool" yaml:"parentPool"`
	ProvisionedCidrs *[]*string `field:"optional" json:"provisionedCidrs" yaml:"provisionedCidrs"`
	PublicIpSource PublicIpSource `field:"optional" json:"publicIpSource" yaml:"publicIpSource"`
	TagRestrictions *map[string]*string `field:"optional" json:"tagRestrictions" yaml:"tagRestrictions"`
}

