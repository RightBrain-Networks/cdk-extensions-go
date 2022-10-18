package kinesisfirehose

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awskinesisfirehose"
)

type DynamicPartitioningConfiguration struct {
	PartitioningConfiguration *awskinesisfirehose.CfnDeliveryStream_DynamicPartitioningConfigurationProperty `field:"required" json:"partitioningConfiguration" yaml:"partitioningConfiguration"`
	Processors *[]DeliveryStreamProcessor `field:"optional" json:"processors" yaml:"processors"`
}

