package aps

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
)

// Alerting configuration to use when setting up an APS workspace.
type WorkspaceAlertingOptions struct {
	// The alert manager configuration to use when setting up alerting.
	//
	// If alerting is enabled and no configuration is given a default
	// configuration that sends all alerts to SNS will be used.
	Configuration IAlertManagerConfiguration `field:"optional" json:"configuration" yaml:"configuration"`
	// Controls whether alerting from the APS workspace should be configured.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// The SNS topic where alerts should be sent when using the default alerting configuration.
	//
	// If a custom alert manager configuration is provided this option is ignored.
	Topic awssns.ITopic `field:"optional" json:"topic" yaml:"topic"`
}

