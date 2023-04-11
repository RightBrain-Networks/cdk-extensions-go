package glue

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Configuration options for the WorkflowCrawlerAction class.
type WorkflowCrawlerActionOptions struct {
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
	// The arguments to use when the associated trigger fires.
	//
	// Jobs run via the associated trigger will have their default arguments
	// replaced with the arguments specified.
	//
	// You can specify arguments here that your own job-execution script
	// consumes, in addition to arguments that AWS Glue itself consumes.
	// See: [Trigger Actions.Arguments](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-trigger-action.html#cfn-glue-trigger-action-arguments)
	//
	Arguments *map[string]*string `field:"optional" json:"arguments" yaml:"arguments"`
	// After a job run starts, the number of minutes to wait before sending a job run delay notification.
	// See: [Trigger Actions.NotificationProperty.NotifyDelayAfter](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-trigger-notificationproperty.html#cfn-glue-trigger-notificationproperty-notifydelayafter)
	//
	NotifyDelayAfter awscdk.Duration `field:"optional" json:"notifyDelayAfter" yaml:"notifyDelayAfter"`
	// The name of the SecurityConfiguration structure to be used with this action.
	// See: [Trigger Actions.SecurityConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-trigger-action.html#cfn-glue-trigger-action-securityconfiguration)
	//
	SecurityConfiguration *string `field:"optional" json:"securityConfiguration" yaml:"securityConfiguration"`
	// The `JobRun` timeout in minutes.
	//
	// This is the maximum time that a job run
	// can consume resources before it is terminated and enters TIMEOUT status.
	// The default is 48 hours. This overrides the timeout value set in the
	// parent job.
	// See: [Trigger Actions.Timeout](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-trigger-action.html#cfn-glue-trigger-action-timeout)
	//
	Timeout awscdk.Duration `field:"optional" json:"timeout" yaml:"timeout"`
}

