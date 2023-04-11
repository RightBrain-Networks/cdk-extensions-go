package lambda

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
)

// Options for configuring logging from an executing state machine.
type ExecutionLogOptions struct {
	// Determines whether execution data is included in your log.
	//
	// When set to
	// `false`, data is excluded.
	// See: [StateMachine LoggingConfiguration.IncludeExecutionData](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-stepfunctions-statemachine-loggingconfiguration.html#cfn-stepfunctions-statemachine-loggingconfiguration-includeexecutiondata)
	//
	IncludeExecutionData *bool `field:"required" json:"includeExecutionData" yaml:"includeExecutionData"`
	// Defines which category of execution history events are logged.
	// See: [StateMachine LoggingConfiguration.Level](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-stepfunctions-statemachine-loggingconfiguration.html#cfn-stepfunctions-statemachine-loggingconfiguration-level)
	//
	Level awsstepfunctions.LogLevel `field:"required" json:"level" yaml:"level"`
	// Controls whether logging from the state machine is enabled.
	// See: [StateMachine LoggingConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-stepfunctions-statemachine.html#cfn-stepfunctions-statemachine-loggingconfiguration)
	//
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// Specifies a log group which will receive execution events from the state machine.
	//
	// If no log group is passed and loggin is enabled, a log group will be
	// created automatically.
	// See: [StateMachine LoggingConfiguration.Destinations.CloudWatchLogsLogGroup.LogGroupArn](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-stepfunctions-statemachine-cloudwatchlogsloggroup.html#cfn-stepfunctions-statemachine-cloudwatchlogsloggroup-loggrouparn)
	//
	LogGroup awslogs.ILogGroup `field:"optional" json:"logGroup" yaml:"logGroup"`
	// The number of days execution logging events should be retained before being deleted.
	//
	// This value is ignored if `logGroup` is passed.
	// See: [LogGroup RetentionInDays](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-loggroup.html#cfn-logs-loggroup-retentionindays)
	//
	Retention awslogs.RetentionDays `field:"optional" json:"retention" yaml:"retention"`
}

