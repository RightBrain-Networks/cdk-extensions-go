package k8saws

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awseks"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
)

// Required configuration for the Fargate logger resource.
type FargateLoggerProps struct {
	// A default list of Fargate profiles that should have permissions configured.
	//
	// Alternatively profiles can be added at any time by calling
	// `addProfile`.
	FargateProfiles *[]awseks.FargateProfile `field:"optional" json:"fargateProfiles" yaml:"fargateProfiles"`
	// The filters that should be applied to logs being processed.
	Filters *[]IFluentBitFilterPlugin `field:"optional" json:"filters" yaml:"filters"`
	// The CloudWatch log group where Farget container logs will be sent.
	LogGroup awslogs.ILogGroup `field:"optional" json:"logGroup" yaml:"logGroup"`
	// The output destinations where logs should be written.
	Outputs *[]IFluentBitOutputPlugin `field:"optional" json:"outputs" yaml:"outputs"`
	// The parsers to be used when reading log files.
	Parsers *[]IFluentBitParserPlugin `field:"optional" json:"parsers" yaml:"parsers"`
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
	// The EKS Cluster to configure Fargate logging for.
	Cluster awseks.ICluster `field:"required" json:"cluster" yaml:"cluster"`
}

