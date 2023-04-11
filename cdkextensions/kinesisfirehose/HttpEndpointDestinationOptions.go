package kinesisfirehose

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

type HttpEndpointDestinationOptions struct {
	AccessKey awscdk.SecretValue `field:"optional" json:"accessKey" yaml:"accessKey"`
	BackupConfiguration BackupConfiguration `field:"optional" json:"backupConfiguration" yaml:"backupConfiguration"`
	Buffering BufferingConfiguration `field:"optional" json:"buffering" yaml:"buffering"`
	CloudwatchLoggingConfiguration CloudWatchLoggingConfiguration `field:"optional" json:"cloudwatchLoggingConfiguration" yaml:"cloudwatchLoggingConfiguration"`
	CommonAttributes *map[string]*string `field:"optional" json:"commonAttributes" yaml:"commonAttributes"`
	ContentEncoding ContentEncoding `field:"optional" json:"contentEncoding" yaml:"contentEncoding"`
	EndpointName *string `field:"optional" json:"endpointName" yaml:"endpointName"`
	ProcessorConfiguration ProcessorConfiguration `field:"optional" json:"processorConfiguration" yaml:"processorConfiguration"`
	RetryDuration awscdk.Duration `field:"optional" json:"retryDuration" yaml:"retryDuration"`
}

