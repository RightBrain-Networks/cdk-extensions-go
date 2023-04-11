package kinesisfirehose

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awskinesisfirehose"
)

type BackupConfigurationResult struct {
	S3BackupConfiguration *awskinesisfirehose.CfnDeliveryStream_S3DestinationConfigurationProperty `field:"required" json:"s3BackupConfiguration" yaml:"s3BackupConfiguration"`
	S3BackupMode *string `field:"optional" json:"s3BackupMode" yaml:"s3BackupMode"`
}

