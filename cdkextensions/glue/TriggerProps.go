package glue

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// Configuration for the GlueTrigger resource.
type TriggerProps struct {
	// The AWS account ID this resource belongs to.
	// Default: - the resource is in the same account as the stack it belongs to.
	//
	Account *string `field:"optional" json:"account" yaml:"account"`
	// ARN to deduce region and account from.
	//
	// The ARN is parsed and the account and region are taken from the ARN.
	// This should be used for imported resources.
	//
	// Cannot be supplied together with either `account` or `region`.
	// Default: - take environment from `account`, `region` parameters, or use Stack environment.
	//
	EnvironmentFromArn *string `field:"optional" json:"environmentFromArn" yaml:"environmentFromArn"`
	// The value passed in by users to the physical name prop of the resource.
	//
	// - `undefined` implies that a physical name will be allocated by
	//   CloudFormation during deployment.
	// - a concrete value implies a specific physical name
	// - `PhysicalName.GENERATE_IF_NEEDED` is a marker that indicates that a physical will only be generated
	//   by the CDK if it is needed for cross-environment references. Otherwise, it will be allocated by CloudFormation.
	// Default: - The physical name will be allocated by CloudFormation at deployment time.
	//
	PhysicalName *string `field:"optional" json:"physicalName" yaml:"physicalName"`
	// The AWS region this resource belongs to.
	// Default: - the resource is in the same region as the stack it belongs to.
	//
	Region *string `field:"optional" json:"region" yaml:"region"`
	// The type of trigger that this is.
	// See: [Trigger Type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-trigger.html#cfn-glue-trigger-type)
	//
	Type TriggerType `field:"required" json:"type" yaml:"type"`
	// A list of actions initiated by this trigger.
	// See: [Trigger Actions](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-trigger.html#cfn-glue-trigger-actions)
	//
	Actions *[]ITriggerAction `field:"optional" json:"actions" yaml:"actions"`
	// A description for the trigger.
	// See: [Trigger Description](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-trigger.html#cfn-glue-trigger-description)
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A name for the trigger.
	// See: [Trigger Name](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-trigger.html#cfn-glue-trigger-name)
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// A list of the conditions that determine when the trigger will fire.
	// See: [Trigger Predicate](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-trigger-predicate.html)
	//
	PredicateConditions *[]ITriggerPredicate `field:"optional" json:"predicateConditions" yaml:"predicateConditions"`
	// Operator for chaining predicate conditions if multiple are given.
	// See: [Trigger Predicate.Logical](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-trigger-predicate.html#cfn-glue-trigger-predicate-logical)
	//
	PredicateOperator PredicateOperator `field:"optional" json:"predicateOperator" yaml:"predicateOperator"`
	// A cron expression used to specify the schedule.
	// See: [Time-Based Schedules for Jobs and Crawlers](https://docs.aws.amazon.com/glue/latest/dg/monitor-data-warehouse-schedule.html)
	//
	Schedule awsevents.Schedule `field:"optional" json:"schedule" yaml:"schedule"`
	// Set to true to start SCHEDULED and CONDITIONAL triggers when created.
	//
	// True
	// is not supported for ON_DEMAND triggers.
	// See: [Trigger StartOnCreation](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-trigger.html#cfn-glue-trigger-startoncreation)
	//
	StartOnCreation *bool `field:"optional" json:"startOnCreation" yaml:"startOnCreation"`
	// The name of the workflow associated with the trigger.
	// See: [Trigger WorkflowName](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-trigger.html#cfn-glue-trigger-workflowname)
	//
	Workflow Workflow `field:"optional" json:"workflow" yaml:"workflow"`
}

