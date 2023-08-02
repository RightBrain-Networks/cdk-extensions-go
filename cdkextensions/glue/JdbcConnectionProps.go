package glue

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

// Configuration for the Glue Workflow resource.
type JdbcConnectionProps struct {
	// The AWS account ID this resource belongs to.
	// Default: - the resource is in the same account as the stack it belongs to.
	//
	Account *string `field:"optional" json:"account" yaml:"account"`
	// ARN to deduce region and account from.
	//
	// The ARN is parsed and the account and region are taken from the ARN.
	// This should be used for imported resources.
	//
	// Cannot be supplied together with either `account` or `region`.
	// Default: - take environment from `account`, `region` parameters, or use Stack environment.
	//
	EnvironmentFromArn *string `field:"optional" json:"environmentFromArn" yaml:"environmentFromArn"`
	// The value passed in by users to the physical name prop of the resource.
	//
	// - `undefined` implies that a physical name will be allocated by
	//   CloudFormation during deployment.
	// - a concrete value implies a specific physical name
	// - `PhysicalName.GENERATE_IF_NEEDED` is a marker that indicates that a physical will only be generated
	//   by the CDK if it is needed for cross-environment references. Otherwise, it will be allocated by CloudFormation.
	// Default: - The physical name will be allocated by CloudFormation at deployment time.
	//
	PhysicalName *string `field:"optional" json:"physicalName" yaml:"physicalName"`
	// The AWS region this resource belongs to.
	// Default: - the resource is in the same region as the stack it belongs to.
	//
	Region *string `field:"optional" json:"region" yaml:"region"`
	// A SecretValue providing the password for the Connection to authenticate to the source with.
	// See: [AWS::Glue::Connection ConnectionInput](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connection-connectioninput.html#cfn-glue-connection-connectioninput-connectiontype)
	//
	Password awscdk.SecretValue `field:"required" json:"password" yaml:"password"`
	// The URL to the source for the Connection.
	// See: [AWS::Glue::Connection ConnectionInput](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connection-connectioninput.html#cfn-glue-connection-connectioninput-connectiontype)
	//
	Url *string `field:"required" json:"url" yaml:"url"`
	// The username for the Connection to authenticate to the source with.
	// See: [AWS::Glue::Connection ConnectionInput](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connection-connectioninput.html#cfn-glue-connection-connectioninput-connectiontype)
	//
	Username *string `field:"required" json:"username" yaml:"username"`
	// VPC to attach to the Connection.
	// See: [IVpc Interface](https://docs.aws.amazon.com/cdk/api/v2/docs/aws-cdk-lib.aws_ec2.IVpc.html)
	//
	Vpc awsec2.IVpc `field:"required" json:"vpc" yaml:"vpc"`
	// A description of the Connection.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Boolean value on whether to require encryption on the Connection.
	// See: [AWS::Glue::Connection ConnectionInput](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connection-connectioninput.html#cfn-glue-connection-connectioninput-connectiontype)
	//
	EnforceSsl *bool `field:"optional" json:"enforceSsl" yaml:"enforceSsl"`
	// A name for the Connection.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// A list of Security Groups to apply to the Connection.
	SecurityGroups *[]awsec2.ISecurityGroup `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// Options for selection of subnets from the VPC to attach to the Connection.
	// See: [CDK SubnetSelection](https://docs.aws.amazon.com/cdk/api/v2/docs/aws-cdk-lib.aws_ec2.SubnetSelection.html)
	//
	Subnets *awsec2.SubnetSelection `field:"optional" json:"subnets" yaml:"subnets"`
}

