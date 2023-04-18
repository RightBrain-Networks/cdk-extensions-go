package ec2patterns


type AddNetworkOptions struct {
	AvailabilityZones *[]*string `field:"optional" json:"availabilityZones" yaml:"availabilityZones"`
	MaxAzs *float64 `field:"optional" json:"maxAzs" yaml:"maxAzs"`
	Netmask *float64 `field:"optional" json:"netmask" yaml:"netmask"`
}

