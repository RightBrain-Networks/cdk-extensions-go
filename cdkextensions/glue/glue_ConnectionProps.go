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
	// The type of the connection.
	// See: [AWS::Glue::Connection ConnectionInput](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connection-connectioninput.html#cfn-glue-connection-connectioninput-connectiontype)
	//
	ConnectionType ConnectionType `field:"required" json:"connectionType" yaml:"connectionType"`
	// A description for the Connection.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the connection.
	//
	// Connection will not function as expected without a name.
	// See: [AWS::Glue::Connection ConnectionInput](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connection-connectioninput.html#cfn-glue-connection-connectioninput-name)
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// List of Key/Value pairs defining the properties of the Connection.
	// See: [AWS::Glue::Connection Properties](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-connection.html#Properties)
	//
	Properties *map[string]*string `field:"optional" json:"properties" yaml:"properties"`
	// Existing Security Group to assign to the Connection.
	//
	// If none is provided a new Security Group will be created.
	SecurityGroups *[]awsec2.ISecurityGroup `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// Options for selection of subnets from the VPC to attach to the Connection.
	// See: [CDK SubnetSelection](https://docs.aws.amazon.com/cdk/api/v2/docs/aws-cdk-lib.aws_ec2.SubnetSelection.html)
	//
	Subnets *awsec2.SubnetSelection `field:"optional" json:"subnets" yaml:"subnets"`
	// VPC to attach to the Connection.
	// See: [IVpc Interface](https://docs.aws.amazon.com/cdk/api/v2/docs/aws-cdk-lib.aws_ec2.IVpc.html)
	//
	Vpc awsec2.IVpc `field:"optional" json:"vpc" yaml:"vpc"`
}

