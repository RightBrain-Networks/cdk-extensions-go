package kinesisfirehose

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
)

type CloudWatchLoggingConfigurationOptions struct {
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	LogGroup awslogs.ILogGroup `field:"optional" json:"logGroup" yaml:"logGroup"`
	LogStream awslogs.ILogStream `field:"optional" json:"logStream" yaml:"logStream"`
}

