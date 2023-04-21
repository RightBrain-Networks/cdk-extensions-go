package ec2patterns

import (
	"github.com/vibe-io/cdk-extensions-go/cdkextensions/ec2"
)

type AddHubOptions struct {
	AvailabilityZones *[]*string `field:"optional" json:"availabilityZones" yaml:"availabilityZones"`
	MaxAzs *float64 `field:"optional" json:"maxAzs" yaml:"maxAzs"`
	Netmask *float64 `field:"optional" json:"netmask" yaml:"netmask"`
	DefaultTransitGatewayRouteTable ec2.ITransitGatewayRouteTable `field:"optional" json:"defaultTransitGatewayRouteTable" yaml:"defaultTransitGatewayRouteTable"`
}

