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
	// Default: `%d{yy/MM/dd HH:mm:ss} %p %c{1}: %m%n`.
	//
	ConversionPattern *string `field:"optional" json:"conversionPattern" yaml:"conversionPattern"`
	// Specify a custom CloudWatch log group name.
	// Default: - a log group is created with name `/aws-glue/jobs/logs-v2/`.
	//
	LogGroup awslogs.ILogGroup `field:"optional" json:"logGroup" yaml:"logGroup"`
	// Specify a custom CloudWatch log stream prefix.
	// Default: - the job run ID.
	//
	LogStreamPrefix *string `field:"optional" json:"logStreamPrefix" yaml:"logStreamPrefix"`
	// Filter out non-useful Apache Spark driver/executor and Apache Hadoop YARN heartbeat log messages.
	// Default: true.
	//
	Quiet *bool `field:"optional" json:"quiet" yaml:"quiet"`
}

