package config

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsconfig"
)

type RemediationTargetConfiguration struct {
	TargetId *string `field:"required" json:"targetId" yaml:"targetId"`
	TargetType RemediationTargetType `field:"required" json:"targetType" yaml:"targetType"`
	Controls *awsconfig.CfnRemediationConfiguration_ExecutionControlsProperty `field:"optional" json:"controls" yaml:"controls"`
	TargetVersion *string `field:"optional" json:"targetVersion" yaml:"targetVersion"`
}

