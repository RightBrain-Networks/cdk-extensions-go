package ec2patterns

import (
	"github.com/aws/constructs-go/constructs/v10"
)

type AddMultiSubnetRouteOptions struct {
	Cidr *string `field:"required" json:"cidr" yaml:"cidr"`
	Description *string `field:"optional" json:"description" yaml:"description"`
	Scope constructs.IConstruct `field:"optional" json:"scope" yaml:"scope"`
}

