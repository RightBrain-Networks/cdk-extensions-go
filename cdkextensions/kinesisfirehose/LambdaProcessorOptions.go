package kinesisfirehose

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
)

type LambdaProcessorOptions struct {
	LambdaFunction awslambda.IFunction `field:"required" json:"lambdaFunction" yaml:"lambdaFunction"`
}

