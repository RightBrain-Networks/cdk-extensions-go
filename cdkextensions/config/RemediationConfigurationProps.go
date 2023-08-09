package config

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsconfig"
)

type RemediationConfigurationProps struct {
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
	ConfigRule awsconfig.IRule `field:"required" json:"configRule" yaml:"configRule"`
	StaticParameters *map[string]*[]interface{} `field:"required" json:"staticParameters" yaml:"staticParameters"`
	Target IRemediationTarget `field:"required" json:"target" yaml:"target"`
	Automatic *bool `field:"optional" json:"automatic" yaml:"automatic"`
	MaximumAutomaticAttempts *float64 `field:"optional" json:"maximumAutomaticAttempts" yaml:"maximumAutomaticAttempts"`
	ResourceParameter *string `field:"optional" json:"resourceParameter" yaml:"resourceParameter"`
	ResourceType *string `field:"optional" json:"resourceType" yaml:"resourceType"`
	RetryInterval awscdk.Duration `field:"optional" json:"retryInterval" yaml:"retryInterval"`
}

