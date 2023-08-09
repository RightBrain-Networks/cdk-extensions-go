package configrules

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsconfig"
)

type IamPasswordPolicyProps struct {
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
	AutoRemediation *bool `field:"optional" json:"autoRemediation" yaml:"autoRemediation"`
	ConfigRuleName *string `field:"optional" json:"configRuleName" yaml:"configRuleName"`
	Description *string `field:"optional" json:"description" yaml:"description"`
	MaximumExecutionFrequency awsconfig.MaximumExecutionFrequency `field:"optional" json:"maximumExecutionFrequency" yaml:"maximumExecutionFrequency"`
	MaxPasswordAge *float64 `field:"optional" json:"maxPasswordAge" yaml:"maxPasswordAge"`
	MinimumPasswordLength *float64 `field:"optional" json:"minimumPasswordLength" yaml:"minimumPasswordLength"`
	PasswordReusePrevention *float64 `field:"optional" json:"passwordReusePrevention" yaml:"passwordReusePrevention"`
	RequireLowercaseCharacters *bool `field:"optional" json:"requireLowercaseCharacters" yaml:"requireLowercaseCharacters"`
	RequireNumbers *bool `field:"optional" json:"requireNumbers" yaml:"requireNumbers"`
	RequireSymbols *bool `field:"optional" json:"requireSymbols" yaml:"requireSymbols"`
	RequireUppercaseCharacters *bool `field:"optional" json:"requireUppercaseCharacters" yaml:"requireUppercaseCharacters"`
}

