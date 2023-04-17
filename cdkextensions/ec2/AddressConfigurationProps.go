package ec2


type AddressConfigurationProps struct {
	DefaultNetmaskLength *float64 `field:"optional" json:"defaultNetmaskLength" yaml:"defaultNetmaskLength"`
	MaxNetmaskLength *float64 `field:"optional" json:"maxNetmaskLength" yaml:"maxNetmaskLength"`
	MinNetmaskLength *float64 `field:"optional" json:"minNetmaskLength" yaml:"minNetmaskLength"`
	Family IpFamily `field:"required" json:"family" yaml:"family"`
	AdvertiseService AdvertiseService `field:"optional" json:"advertiseService" yaml:"advertiseService"`
	PubliclyAdvertisable *bool `field:"optional" json:"publiclyAdvertisable" yaml:"publiclyAdvertisable"`
}

