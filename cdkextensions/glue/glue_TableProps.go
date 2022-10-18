package glue

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Configuration for Table.
type TableProps struct {
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
	Database Database `field:"required" json:"database" yaml:"database"`
	Columns *[]Column `field:"optional" json:"columns" yaml:"columns"`
	Compressed *bool `field:"optional" json:"compressed" yaml:"compressed"`
	DataFormat DataFormat `field:"optional" json:"dataFormat" yaml:"dataFormat"`
	Description *string `field:"optional" json:"description" yaml:"description"`
	Location *string `field:"optional" json:"location" yaml:"location"`
	Name *string `field:"optional" json:"name" yaml:"name"`
	Owner *string `field:"optional" json:"owner" yaml:"owner"`
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	PartitionKeys *[]Column `field:"optional" json:"partitionKeys" yaml:"partitionKeys"`
	Retention awscdk.Duration `field:"optional" json:"retention" yaml:"retention"`
	SerdeName *string `field:"optional" json:"serdeName" yaml:"serdeName"`
	SerdeParameters *map[string]*string `field:"optional" json:"serdeParameters" yaml:"serdeParameters"`
	StorageParameters *map[string]*string `field:"optional" json:"storageParameters" yaml:"storageParameters"`
	StoredAsSubDirectories *bool `field:"optional" json:"storedAsSubDirectories" yaml:"storedAsSubDirectories"`
	TableType TableType `field:"optional" json:"tableType" yaml:"tableType"`
	TargetTable Table `field:"optional" json:"targetTable" yaml:"targetTable"`
	ViewExpandedText *string `field:"optional" json:"viewExpandedText" yaml:"viewExpandedText"`
	ViewOriginalText *string `field:"optional" json:"viewOriginalText" yaml:"viewOriginalText"`
}

