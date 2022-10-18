package kinesisfirehose

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
)

type LambdaPartitioningOptions struct {
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	RetryInterval awscdk.Duration `field:"optional" json:"retryInterval" yaml:"retryInterval"`
	LambdaFunction awslambda.IFunction `field:"required" json:"lambdaFunction" yaml:"lambdaFunction"`
}

