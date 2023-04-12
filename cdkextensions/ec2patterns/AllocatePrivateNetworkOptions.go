package ec2patterns

import (
	"github.com/vibe-io/cdk-extensions-go/cdkextensions/ec2"
)

type AllocatePrivateNetworkOptions struct {
	Consumer ec2.IpamConsumer `field:"optional" json:"consumer" yaml:"consumer"`
	Netmask *float64 `field:"optional" json:"netmask" yaml:"netmask"`
	Pool *string `field:"optional" json:"pool" yaml:"pool"`
}

