package aps


// Optional configuration for an alert manager SNS destination.
type AlertManagerSnsDestinationOptions struct {
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
}

