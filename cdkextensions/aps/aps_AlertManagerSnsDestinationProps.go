package aps

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
)

// Configuration for an alert manager SNS destination.
type AlertManagerSnsDestinationProps struct {
	// The SNS API URL i.e. https://sns.us-east-2.amazonaws.com.
	//
	// If not specified, the SNS API URL from the SNS SDK will be used.
	ApiUrl *string `field:"optional" json:"apiUrl" yaml:"apiUrl"`
	// SNS message attributes.
	Attributes *map[string]*string `field:"optional" json:"attributes" yaml:"attributes"`
	// The message content of the SNS notification.
	Message *string `field:"optional" json:"message" yaml:"message"`
	// Controls whether to notify about resolved alerts.
	SendResolved *bool `field:"optional" json:"sendResolved" yaml:"sendResolved"`
	// Subject line when the message is delivered to email endpoints.
	Subject *string `field:"optional" json:"subject" yaml:"subject"`
	// SNS topic where alerts will be sent.
	//
	// If you are using a FIFO SNS topic you should set a message group interval
	// longer than 5 minutes to prevent messages with the same group key being
	// deduplicated by the SNS default deduplication window.
	Topic awssns.ITopic `field:"required" json:"topic" yaml:"topic"`
}

