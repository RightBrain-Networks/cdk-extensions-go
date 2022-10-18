package kinesisfirehose

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

type BufferingConfigurationOptions struct {
	Interval awscdk.Duration `field:"optional" json:"interval" yaml:"interval"`
	SizeInMb *float64 `field:"optional" json:"sizeInMb" yaml:"sizeInMb"`
}

