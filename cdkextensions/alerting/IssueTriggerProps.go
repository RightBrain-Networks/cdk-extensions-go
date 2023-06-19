package alerting

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

type IssueTriggerProps struct {
	EventPattern *awsevents.EventPattern `field:"required" json:"eventPattern" yaml:"eventPattern"`
	Parser IIssueParser `field:"required" json:"parser" yaml:"parser"`
	Overrides *[]IssueHandlerOverride `field:"optional" json:"overrides" yaml:"overrides"`
}

