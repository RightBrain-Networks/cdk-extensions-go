package glue

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
)

type ContinuousLoggingProps struct {
	// Enable continouous logging.
	Enabled *bool `field:"required" json:"enabled" yaml:"enabled"`
	// Apply the provided conversion pattern.
	//
	// This is a Log4j Conversion Pattern to customize driver and executor logs.
	ConversionPattern *string `field:"optional" json:"conversionPattern" yaml:"conversionPattern"`
	// Specify a custom CloudWatch log group name.
	LogGroup awslogs.ILogGroup `field:"optional" json:"logGroup" yaml:"logGroup"`
	// Specify a custom CloudWatch log stream prefix.
	LogStreamPrefix *string `field:"optional" json:"logStreamPrefix" yaml:"logStreamPrefix"`
	// Filter out non-useful Apache Spark driver/executor and Apache Hadoop YARN heartbeat log messages.
	Quiet *bool `field:"optional" json:"quiet" yaml:"quiet"`
}

