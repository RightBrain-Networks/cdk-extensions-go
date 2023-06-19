package alerting

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
)

type StateMachineLogging struct {
	Destination awslogs.ILogGroup `field:"optional" json:"destination" yaml:"destination"`
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	IncludeExecutionData *bool `field:"optional" json:"includeExecutionData" yaml:"includeExecutionData"`
	Level awsstepfunctions.LogLevel `field:"optional" json:"level" yaml:"level"`
}

