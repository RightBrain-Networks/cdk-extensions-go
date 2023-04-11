package kinesisfirehose

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

type CommonPartitioningOptions struct {
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	RetryInterval awscdk.Duration `field:"optional" json:"retryInterval" yaml:"retryInterval"`
}

