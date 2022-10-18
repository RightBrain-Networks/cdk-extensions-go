package glue

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"cdk-extensions.glue.ArrayColumn",
		reflect.TypeOf((*ArrayColumn)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "comment", GoGetter: "Comment"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "typeString", GoGetter: "TypeString"},
		},
		func() interface{} {
			j := jsiiProxy_ArrayColumn{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Column)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.glue.ArrayColumnProps",
		reflect.TypeOf((*ArrayColumnProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.glue.AssetCode",
		reflect.TypeOf((*AssetCode)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_AssetCode{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Code)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"cdk-extensions.glue.BasicColumn",
		reflect.TypeOf((*BasicColumn)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "comment", GoGetter: "Comment"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "typeString", GoGetter: "TypeString"},
		},
		func() interface{} {
			j := jsiiProxy_BasicColumn{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Column)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.glue.BasicColumnProps",
		reflect.TypeOf((*BasicColumnProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.glue.ClassificationString",
		reflect.TypeOf((*ClassificationString)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_ClassificationString{}
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.glue.CloudWatchEncryption",
		reflect.TypeOf((*CloudWatchEncryption)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cdk-extensions.glue.CloudWatchEncryptionMode",
		reflect.TypeOf((*CloudWatchEncryptionMode)(nil)).Elem(),
		map[string]interface{}{
			"KMS": CloudWatchEncryptionMode_KMS,
		},
	)
	_jsii_.RegisterClass(
		"cdk-extensions.glue.Code",
		reflect.TypeOf((*Code)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			return &jsiiProxy_Code{}
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.glue.CodeConfig",
		reflect.TypeOf((*CodeConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.glue.Column",
		reflect.TypeOf((*Column)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "comment", GoGetter: "Comment"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "typeString", GoGetter: "TypeString"},
		},
		func() interface{} {
			return &jsiiProxy_Column{}
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.glue.ColumnProps",
		reflect.TypeOf((*ColumnProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cdk-extensions.glue.ConfigurationVersion",
		reflect.TypeOf((*ConfigurationVersion)(nil)).Elem(),
		map[string]interface{}{
			"V1_0": ConfigurationVersion_V1_0,
		},
	)
	_jsii_.RegisterClass(
		"cdk-extensions.glue.Connection",
		reflect.TypeOf((*Connection)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addMatchCriteria", GoMethod: "AddMatchCriteria"},
			_jsii_.MemberMethod{JsiiMethod: "addProperty", GoMethod: "AddProperty"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "connectionArn", GoGetter: "ConnectionArn"},
			_jsii_.MemberProperty{JsiiProperty: "connectionName", GoGetter: "ConnectionName"},
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberProperty{JsiiProperty: "connectionType", GoGetter: "ConnectionType"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "resource", GoGetter: "Resource"},
			_jsii_.MemberProperty{JsiiProperty: "securityGroup", GoGetter: "SecurityGroup"},
			_jsii_.MemberProperty{JsiiProperty: "securityGroups", GoGetter: "SecurityGroups"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "subnets", GoGetter: "Subnets"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "vpc", GoGetter: "Vpc"},
		},
		func() interface{} {
			j := jsiiProxy_Connection{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.Type__awsec2IConnectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.glue.ConnectionProps",
		reflect.TypeOf((*ConnectionProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cdk-extensions.glue.ConnectionType",
		reflect.TypeOf((*ConnectionType)(nil)).Elem(),
		map[string]interface{}{
			"JDBC": ConnectionType_JDBC,
			"KAFKA": ConnectionType_KAFKA,
			"MONGODB": ConnectionType_MONGODB,
			"NETWORK": ConnectionType_NETWORK,
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.glue.ContinuousLoggingProps",
		reflect.TypeOf((*ContinuousLoggingProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.glue.Crawler",
		reflect.TypeOf((*Crawler)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addClassifier", GoMethod: "AddClassifier"},
			_jsii_.MemberMethod{JsiiMethod: "addTarget", GoMethod: "AddTarget"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "configuration", GoGetter: "Configuration"},
			_jsii_.MemberProperty{JsiiProperty: "crawlerArn", GoGetter: "CrawlerArn"},
			_jsii_.MemberProperty{JsiiProperty: "crawlerName", GoGetter: "CrawlerName"},
			_jsii_.MemberProperty{JsiiProperty: "database", GoGetter: "Database"},
			_jsii_.MemberProperty{JsiiProperty: "deleteBehavior", GoGetter: "DeleteBehavior"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "recrawlBehavior", GoGetter: "RecrawlBehavior"},
			_jsii_.MemberProperty{JsiiProperty: "resource", GoGetter: "Resource"},
			_jsii_.MemberProperty{JsiiProperty: "role", GoGetter: "Role"},
			_jsii_.MemberProperty{JsiiProperty: "scheduleExpression", GoGetter: "ScheduleExpression"},
			_jsii_.MemberProperty{JsiiProperty: "securityConfiguration", GoGetter: "SecurityConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tablePrefix", GoGetter: "TablePrefix"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updateBehavior", GoGetter: "UpdateBehavior"},
		},
		func() interface{} {
			j := jsiiProxy_Crawler{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.glue.CrawlerConfiguration",
		reflect.TypeOf((*CrawlerConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.glue.CrawlerProps",
		reflect.TypeOf((*CrawlerProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.glue.CrawlerTargetCollection",
		reflect.TypeOf((*CrawlerTargetCollection)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.glue.DataFormat",
		reflect.TypeOf((*DataFormat)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "classificationString", GoGetter: "ClassificationString"},
			_jsii_.MemberProperty{JsiiProperty: "inputFormat", GoGetter: "InputFormat"},
			_jsii_.MemberProperty{JsiiProperty: "outputFormat", GoGetter: "OutputFormat"},
			_jsii_.MemberProperty{JsiiProperty: "serializationLibrary", GoGetter: "SerializationLibrary"},
		},
		func() interface{} {
			return &jsiiProxy_DataFormat{}
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.glue.DataFormatProps",
		reflect.TypeOf((*DataFormatProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.glue.Database",
		reflect.TypeOf((*Database)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "catalogArn", GoGetter: "CatalogArn"},
			_jsii_.MemberProperty{JsiiProperty: "catalogId", GoGetter: "CatalogId"},
			_jsii_.MemberProperty{JsiiProperty: "databaseArn", GoGetter: "DatabaseArn"},
			_jsii_.MemberProperty{JsiiProperty: "databaseName", GoGetter: "DatabaseName"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "locationUri", GoGetter: "LocationUri"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "resource", GoGetter: "Resource"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Database{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.glue.DatabaseProps",
		reflect.TypeOf((*DatabaseProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cdk-extensions.glue.DeleteBehavior",
		reflect.TypeOf((*DeleteBehavior)(nil)).Elem(),
		map[string]interface{}{
			"DELETE_FROM_DATABASE": DeleteBehavior_DELETE_FROM_DATABASE,
			"DEPRECATE_IN_DATABASE": DeleteBehavior_DEPRECATE_IN_DATABASE,
			"LOG": DeleteBehavior_LOG,
		},
	)
	_jsii_.RegisterClass(
		"cdk-extensions.glue.GlueVersion",
		reflect.TypeOf((*GlueVersion)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
		},
		func() interface{} {
			return &jsiiProxy_GlueVersion{}
		},
	)
	_jsii_.RegisterInterface(
		"cdk-extensions.glue.ICrawlerTarget",
		reflect.TypeOf((*ICrawlerTarget)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			return &jsiiProxy_ICrawlerTarget{}
		},
	)
	_jsii_.RegisterInterface(
		"cdk-extensions.glue.ITriggerAction",
		reflect.TypeOf((*ITriggerAction)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			return &jsiiProxy_ITriggerAction{}
		},
	)
	_jsii_.RegisterInterface(
		"cdk-extensions.glue.ITriggerPredicate",
		reflect.TypeOf((*ITriggerPredicate)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			return &jsiiProxy_ITriggerPredicate{}
		},
	)
	_jsii_.RegisterClass(
		"cdk-extensions.glue.InputFormat",
		reflect.TypeOf((*InputFormat)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "className", GoGetter: "ClassName"},
		},
		func() interface{} {
			return &jsiiProxy_InputFormat{}
		},
	)
	_jsii_.RegisterClass(
		"cdk-extensions.glue.JdbcTarget",
		reflect.TypeOf((*JdbcTarget)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addExclusion", GoMethod: "AddExclusion"},
			_jsii_.MemberMethod{JsiiMethod: "addPath", GoMethod: "AddPath"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
		},
		func() interface{} {
			j := jsiiProxy_JdbcTarget{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ICrawlerTarget)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.glue.JdbcTargetOptions",
		reflect.TypeOf((*JdbcTargetOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.glue.Job",
		reflect.TypeOf((*Job)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addArgument", GoMethod: "AddArgument"},
			_jsii_.MemberMethod{JsiiMethod: "addConnection", GoMethod: "AddConnection"},
			_jsii_.MemberProperty{JsiiProperty: "allocatedCapacity", GoGetter: "AllocatedCapacity"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberProperty{JsiiProperty: "continuousLogging", GoGetter: "ContinuousLogging"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "executable", GoGetter: "Executable"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "jobArn", GoGetter: "JobArn"},
			_jsii_.MemberProperty{JsiiProperty: "jobName", GoGetter: "JobName"},
			_jsii_.MemberProperty{JsiiProperty: "logGroup", GoGetter: "LogGroup"},
			_jsii_.MemberProperty{JsiiProperty: "maxCapacity", GoGetter: "MaxCapacity"},
			_jsii_.MemberProperty{JsiiProperty: "maxConcurrentRuns", GoGetter: "MaxConcurrentRuns"},
			_jsii_.MemberProperty{JsiiProperty: "maxRetries", GoGetter: "MaxRetries"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "notifyDelayAfter", GoGetter: "NotifyDelayAfter"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "resource", GoGetter: "Resource"},
			_jsii_.MemberProperty{JsiiProperty: "role", GoGetter: "Role"},
			_jsii_.MemberProperty{JsiiProperty: "securityConfiguration", GoGetter: "SecurityConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "timeout", GoGetter: "Timeout"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "workerCount", GoGetter: "WorkerCount"},
			_jsii_.MemberProperty{JsiiProperty: "workerType", GoGetter: "WorkerType"},
		},
		func() interface{} {
			j := jsiiProxy_Job{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.glue.JobBookmarksEncryption",
		reflect.TypeOf((*JobBookmarksEncryption)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cdk-extensions.glue.JobBookmarksEncryptionMode",
		reflect.TypeOf((*JobBookmarksEncryptionMode)(nil)).Elem(),
		map[string]interface{}{
			"CLIENT_SIDE_KMS": JobBookmarksEncryptionMode_CLIENT_SIDE_KMS,
		},
	)
	_jsii_.RegisterClass(
		"cdk-extensions.glue.JobExecutable",
		reflect.TypeOf((*JobExecutable)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			return &jsiiProxy_JobExecutable{}
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.glue.JobExecutableConfig",
		reflect.TypeOf((*JobExecutableConfig)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cdk-extensions.glue.JobLanguage",
		reflect.TypeOf((*JobLanguage)(nil)).Elem(),
		map[string]interface{}{
			"PYTHON": JobLanguage_PYTHON,
			"SCALA": JobLanguage_SCALA,
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.glue.JobProps",
		reflect.TypeOf((*JobProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.glue.JobType",
		reflect.TypeOf((*JobType)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
		},
		func() interface{} {
			return &jsiiProxy_JobType{}
		},
	)
	_jsii_.RegisterClass(
		"cdk-extensions.glue.OutputFormat",
		reflect.TypeOf((*OutputFormat)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "className", GoGetter: "ClassName"},
		},
		func() interface{} {
			return &jsiiProxy_OutputFormat{}
		},
	)
	_jsii_.RegisterEnum(
		"cdk-extensions.glue.PartitionUpdateBehavior",
		reflect.TypeOf((*PartitionUpdateBehavior)(nil)).Elem(),
		map[string]interface{}{
			"INHERIT_FROM_TABLE": PartitionUpdateBehavior_INHERIT_FROM_TABLE,
		},
	)
	_jsii_.RegisterEnum(
		"cdk-extensions.glue.PredicateOperator",
		reflect.TypeOf((*PredicateOperator)(nil)).Elem(),
		map[string]interface{}{
			"AND": PredicateOperator_AND,
			"OR": PredicateOperator_OR,
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.glue.PythonShellExecutableProps",
		reflect.TypeOf((*PythonShellExecutableProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.glue.PythonSparkJobExecutableProps",
		reflect.TypeOf((*PythonSparkJobExecutableProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cdk-extensions.glue.PythonVersion",
		reflect.TypeOf((*PythonVersion)(nil)).Elem(),
		map[string]interface{}{
			"THREE": PythonVersion_THREE,
			"TWO": PythonVersion_TWO,
		},
	)
	_jsii_.RegisterEnum(
		"cdk-extensions.glue.RecrawlBehavior",
		reflect.TypeOf((*RecrawlBehavior)(nil)).Elem(),
		map[string]interface{}{
			"EVENT_MODE": RecrawlBehavior_EVENT_MODE,
			"EVERYTHING": RecrawlBehavior_EVERYTHING,
			"NEW_FOLDERS_ONLY": RecrawlBehavior_NEW_FOLDERS_ONLY,
		},
	)
	_jsii_.RegisterClass(
		"cdk-extensions.glue.S3Code",
		reflect.TypeOf((*S3Code)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_S3Code{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Code)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.glue.S3Encryption",
		reflect.TypeOf((*S3Encryption)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cdk-extensions.glue.S3EncryptionMode",
		reflect.TypeOf((*S3EncryptionMode)(nil)).Elem(),
		map[string]interface{}{
			"KMS": S3EncryptionMode_KMS,
			"S3_MANAGED": S3EncryptionMode_S3_MANAGED,
		},
	)
	_jsii_.RegisterClass(
		"cdk-extensions.glue.S3Target",
		reflect.TypeOf((*S3Target)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addExclusion", GoMethod: "AddExclusion"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "bucket", GoGetter: "Bucket"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "exclusions", GoGetter: "Exclusions"},
			_jsii_.MemberProperty{JsiiProperty: "keyPrefix", GoGetter: "KeyPrefix"},
			_jsii_.MemberProperty{JsiiProperty: "sampleSize", GoGetter: "SampleSize"},
		},
		func() interface{} {
			j := jsiiProxy_S3Target{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ICrawlerTarget)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.glue.S3TargetOptions",
		reflect.TypeOf((*S3TargetOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.glue.ScalaJobExecutableProps",
		reflect.TypeOf((*ScalaJobExecutableProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.glue.SecurityConfiguration",
		reflect.TypeOf((*SecurityConfiguration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "cloudWatchEncryption", GoGetter: "CloudWatchEncryption"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "jobBookmarksEncryption", GoGetter: "JobBookmarksEncryption"},
			_jsii_.MemberProperty{JsiiProperty: "key", GoGetter: "Key"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "resource", GoGetter: "Resource"},
			_jsii_.MemberProperty{JsiiProperty: "s3Encryption", GoGetter: "S3Encryption"},
			_jsii_.MemberProperty{JsiiProperty: "securityConfigurationName", GoGetter: "SecurityConfigurationName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_SecurityConfiguration{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.glue.SecurityConfigurationProps",
		reflect.TypeOf((*SecurityConfigurationProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.glue.SerializationLibrary",
		reflect.TypeOf((*SerializationLibrary)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "className", GoGetter: "ClassName"},
		},
		func() interface{} {
			return &jsiiProxy_SerializationLibrary{}
		},
	)
	_jsii_.RegisterClass(
		"cdk-extensions.glue.StructColumn",
		reflect.TypeOf((*StructColumn)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addColumn", GoMethod: "AddColumn"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "comment", GoGetter: "Comment"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "typeString", GoGetter: "TypeString"},
		},
		func() interface{} {
			j := jsiiProxy_StructColumn{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Column)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.glue.StructColumnProps",
		reflect.TypeOf((*StructColumnProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.glue.Table",
		reflect.TypeOf((*Table)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addColumn", GoMethod: "AddColumn"},
			_jsii_.MemberMethod{JsiiMethod: "addParameter", GoMethod: "AddParameter"},
			_jsii_.MemberMethod{JsiiMethod: "addPartitionKey", GoMethod: "AddPartitionKey"},
			_jsii_.MemberMethod{JsiiMethod: "addSerdeParameter", GoMethod: "AddSerdeParameter"},
			_jsii_.MemberMethod{JsiiMethod: "addStorageParameter", GoMethod: "AddStorageParameter"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "compressed", GoGetter: "Compressed"},
			_jsii_.MemberProperty{JsiiProperty: "database", GoGetter: "Database"},
			_jsii_.MemberProperty{JsiiProperty: "dataFormat", GoGetter: "DataFormat"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
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
			j := jsiiProxy_Table{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			return &j
		},
	)
	_jsii_.RegisterEnum(
		"cdk-extensions.glue.TableGroupingPolicy",
		reflect.TypeOf((*TableGroupingPolicy)(nil)).Elem(),
		map[string]interface{}{
			"COMBINE_COMPATIBLE_SCHEMAS": TableGroupingPolicy_COMBINE_COMPATIBLE_SCHEMAS,
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.glue.TableProps",
		reflect.TypeOf((*TableProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cdk-extensions.glue.TableType",
		reflect.TypeOf((*TableType)(nil)).Elem(),
		map[string]interface{}{
			"EXTERNAL_TABLE": TableType_EXTERNAL_TABLE,
			"VIRTUAL_VIEW": TableType_VIRTUAL_VIEW,
		},
	)
	_jsii_.RegisterEnum(
		"cdk-extensions.glue.TableUpdateBehavior",
		reflect.TypeOf((*TableUpdateBehavior)(nil)).Elem(),
		map[string]interface{}{
			"MERGE_NEW_COLUMNS": TableUpdateBehavior_MERGE_NEW_COLUMNS,
		},
	)
	_jsii_.RegisterClass(
		"cdk-extensions.glue.Trigger",
		reflect.TypeOf((*Trigger)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addAction", GoMethod: "AddAction"},
			_jsii_.MemberMethod{JsiiMethod: "addPredicate", GoMethod: "AddPredicate"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "predicateOperator", GoGetter: "PredicateOperator"},
			_jsii_.MemberProperty{JsiiProperty: "resource", GoGetter: "Resource"},
			_jsii_.MemberProperty{JsiiProperty: "schedule", GoGetter: "Schedule"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "startOnCreation", GoGetter: "StartOnCreation"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "workflow", GoGetter: "Workflow"},
			_jsii_.MemberProperty{JsiiProperty: "workflowArn", GoGetter: "WorkflowArn"},
			_jsii_.MemberProperty{JsiiProperty: "workflowName", GoGetter: "WorkflowName"},
		},
		func() interface{} {
			j := jsiiProxy_Trigger{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.glue.TriggerOptions",
		reflect.TypeOf((*TriggerOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.glue.TriggerProps",
		reflect.TypeOf((*TriggerProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cdk-extensions.glue.TriggerType",
		reflect.TypeOf((*TriggerType)(nil)).Elem(),
		map[string]interface{}{
			"CONDITIONAL": TriggerType_CONDITIONAL,
			"EVENT": TriggerType_EVENT,
			"ON_DEMAND": TriggerType_ON_DEMAND,
			"SCHEDULED": TriggerType_SCHEDULED,
		},
	)
	_jsii_.RegisterEnum(
		"cdk-extensions.glue.UpdateBehavior",
		reflect.TypeOf((*UpdateBehavior)(nil)).Elem(),
		map[string]interface{}{
			"UPDATE_IN_DATABASE": UpdateBehavior_UPDATE_IN_DATABASE,
			"LOG": UpdateBehavior_LOG,
		},
	)
	_jsii_.RegisterClass(
		"cdk-extensions.glue.WorkerType",
		reflect.TypeOf((*WorkerType)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
		},
		func() interface{} {
			return &jsiiProxy_WorkerType{}
		},
	)
	_jsii_.RegisterClass(
		"cdk-extensions.glue.Workflow",
		reflect.TypeOf((*Workflow)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addTrigger", GoMethod: "AddTrigger"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "resource", GoGetter: "Resource"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "workflowArn", GoGetter: "WorkflowArn"},
			_jsii_.MemberProperty{JsiiProperty: "workflowName", GoGetter: "WorkflowName"},
		},
		func() interface{} {
			j := jsiiProxy_Workflow{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.glue.WorkflowProps",
		reflect.TypeOf((*WorkflowProps)(nil)).Elem(),
	)
}
