package kinesisfirehose

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

type DeliveryStreamAttributes struct {
	DeliveryStreamArn *string `field:"optional" json:"deliveryStreamArn" yaml:"deliveryStreamArn"`
	DeliveryStreamName *string `field:"optional" json:"deliveryStreamName" yaml:"deliveryStreamName"`
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
}

