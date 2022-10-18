package glue

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

// Configuration for the Glue Workflow resource.
type ConnectionProps struct {
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
	//    CloudFormation during deployment.
	// - a concrete value implies a specific physical name
	// - `PhysicalName.GENERATE_IF_NEEDED` is a marker that indicates that a physical will only be generated
	//    by the CDK if it is needed for cross-environment references. Otherwise, it will be allocated by CloudFormation.
	PhysicalName *string `field:"optional" json:"physicalName" yaml:"physicalName"`
	// The AWS region this resource belongs to.
	Region *string `field:"optional" json:"region" yaml:"region"`
	ConnectionType ConnectionType `field:"required" json:"connectionType" yaml:"connectionType"`
	Description *string `field:"optional" json:"description" yaml:"description"`
	Name *string `field:"optional" json:"name" yaml:"name"`
	Properties *map[string]*string `field:"optional" json:"properties" yaml:"properties"`
	SecurityGroups *[]awsec2.ISecurityGroup `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	Subnets *awsec2.SubnetSelection `field:"optional" json:"subnets" yaml:"subnets"`
	Vpc awsec2.IVpc `field:"optional" json:"vpc" yaml:"vpc"`
}

