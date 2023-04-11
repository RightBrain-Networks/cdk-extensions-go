package ec2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

type TunnelOptions struct {
	InsideCidr *string `field:"optional" json:"insideCidr" yaml:"insideCidr"`
	PreSharedKey awscdk.SecretValue `field:"optional" json:"preSharedKey" yaml:"preSharedKey"`
}

