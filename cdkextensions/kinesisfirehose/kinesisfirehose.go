package kinesisfirehose

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"cdk-extensions.kinesis_firehose.AppendDelimiterProcessor",
		reflect.TypeOf((*AppendDelimiterProcessor)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addProcessorParameter", GoMethod: "AddProcessorParameter"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "delimiter", GoGetter: "Delimiter"},
			_jsii_.MemberProperty{JsiiProperty: "processorType", GoGetter: "ProcessorType"},
		},
		func() interface{} {
			j := jsiiProxy_AppendDelimiterProcessor{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_DeliveryStreamProcessor)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.kinesis_firehose.AppendDelimiterProcessorOptions",
		reflect.TypeOf((*AppendDelimiterProcessorOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.kinesis_firehose.BackupConfiguration",
		reflect.TypeOf((*BackupConfiguration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "destination", GoGetter: "Destination"},
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
		},
		func() interface{} {
			return &jsiiProxy_BackupConfiguration{}
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.kinesis_firehose.BackupConfigurationOptions",
		reflect.TypeOf((*BackupConfigurationOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.kinesis_firehose.BackupConfigurationResult",
		reflect.TypeOf((*BackupConfigurationResult)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.kinesis_firehose.BufferingConfiguration",
		reflect.TypeOf((*BufferingConfiguration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "interval", GoGetter: "Interval"},
			_jsii_.MemberProperty{JsiiProperty: "sizeInMb", GoGetter: "SizeInMb"},
		},
		func() interface{} {
			return &jsiiProxy_BufferingConfiguration{}
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.kinesis_firehose.BufferingConfigurationOptions",
		reflect.TypeOf((*BufferingConfigurationOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.kinesis_firehose.CloudWatchLoggingConfiguration",
		reflect.TypeOf((*CloudWatchLoggingConfiguration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
			_jsii_.MemberProperty{JsiiProperty: "logGroup", GoGetter: "LogGroup"},
			_jsii_.MemberProperty{JsiiProperty: "logStream", GoGetter: "LogStream"},
		},
		func() interface{} {
			return &jsiiProxy_CloudWatchLoggingConfiguration{}
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.kinesis_firehose.CloudWatchLoggingConfigurationOptions",
		reflect.TypeOf((*CloudWatchLoggingConfigurationOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.kinesis_firehose.CommonPartitioningOptions",
		reflect.TypeOf((*CommonPartitioningOptions)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cdk-extensions.kinesis_firehose.ContentEncoding",
		reflect.TypeOf((*ContentEncoding)(nil)).Elem(),
		map[string]interface{}{
			"GZIP": ContentEncoding_GZIP,
			"NONE": ContentEncoding_NONE,
		},
	)
	_jsii_.RegisterClass(
		"cdk-extensions.kinesis_firehose.CustomProcessor",
		reflect.TypeOf((*CustomProcessor)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addParameter", GoMethod: "AddParameter"},
			_jsii_.MemberMethod{JsiiMethod: "addProcessorParameter", GoMethod: "AddProcessorParameter"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "processorType", GoGetter: "ProcessorType"},
		},
		func() interface{} {
			j := jsiiProxy_CustomProcessor{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_DeliveryStreamProcessor)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.kinesis_firehose.CustomProcessorOptions",
		reflect.TypeOf((*CustomProcessorOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.kinesis_firehose.DataFormatConversion",
		reflect.TypeOf((*DataFormatConversion)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "catalogId", GoGetter: "CatalogId"},
			_jsii_.MemberProperty{JsiiProperty: "database", GoGetter: "Database"},
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
			_jsii_.MemberProperty{JsiiProperty: "inputFormat", GoGetter: "InputFormat"},
			_jsii_.MemberProperty{JsiiProperty: "outputFormat", GoGetter: "OutputFormat"},
			_jsii_.MemberProperty{JsiiProperty: "region", GoGetter: "Region"},
			_jsii_.MemberProperty{JsiiProperty: "role", GoGetter: "Role"},
			_jsii_.MemberProperty{JsiiProperty: "table", GoGetter: "Table"},
			_jsii_.MemberProperty{JsiiProperty: "version", GoGetter: "Version"},
		},
		func() interface{} {
			return &jsiiProxy_DataFormatConversion{}
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.kinesis_firehose.DataFormatConversionOptions",
		reflect.TypeOf((*DataFormatConversionOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.kinesis_firehose.DelimitedDeaggregationOptions",
		reflect.TypeOf((*DelimitedDeaggregationOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.kinesis_firehose.DeliveryStream",
		reflect.TypeOf((*DeliveryStream)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberProperty{JsiiProperty: "deliveryStreamArn", GoGetter: "DeliveryStreamArn"},
			_jsii_.MemberProperty{JsiiProperty: "deliveryStreamName", GoGetter: "DeliveryStreamName"},
			_jsii_.MemberProperty{JsiiProperty: "destination", GoGetter: "Destination"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberProperty{JsiiProperty: "grantPrincipal", GoGetter: "GrantPrincipal"},
			_jsii_.MemberMethod{JsiiMethod: "grantPutRecords", GoMethod: "GrantPutRecords"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricBackupToS3Bytes", GoMethod: "MetricBackupToS3Bytes"},
			_jsii_.MemberMethod{JsiiMethod: "metricBackupToS3DataFreshness", GoMethod: "MetricBackupToS3DataFreshness"},
			_jsii_.MemberMethod{JsiiMethod: "metricBackupToS3Records", GoMethod: "MetricBackupToS3Records"},
			_jsii_.MemberMethod{JsiiMethod: "metricIncomingBytes", GoMethod: "MetricIncomingBytes"},
			_jsii_.MemberMethod{JsiiMethod: "metricIncomingRecords", GoMethod: "MetricIncomingRecords"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "resource", GoGetter: "Resource"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "streamType", GoGetter: "StreamType"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DeliveryStream{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IDeliveryStream)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.kinesis_firehose.DeliveryStreamAttributes",
		reflect.TypeOf((*DeliveryStreamAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.kinesis_firehose.DeliveryStreamDestination",
		reflect.TypeOf((*DeliveryStreamDestination)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "role", GoGetter: "Role"},
		},
		func() interface{} {
			return &jsiiProxy_DeliveryStreamDestination{}
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.kinesis_firehose.DeliveryStreamDestinationConfiguration",
		reflect.TypeOf((*DeliveryStreamDestinationConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.kinesis_firehose.DeliveryStreamProcessor",
		reflect.TypeOf((*DeliveryStreamProcessor)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addProcessorParameter", GoMethod: "AddProcessorParameter"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "processorType", GoGetter: "ProcessorType"},
		},
		func() interface{} {
			return &jsiiProxy_DeliveryStreamProcessor{}
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.kinesis_firehose.DeliveryStreamProcessorOptions",
		reflect.TypeOf((*DeliveryStreamProcessorOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.kinesis_firehose.DeliveryStreamProps",
		reflect.TypeOf((*DeliveryStreamProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cdk-extensions.kinesis_firehose.DeliveryStreamType",
		reflect.TypeOf((*DeliveryStreamType)(nil)).Elem(),
		map[string]interface{}{
			"DIRECT_PUT": DeliveryStreamType_DIRECT_PUT,
			"KINESIS_STREAM_AS_SOURCE": DeliveryStreamType_KINESIS_STREAM_AS_SOURCE,
		},
	)
	_jsii_.RegisterClass(
		"cdk-extensions.kinesis_firehose.DynamicPartitioning",
		reflect.TypeOf((*DynamicPartitioning)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
			_jsii_.MemberProperty{JsiiProperty: "retryInterval", GoGetter: "RetryInterval"},
		},
		func() interface{} {
			return &jsiiProxy_DynamicPartitioning{}
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.kinesis_firehose.DynamicPartitioningConfiguration",
		reflect.TypeOf((*DynamicPartitioningConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.kinesis_firehose.ExtendedS3Destination",
		reflect.TypeOf((*ExtendedS3Destination)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addProcessor", GoMethod: "AddProcessor"},
			_jsii_.MemberProperty{JsiiProperty: "backupConfiguration", GoGetter: "BackupConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "bucket", GoGetter: "Bucket"},
			_jsii_.MemberProperty{JsiiProperty: "buffering", GoGetter: "Buffering"},
			_jsii_.MemberMethod{JsiiMethod: "buildConfiguration", GoMethod: "BuildConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "cloudwatchLoggingConfiguration", GoGetter: "CloudwatchLoggingConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "compressionFormat", GoGetter: "CompressionFormat"},
			_jsii_.MemberProperty{JsiiProperty: "dataFormatConversion", GoGetter: "DataFormatConversion"},
			_jsii_.MemberProperty{JsiiProperty: "dynamicPartitioning", GoGetter: "DynamicPartitioning"},
			_jsii_.MemberProperty{JsiiProperty: "encryptionEnabled", GoGetter: "EncryptionEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "encryptionKey", GoGetter: "EncryptionKey"},
			_jsii_.MemberProperty{JsiiProperty: "errorOutputPrefix", GoGetter: "ErrorOutputPrefix"},
			_jsii_.MemberProperty{JsiiProperty: "keyPrefix", GoGetter: "KeyPrefix"},
			_jsii_.MemberProperty{JsiiProperty: "processingEnabled", GoGetter: "ProcessingEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "processorConfiguration", GoGetter: "ProcessorConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "processors", GoGetter: "Processors"},
			_jsii_.MemberMethod{JsiiMethod: "renderBackupConfiguration", GoMethod: "RenderBackupConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "renderProcessorConfiguration", GoMethod: "RenderProcessorConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "role", GoGetter: "Role"},
		},
		func() interface{} {
			j := jsiiProxy_ExtendedS3Destination{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_S3Destination)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.kinesis_firehose.ExtendedS3DestinationOptions",
		reflect.TypeOf((*ExtendedS3DestinationOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.kinesis_firehose.HiveJsonInputSerDe",
		reflect.TypeOf((*HiveJsonInputSerDe)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addTimestampFormat", GoMethod: "AddTimestampFormat"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_HiveJsonInputSerDe{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_InputFormat)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.kinesis_firehose.HiveJsonInputSerDeOptions",
		reflect.TypeOf((*HiveJsonInputSerDeOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.kinesis_firehose.HttpEndpointDestination",
		reflect.TypeOf((*HttpEndpointDestination)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "accessKey", GoGetter: "AccessKey"},
			_jsii_.MemberMethod{JsiiMethod: "addCommonAttribute", GoMethod: "AddCommonAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "addProcessor", GoMethod: "AddProcessor"},
			_jsii_.MemberProperty{JsiiProperty: "backupConfiguration", GoGetter: "BackupConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "buffering", GoGetter: "Buffering"},
			_jsii_.MemberMethod{JsiiMethod: "buildBackupConfiguration", GoMethod: "BuildBackupConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "cloudwatchLoggingConfiguration", GoGetter: "CloudwatchLoggingConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "commonAttributes", GoGetter: "CommonAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "contentEncoding", GoGetter: "ContentEncoding"},
			_jsii_.MemberProperty{JsiiProperty: "endpointName", GoGetter: "EndpointName"},
			_jsii_.MemberProperty{JsiiProperty: "endpointUrl", GoGetter: "EndpointUrl"},
			_jsii_.MemberMethod{JsiiMethod: "getOrCreateRole", GoMethod: "GetOrCreateRole"},
			_jsii_.MemberProperty{JsiiProperty: "processingEnabled", GoGetter: "ProcessingEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "processorConfiguration", GoGetter: "ProcessorConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "renderProcessorConfiguration", GoMethod: "RenderProcessorConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "retryDuration", GoGetter: "RetryDuration"},
			_jsii_.MemberProperty{JsiiProperty: "role", GoGetter: "Role"},
		},
		func() interface{} {
			j := jsiiProxy_HttpEndpointDestination{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_DeliveryStreamDestination)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.kinesis_firehose.HttpEndpointDestinationOptions",
		reflect.TypeOf((*HttpEndpointDestinationOptions)(nil)).Elem(),
	)
	_jsii_.RegisterInterface(
		"cdk-extensions.kinesis_firehose.IDeliveryStream",
		reflect.TypeOf((*IDeliveryStream)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberProperty{JsiiProperty: "deliveryStreamArn", GoGetter: "DeliveryStreamArn"},
			_jsii_.MemberProperty{JsiiProperty: "deliveryStreamName", GoGetter: "DeliveryStreamName"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberProperty{JsiiProperty: "grantPrincipal", GoGetter: "GrantPrincipal"},
			_jsii_.MemberMethod{JsiiMethod: "grantPutRecords", GoMethod: "GrantPutRecords"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricBackupToS3Bytes", GoMethod: "MetricBackupToS3Bytes"},
			_jsii_.MemberMethod{JsiiMethod: "metricBackupToS3DataFreshness", GoMethod: "MetricBackupToS3DataFreshness"},
			_jsii_.MemberMethod{JsiiMethod: "metricBackupToS3Records", GoMethod: "MetricBackupToS3Records"},
			_jsii_.MemberMethod{JsiiMethod: "metricIncomingBytes", GoMethod: "MetricIncomingBytes"},
			_jsii_.MemberMethod{JsiiMethod: "metricIncomingRecords", GoMethod: "MetricIncomingRecords"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IDeliveryStream{}
			_jsii_.InitJsiiProxy(&j.Type__awsec2IConnectable)
			_jsii_.InitJsiiProxy(&j.Type__awsiamIGrantable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"cdk-extensions.kinesis_firehose.IDeliveryStreamBackupDestination",
		reflect.TypeOf((*IDeliveryStreamBackupDestination)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "renderBackupConfiguration", GoMethod: "RenderBackupConfiguration"},
		},
		func() interface{} {
			return &jsiiProxy_IDeliveryStreamBackupDestination{}
		},
	)
	_jsii_.RegisterClass(
		"cdk-extensions.kinesis_firehose.InputFormat",
		reflect.TypeOf((*InputFormat)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			return &jsiiProxy_InputFormat{}
		},
	)
	_jsii_.RegisterEnum(
		"cdk-extensions.kinesis_firehose.JsonParsingEngine",
		reflect.TypeOf((*JsonParsingEngine)(nil)).Elem(),
		map[string]interface{}{
			"JQ_1_6": JsonParsingEngine_JQ_1_6,
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.kinesis_firehose.JsonPartitioningOptions",
		reflect.TypeOf((*JsonPartitioningOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.kinesis_firehose.JsonPartitioningSource",
		reflect.TypeOf((*JsonPartitioningSource)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addPartition", GoMethod: "AddPartition"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
			_jsii_.MemberProperty{JsiiProperty: "retryInterval", GoGetter: "RetryInterval"},
		},
		func() interface{} {
			j := jsiiProxy_JsonPartitioningSource{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_DynamicPartitioning)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"cdk-extensions.kinesis_firehose.JsonQuery",
		reflect.TypeOf((*JsonQuery)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addField", GoMethod: "AddField"},
			_jsii_.MemberProperty{JsiiProperty: "query", GoGetter: "Query"},
			_jsii_.MemberMethod{JsiiMethod: "render", GoMethod: "Render"},
		},
		func() interface{} {
			j := jsiiProxy_JsonQuery{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_MetaDataExtractionQuery)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.kinesis_firehose.LambdaPartitioningOptions",
		reflect.TypeOf((*LambdaPartitioningOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.kinesis_firehose.LambdaPartitioningSource",
		reflect.TypeOf((*LambdaPartitioningSource)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
			_jsii_.MemberProperty{JsiiProperty: "lambdaFunction", GoGetter: "LambdaFunction"},
			_jsii_.MemberProperty{JsiiProperty: "retryInterval", GoGetter: "RetryInterval"},
		},
		func() interface{} {
			j := jsiiProxy_LambdaPartitioningSource{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_DynamicPartitioning)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"cdk-extensions.kinesis_firehose.LambdaProcessor",
		reflect.TypeOf((*LambdaProcessor)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addProcessorParameter", GoMethod: "AddProcessorParameter"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "lambdaFunction", GoGetter: "LambdaFunction"},
			_jsii_.MemberProperty{JsiiProperty: "processorType", GoGetter: "ProcessorType"},
		},
		func() interface{} {
			j := jsiiProxy_LambdaProcessor{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_DeliveryStreamProcessor)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.kinesis_firehose.LambdaProcessorOptions",
		reflect.TypeOf((*LambdaProcessorOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.kinesis_firehose.MetaDataExtractionQuery",
		reflect.TypeOf((*MetaDataExtractionQuery)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "query", GoGetter: "Query"},
			_jsii_.MemberMethod{JsiiMethod: "render", GoMethod: "Render"},
		},
		func() interface{} {
			return &jsiiProxy_MetaDataExtractionQuery{}
		},
	)
	_jsii_.RegisterClass(
		"cdk-extensions.kinesis_firehose.MetadataExtractionProcessor",
		reflect.TypeOf((*MetadataExtractionProcessor)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addProcessorParameter", GoMethod: "AddProcessorParameter"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "engine", GoGetter: "Engine"},
			_jsii_.MemberProperty{JsiiProperty: "processorType", GoGetter: "ProcessorType"},
			_jsii_.MemberProperty{JsiiProperty: "query", GoGetter: "Query"},
		},
		func() interface{} {
			j := jsiiProxy_MetadataExtractionProcessor{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_DeliveryStreamProcessor)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.kinesis_firehose.MetadataExtractionProcessorOptions",
		reflect.TypeOf((*MetadataExtractionProcessorOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.kinesis_firehose.OpenxJsonInputSerDe",
		reflect.TypeOf((*OpenxJsonInputSerDe)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addColumnKeyMapping", GoMethod: "AddColumnKeyMapping"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "caseInsensitive", GoGetter: "CaseInsensitive"},
			_jsii_.MemberProperty{JsiiProperty: "convertDotsToUnderscores", GoGetter: "ConvertDotsToUnderscores"},
		},
		func() interface{} {
			j := jsiiProxy_OpenxJsonInputSerDe{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_InputFormat)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.kinesis_firehose.OpenxJsonInputSerDeOptions",
		reflect.TypeOf((*OpenxJsonInputSerDeOptions)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cdk-extensions.kinesis_firehose.OrcCompressionFormat",
		reflect.TypeOf((*OrcCompressionFormat)(nil)).Elem(),
		map[string]interface{}{
			"NONE": OrcCompressionFormat_NONE,
			"SNAPPY": OrcCompressionFormat_SNAPPY,
			"ZLIB": OrcCompressionFormat_ZLIB,
		},
	)
	_jsii_.RegisterEnum(
		"cdk-extensions.kinesis_firehose.OrcFormatVersion",
		reflect.TypeOf((*OrcFormatVersion)(nil)).Elem(),
		map[string]interface{}{
			"V0_11": OrcFormatVersion_V0_11,
			"V0_12": OrcFormatVersion_V0_12,
		},
	)
	_jsii_.RegisterClass(
		"cdk-extensions.kinesis_firehose.OrcOutputSerDe",
		reflect.TypeOf((*OrcOutputSerDe)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addBloomFilterColumn", GoMethod: "AddBloomFilterColumn"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "blockSizeBytes", GoGetter: "BlockSizeBytes"},
			_jsii_.MemberProperty{JsiiProperty: "bloomFilterColumns", GoGetter: "BloomFilterColumns"},
			_jsii_.MemberProperty{JsiiProperty: "bloomFilterFalsePositiveProbability", GoGetter: "BloomFilterFalsePositiveProbability"},
			_jsii_.MemberProperty{JsiiProperty: "compression", GoGetter: "Compression"},
			_jsii_.MemberProperty{JsiiProperty: "dictionaryKeyThreshold", GoGetter: "DictionaryKeyThreshold"},
			_jsii_.MemberProperty{JsiiProperty: "enablePadding", GoGetter: "EnablePadding"},
			_jsii_.MemberProperty{JsiiProperty: "formatVersion", GoGetter: "FormatVersion"},
			_jsii_.MemberProperty{JsiiProperty: "paddingTolerance", GoGetter: "PaddingTolerance"},
			_jsii_.MemberProperty{JsiiProperty: "rowIndexStride", GoGetter: "RowIndexStride"},
			_jsii_.MemberProperty{JsiiProperty: "stripeSizeBytes", GoGetter: "StripeSizeBytes"},
		},
		func() interface{} {
			j := jsiiProxy_OrcOutputSerDe{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_OutputFormat)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.kinesis_firehose.OrcOutputSerDeOptions",
		reflect.TypeOf((*OrcOutputSerDeOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.kinesis_firehose.OutputFormat",
		reflect.TypeOf((*OutputFormat)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			return &jsiiProxy_OutputFormat{}
		},
	)
	_jsii_.RegisterEnum(
		"cdk-extensions.kinesis_firehose.ParquetCompressionFormat",
		reflect.TypeOf((*ParquetCompressionFormat)(nil)).Elem(),
		map[string]interface{}{
			"GZIP": ParquetCompressionFormat_GZIP,
			"SNAPPY": ParquetCompressionFormat_SNAPPY,
			"UNCOMPRESSED": ParquetCompressionFormat_UNCOMPRESSED,
		},
	)
	_jsii_.RegisterClass(
		"cdk-extensions.kinesis_firehose.ParquetOutputSerDe",
		reflect.TypeOf((*ParquetOutputSerDe)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "blockSizeBytes", GoGetter: "BlockSizeBytes"},
			_jsii_.MemberProperty{JsiiProperty: "compression", GoGetter: "Compression"},
			_jsii_.MemberProperty{JsiiProperty: "enableDictionaryCompression", GoGetter: "EnableDictionaryCompression"},
			_jsii_.MemberProperty{JsiiProperty: "maxPaddingBytes", GoGetter: "MaxPaddingBytes"},
			_jsii_.MemberProperty{JsiiProperty: "pageSizeBytes", GoGetter: "PageSizeBytes"},
			_jsii_.MemberProperty{JsiiProperty: "writerVersion", GoGetter: "WriterVersion"},
		},
		func() interface{} {
			j := jsiiProxy_ParquetOutputSerDe{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_OutputFormat)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.kinesis_firehose.ParquetOutputSerDeOptions",
		reflect.TypeOf((*ParquetOutputSerDeOptions)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cdk-extensions.kinesis_firehose.ParquetWriterVersion",
		reflect.TypeOf((*ParquetWriterVersion)(nil)).Elem(),
		map[string]interface{}{
			"V1": ParquetWriterVersion_V1,
			"V2": ParquetWriterVersion_V2,
		},
	)
	_jsii_.RegisterClass(
		"cdk-extensions.kinesis_firehose.ProcessorConfiguration",
		reflect.TypeOf((*ProcessorConfiguration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
			_jsii_.MemberProperty{JsiiProperty: "processors", GoGetter: "Processors"},
		},
		func() interface{} {
			return &jsiiProxy_ProcessorConfiguration{}
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.kinesis_firehose.ProcessorConfigurationOptions",
		reflect.TypeOf((*ProcessorConfigurationOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.kinesis_firehose.ProcessorConfigurationResult",
		reflect.TypeOf((*ProcessorConfigurationResult)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.kinesis_firehose.ProcessorType",
		reflect.TypeOf((*ProcessorType)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
		},
		func() interface{} {
			return &jsiiProxy_ProcessorType{}
		},
	)
	_jsii_.RegisterClass(
		"cdk-extensions.kinesis_firehose.RecordDeaggregationProcessor",
		reflect.TypeOf((*RecordDeaggregationProcessor)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addProcessorParameter", GoMethod: "AddProcessorParameter"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "delimiter", GoGetter: "Delimiter"},
			_jsii_.MemberProperty{JsiiProperty: "processorType", GoGetter: "ProcessorType"},
			_jsii_.MemberProperty{JsiiProperty: "subRecordType", GoGetter: "SubRecordType"},
		},
		func() interface{} {
			j := jsiiProxy_RecordDeaggregationProcessor{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_DeliveryStreamProcessor)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.kinesis_firehose.RecordDeaggregationProcessorOptions",
		reflect.TypeOf((*RecordDeaggregationProcessorOptions)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cdk-extensions.kinesis_firehose.S3CompressionFormat",
		reflect.TypeOf((*S3CompressionFormat)(nil)).Elem(),
		map[string]interface{}{
			"GZIP": S3CompressionFormat_GZIP,
			"HADOOP_SNAPPY": S3CompressionFormat_HADOOP_SNAPPY,
			"SNAPPY": S3CompressionFormat_SNAPPY,
			"UNCOMPRESSED": S3CompressionFormat_UNCOMPRESSED,
			"ZIP": S3CompressionFormat_ZIP,
		},
	)
	_jsii_.RegisterClass(
		"cdk-extensions.kinesis_firehose.S3Destination",
		reflect.TypeOf((*S3Destination)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "bucket", GoGetter: "Bucket"},
			_jsii_.MemberProperty{JsiiProperty: "buffering", GoGetter: "Buffering"},
			_jsii_.MemberMethod{JsiiMethod: "buildConfiguration", GoMethod: "BuildConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "cloudwatchLoggingConfiguration", GoGetter: "CloudwatchLoggingConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "compressionFormat", GoGetter: "CompressionFormat"},
			_jsii_.MemberProperty{JsiiProperty: "encryptionEnabled", GoGetter: "EncryptionEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "encryptionKey", GoGetter: "EncryptionKey"},
			_jsii_.MemberProperty{JsiiProperty: "errorOutputPrefix", GoGetter: "ErrorOutputPrefix"},
			_jsii_.MemberProperty{JsiiProperty: "keyPrefix", GoGetter: "KeyPrefix"},
			_jsii_.MemberMethod{JsiiMethod: "renderBackupConfiguration", GoMethod: "RenderBackupConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "role", GoGetter: "Role"},
		},
		func() interface{} {
			j := jsiiProxy_S3Destination{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_DeliveryStreamDestination)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IDeliveryStreamBackupDestination)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.kinesis_firehose.S3DestinationOptions",
		reflect.TypeOf((*S3DestinationOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.kinesis_firehose.SubRecordType",
		reflect.TypeOf((*SubRecordType)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
		},
		func() interface{} {
			return &jsiiProxy_SubRecordType{}
		},
	)
	_jsii_.RegisterClass(
		"cdk-extensions.kinesis_firehose.TableVersion",
		reflect.TypeOf((*TableVersion)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "version", GoGetter: "Version"},
		},
		func() interface{} {
			return &jsiiProxy_TableVersion{}
		},
	)
}
