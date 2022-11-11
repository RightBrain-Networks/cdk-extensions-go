package k8saws

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awseks"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
)

// Optional configuration for the FargateLogger resource.
type FargateLoggerOptions struct {
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
}

