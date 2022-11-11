package k8saws

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// The output of a Fluent Bit configuration object for consumption be the resource configuring Fluent Bit.
type ResolvedFluentBitConfiguration struct {
	// The configuration rended as a configuration file that can be read by the Fluent Bit service.
	ConfigFile *string `field:"required" json:"configFile" yaml:"configFile"`
	// A list of parsers referenced by this plugin.
	Parsers *[]IFluentBitParserPlugin `field:"optional" json:"parsers" yaml:"parsers"`
	// IAM permissions required by resources that will be using this plugin.
	Permissions *[]awsiam.PolicyStatement `field:"optional" json:"permissions" yaml:"permissions"`
}

