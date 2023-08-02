package glue

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Configuration for Table.
type TableProps struct {
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
	// Database object to add Table to.
	Database Database `field:"required" json:"database" yaml:"database"`
	// A list of the Columns in the table.
	// See: [AWS::Glue::Table StorageDescriptor](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-storagedescriptor.html#cfn-glue-table-storagedescriptor-columns)
	//
	Columns *[]Column `field:"optional" json:"columns" yaml:"columns"`
	// True if the data in the table is compressed, or False if not.
	// See: [AWS::Glue::Table StorageDescriptor](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-storagedescriptor.html#cfn-glue-table-storagedescriptor-compressed)
	//
	Compressed *bool `field:"optional" json:"compressed" yaml:"compressed"`
	// DataFormat object indicating the expected input/output format.
	DataFormat DataFormat `field:"optional" json:"dataFormat" yaml:"dataFormat"`
	// A description for the Table.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The physical location of the table.
	//
	// By default, this takes the form of the warehouse location, followed by the database location in the warehouse, followed by the table name.
	// See: [AWS::Glue::Table StorageDescriptor](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-storagedescriptor.html#cfn-glue-table-storagedescriptor-location)
	//
	Location *string `field:"optional" json:"location" yaml:"location"`
	// A name for the Table.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The table owner.
	//
	// Included for Apache Hive compatibility. Not used in the normal course of AWS Glue operations.
	// See: [AWS::Glue::Table TableInput](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-tableinput.html#cfn-glue-table-tableinput-owner)
	//
	Owner *string `field:"optional" json:"owner" yaml:"owner"`
	// These key-value pairs define properties associated with the table.
	// See: [AWS::Glue::Table TableInput](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-tableinput.html#cfn-glue-table-tableinput-parameters)
	//
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// A list of columns by which the table is partitioned.
	//
	// Only primitive types are supported as partition keys.
	// See: [AWS::Glue::Table TableInput](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-tableinput.html#cfn-glue-table-tableinput-parameterskeys)
	//
	PartitionKeys *[]Column `field:"optional" json:"partitionKeys" yaml:"partitionKeys"`
	// The retention time for this table.
	// See: [AWS::Glue::Table TableInput](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-tableinput.html#cfn-glue-table-tableinput-retention)
	//
	Retention awscdk.Duration `field:"optional" json:"retention" yaml:"retention"`
	// Name of the SerDe.
	// See: [AWS::Glue::Table SerdeInfo](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-serdeinfo.html#cfn-glue-table-serdeinfo-name)
	//
	SerdeName *string `field:"optional" json:"serdeName" yaml:"serdeName"`
	// These key-value pairs define initialization parameters for the SerDe.
	// See: [AWS::Glue::Table SerdeInfo](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-serdeinfo.html#cfn-glue-table-serdeinfo-parameters)
	//
	SerdeParameters *map[string]*string `field:"optional" json:"serdeParameters" yaml:"serdeParameters"`
	// The user-supplied properties in key-value form.
	// See: [AWS::Glue::Table StorageDescriptor](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-storagedescriptor.html#cfn-glue-table-storagedescriptor-parameters)
	//
	StorageParameters *map[string]*string `field:"optional" json:"storageParameters" yaml:"storageParameters"`
	// True if the table data is stored in subdirectories, or False if not.
	// See: [AWS::Glue::Table StorageDescriptor](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-storagedescriptor.html#cfn-glue-table-storagedescriptor-storedassubdirectories)
	//
	StoredAsSubDirectories *bool `field:"optional" json:"storedAsSubDirectories" yaml:"storedAsSubDirectories"`
	// The type of this table.
	//
	// AWS Glue will create tables with the EXTERNAL_TABLE type. Other services, such as Athena, may create tables with additional table types.
	// See: [AWS::Glue::Table TableInput](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-tableinput.html#cfn-glue-table-tableinput-tabletype)
	//
	TableType TableType `field:"optional" json:"tableType" yaml:"tableType"`
	// A TableIdentifier structure that describes a target table for resource linking.
	// See: [AWS::Glue::Table TableInput](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-tableinput.html#cfn-glue-table-tableinput-targettable)
	//
	TargetTable Table `field:"optional" json:"targetTable" yaml:"targetTable"`
	// Included for Apache Hive compatibility.
	//
	// Not used in the normal course of AWS Glue operations.
	// See: [AWS::Glue::Table TableInput](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-tableinput.html#cfn-glue-table-tableinput-viewexpandedtext)
	//
	ViewExpandedText *string `field:"optional" json:"viewExpandedText" yaml:"viewExpandedText"`
	// Included for Apache Hive compatibility.
	//
	// Not used in the normal course of AWS Glue operations. If the table is a VIRTUAL_VIEW, certain Athena configuration encoded in base64.
	// See: [AWS::Glue::Table TableInput](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-tableinput.html#cfn-glue-table-tableinput-vieworiginaltext)
	//
	ViewOriginalText *string `field:"optional" json:"viewOriginalText" yaml:"viewOriginalText"`
}

