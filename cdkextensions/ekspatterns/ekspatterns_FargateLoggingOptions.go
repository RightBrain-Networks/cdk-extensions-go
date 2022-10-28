package ekspatterns

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
)

type FargateLoggingOptions struct {
	// Controls whether logging will be set up for pods using the default Fargate provide on the EKS cluster.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// The CloudWatch log group where Farget container logs will be sent.
	LogGroup awslogs.ILogGroup `field:"optional" json:"logGroup" yaml:"logGroup"`
	// The prefix to add to the start of log streams created by the Fargate logger.
	LogStreamPrefix *string `field:"optional" json:"logStreamPrefix" yaml:"logStreamPrefix"`
	// The number of days logs sent to CloudWatch from Fluent Bit should be retained before they are automatically removed.
	Retention awslogs.RetentionDays `field:"optional" json:"retention" yaml:"retention"`
}

