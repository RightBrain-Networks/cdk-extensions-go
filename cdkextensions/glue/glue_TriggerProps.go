package glue

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// Configuration for the Glue Trigger resource.
type TriggerProps struct {
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
	//    CloudFormation during deployment.
	// - a concrete value implies a specific physical name
	// - `PhysicalName.GENERATE_IF_NEEDED` is a marker that indicates that a physical will only be generated
	//    by the CDK if it is needed for cross-environment references. Otherwise, it will be allocated by CloudFormation.
	PhysicalName *string `field:"optional" json:"physicalName" yaml:"physicalName"`
	// The AWS region this resource belongs to.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// The type of trigger that this is.
	// See: [AWS::Glue::Trigger](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-trigger.html#cfn-glue-trigger-type)
	//
	Type TriggerType `field:"required" json:"type" yaml:"type"`
	// A list of actions initiated by this trigger.
	// See: [AWS::Glue::Trigger Action](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-trigger-action.html)
	//
	Actions *[]ITriggerAction `field:"optional" json:"actions" yaml:"actions"`
	// A description for the Trigger.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A name for the Trigger.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// A list of conditions predicating the trigger, determining when it will fire.
	// See: [AWS::Glue::Trigger Predicate](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-trigger-predicate.html)
	//
	PredicateConditions *[]ITriggerPredicate `field:"optional" json:"predicateConditions" yaml:"predicateConditions"`
	// Operator for chaining predicate conditions (AND/OR).
	PredicateOperator PredicateOperator `field:"optional" json:"predicateOperator" yaml:"predicateOperator"`
	// A cron expression used to specify the schedule.
	// See: [AWS::Glue::Trigger](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-trigger.html#cfn-glue-trigger-schedule)
	//
	Schedule awsevents.Schedule `field:"optional" json:"schedule" yaml:"schedule"`
	// Set to true to start SCHEDULED and CONDITIONAL triggers when created.
	//
	// True is not supported for ON_DEMAND triggers.
	// See: [AWS::Glue::Trigger](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-trigger.html#cfn-glue-trigger-startoncreation)
	//
	StartOnCreation *bool `field:"optional" json:"startOnCreation" yaml:"startOnCreation"`
	// Workflow object the Trigger should be attached to.
	Workflow Workflow `field:"optional" json:"workflow" yaml:"workflow"`
}

