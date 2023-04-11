package gluetables

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"cdk-extensions.glue_tables.AlbLogsTable",
		reflect.TypeOf((*AlbLogsTable)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addColumn", GoMethod: "AddColumn"},
			_jsii_.MemberMethod{JsiiMethod: "addParameter", GoMethod: "AddParameter"},
			_jsii_.MemberMethod{JsiiMethod: "addPartitionKey", GoMethod: "AddPartitionKey"},
			_jsii_.MemberMethod{JsiiMethod: "addSerdeParameter", GoMethod: "AddSerdeParameter"},
			_jsii_.MemberMethod{JsiiMethod: "addStorageParameter", GoMethod: "AddStorageParameter"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "compressed", GoGetter: "Compressed"},
			_jsii_.MemberProperty{JsiiProperty: "createQueries", GoGetter: "CreateQueries"},
			_jsii_.MemberProperty{JsiiProperty: "database", GoGetter: "Database"},
			_jsii_.MemberProperty{JsiiProperty: "dataFormat", GoGetter: "DataFormat"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyQueryNames", GoGetter: "FriendlyQueryNames"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "location", GoGetter: "Location"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "owner", GoGetter: "Owner"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "renderStorageDescriptor", GoMethod: "RenderStorageDescriptor"},
			_jsii_.MemberProperty{JsiiProperty: "resource", GoGetter: "Resource"},
			_jsii_.MemberProperty{JsiiProperty: "retention", GoGetter: "Retention"},
			_jsii_.MemberProperty{JsiiProperty: "serdeName", GoGetter: "SerdeName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "status5xxNamedQuery", GoGetter: "Status5xxNamedQuery"},
			_jsii_.MemberProperty{JsiiProperty: "storedAsSubDirectories", GoGetter: "StoredAsSubDirectories"},
			_jsii_.MemberProperty{JsiiProperty: "tableArn", GoGetter: "TableArn"},
			_jsii_.MemberProperty{JsiiProperty: "tableName", GoGetter: "TableName"},
			_jsii_.MemberProperty{JsiiProperty: "tableType", GoGetter: "TableType"},
			_jsii_.MemberProperty{JsiiProperty: "targetTable", GoGetter: "TargetTable"},
			_jsii_.MemberProperty{JsiiProperty: "topIpsNamedQuery", GoGetter: "TopIpsNamedQuery"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "viewExpandedText", GoGetter: "ViewExpandedText"},
			_jsii_.MemberProperty{JsiiProperty: "viewOriginalText", GoGetter: "ViewOriginalText"},
		},
		func() interface{} {
			j := jsiiProxy_AlbLogsTable{}
			_jsii_.InitJsiiProxy(&j.Type__glueTable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.glue_tables.AlbLogsTableProps",
		reflect.TypeOf((*AlbLogsTableProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.glue_tables.CloudfrontLogsTable",
		reflect.TypeOf((*CloudfrontLogsTable)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addColumn", GoMethod: "AddColumn"},
			_jsii_.MemberMethod{JsiiMethod: "addParameter", GoMethod: "AddParameter"},
			_jsii_.MemberMethod{JsiiMethod: "addPartitionKey", GoMethod: "AddPartitionKey"},
			_jsii_.MemberMethod{JsiiMethod: "addSerdeParameter", GoMethod: "AddSerdeParameter"},
			_jsii_.MemberMethod{JsiiMethod: "addStorageParameter", GoMethod: "AddStorageParameter"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "compressed", GoGetter: "Compressed"},
			_jsii_.MemberProperty{JsiiProperty: "createQueries", GoGetter: "CreateQueries"},
			_jsii_.MemberProperty{JsiiProperty: "database", GoGetter: "Database"},
			_jsii_.MemberProperty{JsiiProperty: "dataFormat", GoGetter: "DataFormat"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "distributionStatisticsNamedQuery", GoGetter: "DistributionStatisticsNamedQuery"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyQueryNames", GoGetter: "FriendlyQueryNames"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "location", GoGetter: "Location"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "owner", GoGetter: "Owner"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "renderStorageDescriptor", GoMethod: "RenderStorageDescriptor"},
			_jsii_.MemberProperty{JsiiProperty: "requestErrorsNamedQuery", GoGetter: "RequestErrorsNamedQuery"},
			_jsii_.MemberProperty{JsiiProperty: "resource", GoGetter: "Resource"},
			_jsii_.MemberProperty{JsiiProperty: "retention", GoGetter: "Retention"},
			_jsii_.MemberProperty{JsiiProperty: "serdeName", GoGetter: "SerdeName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "storedAsSubDirectories", GoGetter: "StoredAsSubDirectories"},
			_jsii_.MemberProperty{JsiiProperty: "tableArn", GoGetter: "TableArn"},
			_jsii_.MemberProperty{JsiiProperty: "tableName", GoGetter: "TableName"},
			_jsii_.MemberProperty{JsiiProperty: "tableType", GoGetter: "TableType"},
			_jsii_.MemberProperty{JsiiProperty: "targetTable", GoGetter: "TargetTable"},
			_jsii_.MemberProperty{JsiiProperty: "topIpsNamedQuery", GoGetter: "TopIpsNamedQuery"},
			_jsii_.MemberProperty{JsiiProperty: "topObjectsNamedQuery", GoGetter: "TopObjectsNamedQuery"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "viewExpandedText", GoGetter: "ViewExpandedText"},
			_jsii_.MemberProperty{JsiiProperty: "viewOriginalText", GoGetter: "ViewOriginalText"},
		},
		func() interface{} {
			j := jsiiProxy_CloudfrontLogsTable{}
			_jsii_.InitJsiiProxy(&j.Type__glueTable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.glue_tables.CloudfrontLogsTableProps",
		reflect.TypeOf((*CloudfrontLogsTableProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.glue_tables.CloudtrailTable",
		reflect.TypeOf((*CloudtrailTable)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addColumn", GoMethod: "AddColumn"},
			_jsii_.MemberMethod{JsiiMethod: "addParameter", GoMethod: "AddParameter"},
			_jsii_.MemberMethod{JsiiMethod: "addPartitionKey", GoMethod: "AddPartitionKey"},
			_jsii_.MemberMethod{JsiiMethod: "addSerdeParameter", GoMethod: "AddSerdeParameter"},
			_jsii_.MemberMethod{JsiiMethod: "addStorageParameter", GoMethod: "AddStorageParameter"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "compressed", GoGetter: "Compressed"},
			_jsii_.MemberProperty{JsiiProperty: "createQueries", GoGetter: "CreateQueries"},
			_jsii_.MemberProperty{JsiiProperty: "database", GoGetter: "Database"},
			_jsii_.MemberProperty{JsiiProperty: "dataFormat", GoGetter: "DataFormat"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyQueryNames", GoGetter: "FriendlyQueryNames"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "location", GoGetter: "Location"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "owner", GoGetter: "Owner"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "renderStorageDescriptor", GoMethod: "RenderStorageDescriptor"},
			_jsii_.MemberProperty{JsiiProperty: "resource", GoGetter: "Resource"},
			_jsii_.MemberProperty{JsiiProperty: "retention", GoGetter: "Retention"},
			_jsii_.MemberProperty{JsiiProperty: "serdeName", GoGetter: "SerdeName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "storedAsSubDirectories", GoGetter: "StoredAsSubDirectories"},
			_jsii_.MemberProperty{JsiiProperty: "tableArn", GoGetter: "TableArn"},
			_jsii_.MemberProperty{JsiiProperty: "tableName", GoGetter: "TableName"},
			_jsii_.MemberProperty{JsiiProperty: "tableType", GoGetter: "TableType"},
			_jsii_.MemberProperty{JsiiProperty: "targetTable", GoGetter: "TargetTable"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "unauthorizedNamedQuery", GoGetter: "UnauthorizedNamedQuery"},
			_jsii_.MemberProperty{JsiiProperty: "userLoginsNamedQuery", GoGetter: "UserLoginsNamedQuery"},
			_jsii_.MemberProperty{JsiiProperty: "viewExpandedText", GoGetter: "ViewExpandedText"},
			_jsii_.MemberProperty{JsiiProperty: "viewOriginalText", GoGetter: "ViewOriginalText"},
		},
		func() interface{} {
			j := jsiiProxy_CloudtrailTable{}
			_jsii_.InitJsiiProxy(&j.Type__glueTable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.glue_tables.CloudtrailTableProps",
		reflect.TypeOf((*CloudtrailTableProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.glue_tables.FlowLogsTable",
		reflect.TypeOf((*FlowLogsTable)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addColumn", GoMethod: "AddColumn"},
			_jsii_.MemberMethod{JsiiMethod: "addParameter", GoMethod: "AddParameter"},
			_jsii_.MemberMethod{JsiiMethod: "addPartitionKey", GoMethod: "AddPartitionKey"},
			_jsii_.MemberMethod{JsiiMethod: "addSerdeParameter", GoMethod: "AddSerdeParameter"},
			_jsii_.MemberMethod{JsiiMethod: "addStorageParameter", GoMethod: "AddStorageParameter"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "compressed", GoGetter: "Compressed"},
			_jsii_.MemberProperty{JsiiProperty: "createQueries", GoGetter: "CreateQueries"},
			_jsii_.MemberProperty{JsiiProperty: "database", GoGetter: "Database"},
			_jsii_.MemberProperty{JsiiProperty: "dataFormat", GoGetter: "DataFormat"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "format", GoGetter: "Format"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyQueryNames", GoGetter: "FriendlyQueryNames"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalRejectedNamedQuery", GoGetter: "InternalRejectedNamedQuery"},
			_jsii_.MemberProperty{JsiiProperty: "location", GoGetter: "Location"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "owner", GoGetter: "Owner"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "renderStorageDescriptor", GoMethod: "RenderStorageDescriptor"},
			_jsii_.MemberProperty{JsiiProperty: "resource", GoGetter: "Resource"},
			_jsii_.MemberProperty{JsiiProperty: "retention", GoGetter: "Retention"},
			_jsii_.MemberProperty{JsiiProperty: "serdeName", GoGetter: "SerdeName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "storedAsSubDirectories", GoGetter: "StoredAsSubDirectories"},
			_jsii_.MemberProperty{JsiiProperty: "tableArn", GoGetter: "TableArn"},
			_jsii_.MemberProperty{JsiiProperty: "tableName", GoGetter: "TableName"},
			_jsii_.MemberProperty{JsiiProperty: "tableType", GoGetter: "TableType"},
			_jsii_.MemberProperty{JsiiProperty: "targetTable", GoGetter: "TargetTable"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "viewExpandedText", GoGetter: "ViewExpandedText"},
			_jsii_.MemberProperty{JsiiProperty: "viewOriginalText", GoGetter: "ViewOriginalText"},
		},
		func() interface{} {
			j := jsiiProxy_FlowLogsTable{}
			_jsii_.InitJsiiProxy(&j.Type__glueTable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.glue_tables.FlowLogsTableProps",
		reflect.TypeOf((*FlowLogsTableProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.glue_tables.S3AccessLogsTable",
		reflect.TypeOf((*S3AccessLogsTable)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addColumn", GoMethod: "AddColumn"},
			_jsii_.MemberMethod{JsiiMethod: "addParameter", GoMethod: "AddParameter"},
			_jsii_.MemberMethod{JsiiMethod: "addPartitionKey", GoMethod: "AddPartitionKey"},
			_jsii_.MemberMethod{JsiiMethod: "addSerdeParameter", GoMethod: "AddSerdeParameter"},
			_jsii_.MemberMethod{JsiiMethod: "addStorageParameter", GoMethod: "AddStorageParameter"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "compressed", GoGetter: "Compressed"},
			_jsii_.MemberProperty{JsiiProperty: "createQueries", GoGetter: "CreateQueries"},
			_jsii_.MemberProperty{JsiiProperty: "database", GoGetter: "Database"},
			_jsii_.MemberProperty{JsiiProperty: "dataFormat", GoGetter: "DataFormat"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyQueryNames", GoGetter: "FriendlyQueryNames"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "location", GoGetter: "Location"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "owner", GoGetter: "Owner"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "renderStorageDescriptor", GoMethod: "RenderStorageDescriptor"},
			_jsii_.MemberProperty{JsiiProperty: "requestErrorsNamedQuery", GoGetter: "RequestErrorsNamedQuery"},
			_jsii_.MemberProperty{JsiiProperty: "resource", GoGetter: "Resource"},
			_jsii_.MemberProperty{JsiiProperty: "retention", GoGetter: "Retention"},
			_jsii_.MemberProperty{JsiiProperty: "serdeName", GoGetter: "SerdeName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "storedAsSubDirectories", GoGetter: "StoredAsSubDirectories"},
			_jsii_.MemberProperty{JsiiProperty: "tableArn", GoGetter: "TableArn"},
			_jsii_.MemberProperty{JsiiProperty: "tableName", GoGetter: "TableName"},
			_jsii_.MemberProperty{JsiiProperty: "tableType", GoGetter: "TableType"},
			_jsii_.MemberProperty{JsiiProperty: "targetTable", GoGetter: "TargetTable"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "viewExpandedText", GoGetter: "ViewExpandedText"},
			_jsii_.MemberProperty{JsiiProperty: "viewOriginalText", GoGetter: "ViewOriginalText"},
		},
		func() interface{} {
			j := jsiiProxy_S3AccessLogsTable{}
			_jsii_.InitJsiiProxy(&j.Type__glueTable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.glue_tables.S3AccessLogsTableProps",
		reflect.TypeOf((*S3AccessLogsTableProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.glue_tables.SesLogsTable",
		reflect.TypeOf((*SesLogsTable)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addColumn", GoMethod: "AddColumn"},
			_jsii_.MemberMethod{JsiiMethod: "addParameter", GoMethod: "AddParameter"},
			_jsii_.MemberMethod{JsiiMethod: "addPartitionKey", GoMethod: "AddPartitionKey"},
			_jsii_.MemberMethod{JsiiMethod: "addSerdeParameter", GoMethod: "AddSerdeParameter"},
			_jsii_.MemberMethod{JsiiMethod: "addStorageParameter", GoMethod: "AddStorageParameter"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "bouncesQuery", GoGetter: "BouncesQuery"},
			_jsii_.MemberProperty{JsiiProperty: "complaintsQuery", GoGetter: "ComplaintsQuery"},
			_jsii_.MemberProperty{JsiiProperty: "compressed", GoGetter: "Compressed"},
			_jsii_.MemberProperty{JsiiProperty: "createQueries", GoGetter: "CreateQueries"},
			_jsii_.MemberProperty{JsiiProperty: "database", GoGetter: "Database"},
			_jsii_.MemberProperty{JsiiProperty: "dataFormat", GoGetter: "DataFormat"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyQueryNames", GoGetter: "FriendlyQueryNames"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "location", GoGetter: "Location"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "owner", GoGetter: "Owner"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "renderStorageDescriptor", GoMethod: "RenderStorageDescriptor"},
			_jsii_.MemberProperty{JsiiProperty: "resource", GoGetter: "Resource"},
			_jsii_.MemberProperty{JsiiProperty: "retention", GoGetter: "Retention"},
			_jsii_.MemberProperty{JsiiProperty: "serdeName", GoGetter: "SerdeName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "storedAsSubDirectories", GoGetter: "StoredAsSubDirectories"},
			_jsii_.MemberProperty{JsiiProperty: "tableArn", GoGetter: "TableArn"},
			_jsii_.MemberProperty{JsiiProperty: "tableName", GoGetter: "TableName"},
			_jsii_.MemberProperty{JsiiProperty: "tableType", GoGetter: "TableType"},
			_jsii_.MemberProperty{JsiiProperty: "targetTable", GoGetter: "TargetTable"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "viewExpandedText", GoGetter: "ViewExpandedText"},
			_jsii_.MemberProperty{JsiiProperty: "viewOriginalText", GoGetter: "ViewOriginalText"},
		},
		func() interface{} {
			j := jsiiProxy_SesLogsTable{}
			_jsii_.InitJsiiProxy(&j.Type__glueTable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.glue_tables.SesLogsTableProps",
		reflect.TypeOf((*SesLogsTableProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.glue_tables.WafLogsTable",
		reflect.TypeOf((*WafLogsTable)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addColumn", GoMethod: "AddColumn"},
			_jsii_.MemberMethod{JsiiMethod: "addParameter", GoMethod: "AddParameter"},
			_jsii_.MemberMethod{JsiiMethod: "addPartitionKey", GoMethod: "AddPartitionKey"},
			_jsii_.MemberMethod{JsiiMethod: "addSerdeParameter", GoMethod: "AddSerdeParameter"},
			_jsii_.MemberMethod{JsiiMethod: "addStorageParameter", GoMethod: "AddStorageParameter"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "compressed", GoGetter: "Compressed"},
			_jsii_.MemberProperty{JsiiProperty: "createQueries", GoGetter: "CreateQueries"},
			_jsii_.MemberProperty{JsiiProperty: "database", GoGetter: "Database"},
			_jsii_.MemberProperty{JsiiProperty: "dataFormat", GoGetter: "DataFormat"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyQueryNames", GoGetter: "FriendlyQueryNames"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "location", GoGetter: "Location"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "owner", GoGetter: "Owner"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "renderStorageDescriptor", GoMethod: "RenderStorageDescriptor"},
			_jsii_.MemberProperty{JsiiProperty: "resource", GoGetter: "Resource"},
			_jsii_.MemberProperty{JsiiProperty: "retention", GoGetter: "Retention"},
			_jsii_.MemberProperty{JsiiProperty: "serdeName", GoGetter: "SerdeName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "status5xxNamedQuery", GoGetter: "Status5xxNamedQuery"},
			_jsii_.MemberProperty{JsiiProperty: "storedAsSubDirectories", GoGetter: "StoredAsSubDirectories"},
			_jsii_.MemberProperty{JsiiProperty: "tableArn", GoGetter: "TableArn"},
			_jsii_.MemberProperty{JsiiProperty: "tableName", GoGetter: "TableName"},
			_jsii_.MemberProperty{JsiiProperty: "tableType", GoGetter: "TableType"},
			_jsii_.MemberProperty{JsiiProperty: "targetTable", GoGetter: "TargetTable"},
			_jsii_.MemberProperty{JsiiProperty: "topIpsNamedQuery", GoGetter: "TopIpsNamedQuery"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "viewExpandedText", GoGetter: "ViewExpandedText"},
			_jsii_.MemberProperty{JsiiProperty: "viewOriginalText", GoGetter: "ViewOriginalText"},
		},
		func() interface{} {
			j := jsiiProxy_WafLogsTable{}
			_jsii_.InitJsiiProxy(&j.Type__glueTable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.glue_tables.WafLogsTableProps",
		reflect.TypeOf((*WafLogsTableProps)(nil)).Elem(),
	)
}
