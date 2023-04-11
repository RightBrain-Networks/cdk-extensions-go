package kinesisfirehose

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

type ExtendedS3DestinationOptions struct {
	Buffering BufferingConfiguration `field:"optional" json:"buffering" yaml:"buffering"`
	CloudwatchLoggingConfiguration CloudWatchLoggingConfiguration `field:"optional" json:"cloudwatchLoggingConfiguration" yaml:"cloudwatchLoggingConfiguration"`
	CompressionFormat S3CompressionFormat `field:"optional" json:"compressionFormat" yaml:"compressionFormat"`
	EncryptionEnabled *bool `field:"optional" json:"encryptionEnabled" yaml:"encryptionEnabled"`
	EncryptionKey awskms.IKey `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	ErrorOutputPrefix *string `field:"optional" json:"errorOutputPrefix" yaml:"errorOutputPrefix"`
	KeyPrefix *string `field:"optional" json:"keyPrefix" yaml:"keyPrefix"`
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	BackupConfiguration BackupConfiguration `field:"optional" json:"backupConfiguration" yaml:"backupConfiguration"`
	DataFormatConversion DataFormatConversion `field:"optional" json:"dataFormatConversion" yaml:"dataFormatConversion"`
	DynamicPartitioning DynamicPartitioning `field:"optional" json:"dynamicPartitioning" yaml:"dynamicPartitioning"`
	ProcessorConfiguration ProcessorConfiguration `field:"optional" json:"processorConfiguration" yaml:"processorConfiguration"`
}

