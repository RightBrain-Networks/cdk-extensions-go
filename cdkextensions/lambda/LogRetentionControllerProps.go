package lambda

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
)

// Configuration for the LogRetentionController resource.
type LogRetentionControllerProps struct {
	// The AWS account ID this resource belongs to.
	Account *string `field:"optional" json:"account" yaml:"account"`
	// ARN to deduce region and account from.
	//
	// The ARN is parsed and the account and region are taken from the ARN.
	// This should be used for imported resources.
	//
	// Cannot be supplied together with either `account` or `region`.
	EnvironmentFromArn *string `field:"optional" json:"environmentFromArn" yaml:"environmentFromArn"`
	// The value passed in by users to the physical name prop of the resource.
	//
	// - `undefined` implies that a physical name will be allocated by
	//   CloudFormation during deployment.
	// - a concrete value implies a specific physical name
	// - `PhysicalName.GENERATE_IF_NEEDED` is a marker that indicates that a physical will only be generated
	//   by the CDK if it is needed for cross-environment references. Otherwise, it will be allocated by CloudFormation.
	PhysicalName *string `field:"optional" json:"physicalName" yaml:"physicalName"`
	// The AWS region this resource belongs to.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Execution logging configuration for the state machine that is used to configure log retention for log groups created via AWS Lambda.
	// See: [StateMachine LoggingConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-stepfunctions-statemachine.html#cfn-stepfunctions-statemachine-loggingconfiguration)
	//
	ExecutionLogging *ExecutionLogOptions `field:"optional" json:"executionLogging" yaml:"executionLogging"`
	// The length of time logs sent to log groups created by AWS Lambda should be retained before being deleted.
	// See: [LogGroup RetentionInDays](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-loggroup.html#cfn-logs-loggroup-retentionindays)
	//
	Retention awslogs.RetentionDays `field:"optional" json:"retention" yaml:"retention"`
}

