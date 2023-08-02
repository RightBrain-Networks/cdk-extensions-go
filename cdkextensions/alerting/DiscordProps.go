package alerting

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager"
)

// Configuration controlling how Discord messages should be sent in response to events.
type DiscordProps struct {
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
	Logging *StateMachineLogging `field:"optional" json:"logging" yaml:"logging"`
	Channel *string `field:"required" json:"channel" yaml:"channel"`
	Token awssecretsmanager.ISecret `field:"required" json:"token" yaml:"token"`
	EventBus awsevents.IEventBus `field:"optional" json:"eventBus" yaml:"eventBus"`
	Mentions *[]*string `field:"optional" json:"mentions" yaml:"mentions"`
	Name *string `field:"optional" json:"name" yaml:"name"`
	Timeout awscdk.Duration `field:"optional" json:"timeout" yaml:"timeout"`
}

