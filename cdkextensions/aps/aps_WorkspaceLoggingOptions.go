package aps

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
)

// Logging configuration to use when setting up the APS workspace.
type WorkspaceLoggingOptions struct {
	// Controls whether logging for the workspace should be enabled.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// The log group where events from the APS workspace will be written.
	//
	// If logging is enabled and no log group is provided a new log group will be
	// created.
	// See: [Workspace LoggingConfiguration.LogGroupArn](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-aps-workspace-loggingconfiguration.html#cfn-aps-workspace-loggingconfiguration-loggrouparn)
	//
	LogGroup awslogs.ILogGroup `field:"optional" json:"logGroup" yaml:"logGroup"`
	// The length of time that logs from the APS workspace should be kept when a new log group is created for the workspace.
	//
	// This property is ignored when a log group is passed as part of the logging
	// configuration.
	// See: [LogGroup RetentionInDays](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-loggroup.html#cfn-logs-loggroup-retentionindays)
	//
	Retention awslogs.RetentionDays `field:"optional" json:"retention" yaml:"retention"`
}

