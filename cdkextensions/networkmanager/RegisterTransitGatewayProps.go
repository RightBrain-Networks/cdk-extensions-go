package networkmanager

import (
	"github.com/vibe-io/cdk-extensions-go/cdkextensions/ec2"
)

type RegisterTransitGatewayProps struct {
	TransitGateway ec2.ITransitGateway `field:"required" json:"transitGateway" yaml:"transitGateway"`
}

