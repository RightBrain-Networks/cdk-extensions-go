package ec2


type Ipv4ConfigurationOptions struct {
	DefaultNetmaskLength *float64 `field:"optional" json:"defaultNetmaskLength" yaml:"defaultNetmaskLength"`
	MaxNetmaskLength *float64 `field:"optional" json:"maxNetmaskLength" yaml:"maxNetmaskLength"`
	MinNetmaskLength *float64 `field:"optional" json:"minNetmaskLength" yaml:"minNetmaskLength"`
}

