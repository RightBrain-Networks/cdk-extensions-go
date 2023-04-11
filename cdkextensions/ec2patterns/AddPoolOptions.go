package ec2patterns


type AddPoolOptions struct {
	Cidrs *[]*string `field:"optional" json:"cidrs" yaml:"cidrs"`
	DefaultNetmaskLength *float64 `field:"optional" json:"defaultNetmaskLength" yaml:"defaultNetmaskLength"`
}

