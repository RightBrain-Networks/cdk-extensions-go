package ec2patterns


type AllocatePrivateNetworkOptions struct {
	Netmask *float64 `field:"optional" json:"netmask" yaml:"netmask"`
	Pool *string `field:"optional" json:"pool" yaml:"pool"`
}

