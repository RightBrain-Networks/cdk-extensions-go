package ec2

import (
	"github.com/aws/constructs-go/constructs/v10"
)

type AllocateCidrFromPoolOptions struct {
	Allocation IIpamAllocationConfiguration `field:"optional" json:"allocation" yaml:"allocation"`
	Description *string `field:"optional" json:"description" yaml:"description"`
	Scope constructs.IConstruct `field:"optional" json:"scope" yaml:"scope"`
}

