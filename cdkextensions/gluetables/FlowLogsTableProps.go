package gluetables

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/vibe-io/cdk-extensions-go/cdkextensions/athena"
	"github.com/vibe-io/cdk-extensions-go/cdkextensions/ec2"
	"github.com/vibe-io/cdk-extensions-go/cdkextensions/glue"
)

// Configuration for FlowLogsTable.
type FlowLogsTableProps struct {
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
	// A bucket where logs will be stored.
	// See: [AWS S3 iBucket](https://docs.aws.amazon.com/cdk/api/v2/docs/aws-cdk-lib.aws_s3.IBucket.html)
	//
	Bucket awss3.IBucket `field:"required" json:"bucket" yaml:"bucket"`
	// A cdk-extensions/glue {@link aws-glue !Database } object that the table should be created in.
	// See: [AWS::Glue::Database](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-database.html)
	//
	Database glue.Database `field:"required" json:"database" yaml:"database"`
	// Boolean indicating whether to create default Athena queries for the Flow Logs.
	// See: [`CfnNamedQueries`](https://docs.aws.amazon.com/cdk/api/v1/python/aws_cdk.aws_athena/CfnNamedQuery.html)
	//
	CreateQueries *bool `field:"optional" json:"createQueries" yaml:"createQueries"`
	// A cdk-extentions/ec2 {@link aws-ec2 !FlowLogFormat } object defining the desired formatting for Flow Logs.
	Format ec2.FlowLogFormat `field:"optional" json:"format" yaml:"format"`
	// Boolean for adding "friendly names" for the created Athena queries.
	FriendlyQueryNames *bool `field:"optional" json:"friendlyQueryNames" yaml:"friendlyQueryNames"`
	// Name for Flow Logs Table.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Set a custom prefix for the S3 Bucket.
	S3Prefix *string `field:"optional" json:"s3Prefix" yaml:"s3Prefix"`
	// The name of the workgroup where namedqueries should be created.
	// See: [Setting up workgroups](https://docs.aws.amazon.com/athena/latest/ug/workgroups-procedure.html)
	//
	WorkGroup athena.IWorkGroup `field:"optional" json:"workGroup" yaml:"workGroup"`
}

