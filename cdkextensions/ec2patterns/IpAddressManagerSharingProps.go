package ec2patterns


type IpAddressManagerSharingProps struct {
	AllowExternalPricipals *bool `field:"optional" json:"allowExternalPricipals" yaml:"allowExternalPricipals"`
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
}

