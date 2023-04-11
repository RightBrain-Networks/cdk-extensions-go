package ec2


type IpamResourceDiscoveryAttributes struct {
	IsDefault *bool `field:"optional" json:"isDefault" yaml:"isDefault"`
	OwnerId *string `field:"optional" json:"ownerId" yaml:"ownerId"`
	Region *string `field:"optional" json:"region" yaml:"region"`
	ResourceDiscoveryArn *string `field:"optional" json:"resourceDiscoveryArn" yaml:"resourceDiscoveryArn"`
	ResourceDiscoveryId *string `field:"optional" json:"resourceDiscoveryId" yaml:"resourceDiscoveryId"`
	State *string `field:"optional" json:"state" yaml:"state"`
}

