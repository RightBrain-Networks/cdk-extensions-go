package ekspatterns

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awseks"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/vibe-io/cdk-extensions-go/cdkextensions/k8saws"
)

// Configuration options for enabling persistent logging for Fargate containers on the cluster.
type ClusterFargateLoggingOptions struct {
	// A default list of Fargate profiles that should have permissions configured.
	//
	// Alternatively profiles can be added at any time by calling
	// `addProfile`.
	FargateProfiles *[]awseks.FargateProfile `field:"optional" json:"fargateProfiles" yaml:"fargateProfiles"`
	// The filters that should be applied to logs being processed.
	Filters *[]k8saws.IFluentBitFilterPlugin `field:"optional" json:"filters" yaml:"filters"`
	// The CloudWatch log group where Farget container logs will be sent.
	LogGroup awslogs.ILogGroup `field:"optional" json:"logGroup" yaml:"logGroup"`
	// The output destinations where logs should be written.
	Outputs *[]k8saws.IFluentBitOutputPlugin `field:"optional" json:"outputs" yaml:"outputs"`
	// The parsers to be used when reading log files.
	Parsers *[]k8saws.IFluentBitParserPlugin `field:"optional" json:"parsers" yaml:"parsers"`
	// Controls whether logging will be set up for pods using the default Fargate provide on the EKS cluster.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
}

