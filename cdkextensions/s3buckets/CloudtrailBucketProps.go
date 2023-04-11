package s3buckets

import (
	"github.com/vibe-io/cdk-extensions-go/cdkextensions/glue"
)

// Configuration for objects bucket.
type CloudtrailBucketProps struct {
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
	//   CloudFormation during deployment.
	// - a concrete value implies a specific physical name
	// - `PhysicalName.GENERATE_IF_NEEDED` is a marker that indicates that a physical will only be generated
	//   by the CDK if it is needed for cross-environment references. Otherwise, it will be allocated by CloudFormation.
	PhysicalName *string `field:"optional" json:"physicalName" yaml:"physicalName"`
	// The AWS region this resource belongs to.
	Region *string `field:"optional" json:"region" yaml:"region"`
	BucketName *string `field:"optional" json:"bucketName" yaml:"bucketName"`
	CreateQueries *bool `field:"optional" json:"createQueries" yaml:"createQueries"`
	Database glue.Database `field:"optional" json:"database" yaml:"database"`
	FriendlyQueryNames *bool `field:"optional" json:"friendlyQueryNames" yaml:"friendlyQueryNames"`
	TableName *string `field:"optional" json:"tableName" yaml:"tableName"`
}

