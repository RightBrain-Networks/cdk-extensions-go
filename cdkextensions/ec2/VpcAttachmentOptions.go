package ec2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

type VpcAttachmentOptions struct {
	Name *string `field:"optional" json:"name" yaml:"name"`
	Subnets *awsec2.SubnetSelection `field:"optional" json:"subnets" yaml:"subnets"`
}

