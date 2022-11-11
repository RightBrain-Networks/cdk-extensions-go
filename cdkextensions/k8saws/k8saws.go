package k8saws

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"cdk-extensions.k8s_aws.AdotCollector",
		reflect.TypeOf((*AdotCollector)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "cluster", GoGetter: "Cluster"},
			_jsii_.MemberProperty{JsiiProperty: "createNamespace", GoGetter: "CreateNamespace"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "manifest", GoGetter: "Manifest"},
			_jsii_.MemberProperty{JsiiProperty: "namespace", GoGetter: "Namespace"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "serviceAccount", GoGetter: "ServiceAccount"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AdotCollector{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.k8s_aws.AdotCollectorProps",
		reflect.TypeOf((*AdotCollectorProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.k8s_aws.AppendedRecord",
		reflect.TypeOf((*AppendedRecord)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.k8s_aws.AwsSecretStore",
		reflect.TypeOf((*AwsSecretStore)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "cluster", GoGetter: "Cluster"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "manifest", GoGetter: "Manifest"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "namespace", GoGetter: "Namespace"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "secretStoreName", GoGetter: "SecretStoreName"},
			_jsii_.MemberProperty{JsiiProperty: "service", GoGetter: "Service"},
			_jsii_.MemberProperty{JsiiProperty: "serviceAccount", GoGetter: "ServiceAccount"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AwsSecretStore{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ISecretStore)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.k8s_aws.AwsSecretStoreProps",
		reflect.TypeOf((*AwsSecretStoreProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.k8s_aws.AwsServiceDiscoveryRegistry",
		reflect.TypeOf((*AwsServiceDiscoveryRegistry)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "registryType", GoGetter: "RegistryType"},
		},
		func() interface{} {
			j := jsiiProxy_AwsServiceDiscoveryRegistry{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IExternalDnsRegistry)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"cdk-extensions.k8s_aws.Echoserver",
		reflect.TypeOf((*Echoserver)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "cluster", GoGetter: "Cluster"},
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberProperty{JsiiProperty: "domainDiscovery", GoGetter: "DomainDiscovery"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "loadBalancerSubnets", GoGetter: "LoadBalancerSubnets"},
			_jsii_.MemberProperty{JsiiProperty: "manifest", GoGetter: "Manifest"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "namespace", GoGetter: "Namespace"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "port", GoGetter: "Port"},
			_jsii_.MemberMethod{JsiiMethod: "registerDomain", GoMethod: "RegisterDomain"},
			_jsii_.MemberProperty{JsiiProperty: "replicas", GoGetter: "Replicas"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "subdomain", GoGetter: "Subdomain"},
			_jsii_.MemberProperty{JsiiProperty: "tag", GoGetter: "Tag"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Echoserver{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.Type__awsec2IConnectable)
			_jsii_.InitJsiiProxy(&j.Type__route53IDnsResolvable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.k8s_aws.EchoserverProps",
		reflect.TypeOf((*EchoserverProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cdk-extensions.k8s_aws.ElasticsearchCompressionFormat",
		reflect.TypeOf((*ElasticsearchCompressionFormat)(nil)).Elem(),
		map[string]interface{}{
			"GZIP": ElasticsearchCompressionFormat_GZIP,
		},
	)
	_jsii_.RegisterClass(
		"cdk-extensions.k8s_aws.ElasticsearchOutputBufferSize",
		reflect.TypeOf((*ElasticsearchOutputBufferSize)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_ElasticsearchOutputBufferSize{}
		},
	)
	_jsii_.RegisterClass(
		"cdk-extensions.k8s_aws.EmitterStorageType",
		reflect.TypeOf((*EmitterStorageType)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
		},
		func() interface{} {
			return &jsiiProxy_EmitterStorageType{}
		},
	)
	_jsii_.RegisterEnum(
		"cdk-extensions.k8s_aws.ExternalDnsLogFormat",
		reflect.TypeOf((*ExternalDnsLogFormat)(nil)).Elem(),
		map[string]interface{}{
			"JSON": ExternalDnsLogFormat_JSON,
			"TEXT": ExternalDnsLogFormat_TEXT,
		},
	)
	_jsii_.RegisterEnum(
		"cdk-extensions.k8s_aws.ExternalDnsLogLevel",
		reflect.TypeOf((*ExternalDnsLogLevel)(nil)).Elem(),
		map[string]interface{}{
			"PANIC": ExternalDnsLogLevel_PANIC,
			"DEBUG": ExternalDnsLogLevel_DEBUG,
			"INFO": ExternalDnsLogLevel_INFO,
			"WARNING": ExternalDnsLogLevel_WARNING,
			"ERROR": ExternalDnsLogLevel_ERROR,
			"FATAL": ExternalDnsLogLevel_FATAL,
			"TRACE": ExternalDnsLogLevel_TRACE,
		},
	)
	_jsii_.RegisterClass(
		"cdk-extensions.k8s_aws.ExternalDnsRegistry",
		reflect.TypeOf((*ExternalDnsRegistry)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_ExternalDnsRegistry{}
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.k8s_aws.ExternalDnsRegistryConfiguration",
		reflect.TypeOf((*ExternalDnsRegistryConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cdk-extensions.k8s_aws.ExternalDnsSyncPolicy",
		reflect.TypeOf((*ExternalDnsSyncPolicy)(nil)).Elem(),
		map[string]interface{}{
			"SYNC": ExternalDnsSyncPolicy_SYNC,
			"UPSERT_ONLY": ExternalDnsSyncPolicy_UPSERT_ONLY,
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.k8s_aws.ExternalDnsZoneTag",
		reflect.TypeOf((*ExternalDnsZoneTag)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cdk-extensions.k8s_aws.ExternalDnsZoneType",
		reflect.TypeOf((*ExternalDnsZoneType)(nil)).Elem(),
		map[string]interface{}{
			"ALL": ExternalDnsZoneType_ALL,
			"PRIVATE": ExternalDnsZoneType_PRIVATE,
			"PUBLIC": ExternalDnsZoneType_PUBLIC,
		},
	)
	_jsii_.RegisterClass(
		"cdk-extensions.k8s_aws.ExternalSecret",
		reflect.TypeOf((*ExternalSecret)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addSecret", GoMethod: "AddSecret"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "cluster", GoGetter: "Cluster"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "manifest", GoGetter: "Manifest"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "namespace", GoGetter: "Namespace"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "refreshInterval", GoGetter: "RefreshInterval"},
			_jsii_.MemberProperty{JsiiProperty: "secretName", GoGetter: "SecretName"},
			_jsii_.MemberProperty{JsiiProperty: "secrets", GoGetter: "Secrets"},
			_jsii_.MemberProperty{JsiiProperty: "secretStore", GoGetter: "SecretStore"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ExternalSecret{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.k8s_aws.ExternalSecretOptions",
		reflect.TypeOf((*ExternalSecretOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.k8s_aws.ExternalSecretProps",
		reflect.TypeOf((*ExternalSecretProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.k8s_aws.ExternalSecretsOperator",
		reflect.TypeOf((*ExternalSecretsOperator)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "cluster", GoGetter: "Cluster"},
			_jsii_.MemberProperty{JsiiProperty: "createNamespace", GoGetter: "CreateNamespace"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "helmChart", GoGetter: "HelmChart"},
			_jsii_.MemberProperty{JsiiProperty: "namespace", GoGetter: "Namespace"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "registerSecretsManagerSecret", GoMethod: "RegisterSecretsManagerSecret"},
			_jsii_.MemberMethod{JsiiMethod: "registerSsmParameterSecret", GoMethod: "RegisterSsmParameterSecret"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ExternalSecretsOperator{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.k8s_aws.ExternalSecretsOperatorProps",
		reflect.TypeOf((*ExternalSecretsOperatorProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.k8s_aws.FargateLogger",
		reflect.TypeOf((*FargateLogger)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addFargateProfile", GoMethod: "AddFargateProfile"},
			_jsii_.MemberMethod{JsiiMethod: "addFilter", GoMethod: "AddFilter"},
			_jsii_.MemberMethod{JsiiMethod: "addOutput", GoMethod: "AddOutput"},
			_jsii_.MemberMethod{JsiiMethod: "addParser", GoMethod: "AddParser"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "cluster", GoGetter: "Cluster"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "filters", GoGetter: "Filters"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "manifest", GoGetter: "Manifest"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "outputs", GoGetter: "Outputs"},
			_jsii_.MemberProperty{JsiiProperty: "parsers", GoGetter: "Parsers"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_FargateLogger{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.k8s_aws.FargateLoggerOptions",
		reflect.TypeOf((*FargateLoggerOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.k8s_aws.FargateLoggerProps",
		reflect.TypeOf((*FargateLoggerProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.k8s_aws.FluentBitCloudWatchLogsOutput",
		reflect.TypeOf((*FluentBitCloudWatchLogsOutput)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "autoCreateGroup", GoGetter: "AutoCreateGroup"},
			_jsii_.MemberProperty{JsiiProperty: "autoRetryRequests", GoGetter: "AutoRetryRequests"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "endpoint", GoGetter: "Endpoint"},
			_jsii_.MemberProperty{JsiiProperty: "logFormat", GoGetter: "LogFormat"},
			_jsii_.MemberProperty{JsiiProperty: "logGroup", GoGetter: "LogGroup"},
			_jsii_.MemberProperty{JsiiProperty: "logGroupTemplate", GoGetter: "LogGroupTemplate"},
			_jsii_.MemberProperty{JsiiProperty: "logKey", GoGetter: "LogKey"},
			_jsii_.MemberProperty{JsiiProperty: "logRetention", GoGetter: "LogRetention"},
			_jsii_.MemberProperty{JsiiProperty: "logStream", GoGetter: "LogStream"},
			_jsii_.MemberProperty{JsiiProperty: "logStreamTemplate", GoGetter: "LogStreamTemplate"},
			_jsii_.MemberProperty{JsiiProperty: "match", GoGetter: "Match"},
			_jsii_.MemberProperty{JsiiProperty: "metricDimensions", GoGetter: "MetricDimensions"},
			_jsii_.MemberProperty{JsiiProperty: "metricNamespace", GoGetter: "MetricNamespace"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "pluginType", GoGetter: "PluginType"},
			_jsii_.MemberProperty{JsiiProperty: "region", GoGetter: "Region"},
			_jsii_.MemberMethod{JsiiMethod: "renderConfigFile", GoMethod: "RenderConfigFile"},
			_jsii_.MemberProperty{JsiiProperty: "role", GoGetter: "Role"},
			_jsii_.MemberProperty{JsiiProperty: "stsEndpoint", GoGetter: "StsEndpoint"},
		},
		func() interface{} {
			j := jsiiProxy_FluentBitCloudWatchLogsOutput{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_FluentBitOutputPluginBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.k8s_aws.FluentBitCloudWatchLogsOutputOptions",
		reflect.TypeOf((*FluentBitCloudWatchLogsOutputOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.k8s_aws.FluentBitElasticsearchOutput",
		reflect.TypeOf((*FluentBitElasticsearchOutput)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "awsAuth", GoGetter: "AwsAuth"},
			_jsii_.MemberProperty{JsiiProperty: "awsExternalId", GoGetter: "AwsExternalId"},
			_jsii_.MemberProperty{JsiiProperty: "awsRegion", GoGetter: "AwsRegion"},
			_jsii_.MemberProperty{JsiiProperty: "awsRole", GoGetter: "AwsRole"},
			_jsii_.MemberProperty{JsiiProperty: "awsStsEndpoint", GoGetter: "AwsStsEndpoint"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "bufferSize", GoGetter: "BufferSize"},
			_jsii_.MemberProperty{JsiiProperty: "cloudAuth", GoGetter: "CloudAuth"},
			_jsii_.MemberProperty{JsiiProperty: "cloudId", GoGetter: "CloudId"},
			_jsii_.MemberProperty{JsiiProperty: "compress", GoGetter: "Compress"},
			_jsii_.MemberProperty{JsiiProperty: "currentTimeIndex", GoGetter: "CurrentTimeIndex"},
			_jsii_.MemberProperty{JsiiProperty: "generateId", GoGetter: "GenerateId"},
			_jsii_.MemberProperty{JsiiProperty: "host", GoGetter: "Host"},
			_jsii_.MemberProperty{JsiiProperty: "httpPasswd", GoGetter: "HttpPasswd"},
			_jsii_.MemberProperty{JsiiProperty: "httpUser", GoGetter: "HttpUser"},
			_jsii_.MemberProperty{JsiiProperty: "idKey", GoGetter: "IdKey"},
			_jsii_.MemberProperty{JsiiProperty: "includeTagKey", GoGetter: "IncludeTagKey"},
			_jsii_.MemberProperty{JsiiProperty: "index", GoGetter: "Index"},
			_jsii_.MemberProperty{JsiiProperty: "logstashDateFormat", GoGetter: "LogstashDateFormat"},
			_jsii_.MemberProperty{JsiiProperty: "logstashFormat", GoGetter: "LogstashFormat"},
			_jsii_.MemberProperty{JsiiProperty: "logstashPrefix", GoGetter: "LogstashPrefix"},
			_jsii_.MemberProperty{JsiiProperty: "logstashPrefixKey", GoGetter: "LogstashPrefixKey"},
			_jsii_.MemberProperty{JsiiProperty: "match", GoGetter: "Match"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "path", GoGetter: "Path"},
			_jsii_.MemberProperty{JsiiProperty: "pipeline", GoGetter: "Pipeline"},
			_jsii_.MemberProperty{JsiiProperty: "pluginType", GoGetter: "PluginType"},
			_jsii_.MemberProperty{JsiiProperty: "port", GoGetter: "Port"},
			_jsii_.MemberMethod{JsiiMethod: "renderConfigFile", GoMethod: "RenderConfigFile"},
			_jsii_.MemberProperty{JsiiProperty: "replaceDots", GoGetter: "ReplaceDots"},
			_jsii_.MemberProperty{JsiiProperty: "suppressTypeName", GoGetter: "SuppressTypeName"},
			_jsii_.MemberProperty{JsiiProperty: "tagKey", GoGetter: "TagKey"},
			_jsii_.MemberProperty{JsiiProperty: "timeKey", GoGetter: "TimeKey"},
			_jsii_.MemberProperty{JsiiProperty: "timeKeyFormat", GoGetter: "TimeKeyFormat"},
			_jsii_.MemberProperty{JsiiProperty: "timeKeyNanos", GoGetter: "TimeKeyNanos"},
			_jsii_.MemberProperty{JsiiProperty: "traceError", GoGetter: "TraceError"},
			_jsii_.MemberProperty{JsiiProperty: "traceOutput", GoGetter: "TraceOutput"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "workers", GoGetter: "Workers"},
			_jsii_.MemberProperty{JsiiProperty: "writeOperation", GoGetter: "WriteOperation"},
		},
		func() interface{} {
			j := jsiiProxy_FluentBitElasticsearchOutput{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_FluentBitOutputPluginBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.k8s_aws.FluentBitElasticsearchOutputOptions",
		reflect.TypeOf((*FluentBitElasticsearchOutputOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.k8s_aws.FluentBitFilter",
		reflect.TypeOf((*FluentBitFilter)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_FluentBitFilter{}
		},
	)
	_jsii_.RegisterClass(
		"cdk-extensions.k8s_aws.FluentBitFilterPluginBase",
		reflect.TypeOf((*FluentBitFilterPluginBase)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "match", GoGetter: "Match"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "pluginType", GoGetter: "PluginType"},
			_jsii_.MemberMethod{JsiiMethod: "renderConfigFile", GoMethod: "RenderConfigFile"},
		},
		func() interface{} {
			j := jsiiProxy_FluentBitFilterPluginBase{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_FluentBitPlugin)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IFluentBitFilterPlugin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.k8s_aws.FluentBitFilterPluginCommonOptions",
		reflect.TypeOf((*FluentBitFilterPluginCommonOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.k8s_aws.FluentBitGrepFilter",
		reflect.TypeOf((*FluentBitGrepFilter)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "match", GoGetter: "Match"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "pattern", GoGetter: "Pattern"},
			_jsii_.MemberProperty{JsiiProperty: "pluginType", GoGetter: "PluginType"},
			_jsii_.MemberMethod{JsiiMethod: "renderConfigFile", GoMethod: "RenderConfigFile"},
		},
		func() interface{} {
			j := jsiiProxy_FluentBitGrepFilter{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_FluentBitFilterPluginBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.k8s_aws.FluentBitGrepFilterOptions",
		reflect.TypeOf((*FluentBitGrepFilterOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.k8s_aws.FluentBitGrepRegex",
		reflect.TypeOf((*FluentBitGrepRegex)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.k8s_aws.FluentBitJsonParser",
		reflect.TypeOf((*FluentBitJsonParser)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "format", GoGetter: "Format"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "pluginType", GoGetter: "PluginType"},
			_jsii_.MemberMethod{JsiiMethod: "renderConfigFile", GoMethod: "RenderConfigFile"},
			_jsii_.MemberProperty{JsiiProperty: "timeFormat", GoGetter: "TimeFormat"},
			_jsii_.MemberProperty{JsiiProperty: "timeKey", GoGetter: "TimeKey"},
		},
		func() interface{} {
			j := jsiiProxy_FluentBitJsonParser{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_FluentBitParserPluginBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.k8s_aws.FluentBitJsonParserOptions",
		reflect.TypeOf((*FluentBitJsonParserOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.k8s_aws.FluentBitKinesisFirehoseOutput",
		reflect.TypeOf((*FluentBitKinesisFirehoseOutput)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "autoRetryRequests", GoGetter: "AutoRetryRequests"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "compression", GoGetter: "Compression"},
			_jsii_.MemberProperty{JsiiProperty: "deliveryStream", GoGetter: "DeliveryStream"},
			_jsii_.MemberProperty{JsiiProperty: "endpoint", GoGetter: "Endpoint"},
			_jsii_.MemberProperty{JsiiProperty: "logKey", GoGetter: "LogKey"},
			_jsii_.MemberProperty{JsiiProperty: "match", GoGetter: "Match"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "pluginType", GoGetter: "PluginType"},
			_jsii_.MemberProperty{JsiiProperty: "region", GoGetter: "Region"},
			_jsii_.MemberMethod{JsiiMethod: "renderConfigFile", GoMethod: "RenderConfigFile"},
			_jsii_.MemberProperty{JsiiProperty: "role", GoGetter: "Role"},
			_jsii_.MemberProperty{JsiiProperty: "stsEndpoint", GoGetter: "StsEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "timeKey", GoGetter: "TimeKey"},
			_jsii_.MemberProperty{JsiiProperty: "timeKeyFormat", GoGetter: "TimeKeyFormat"},
		},
		func() interface{} {
			j := jsiiProxy_FluentBitKinesisFirehoseOutput{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_FluentBitOutputPluginBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.k8s_aws.FluentBitKinesisFirehoseOutputOptions",
		reflect.TypeOf((*FluentBitKinesisFirehoseOutputOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.k8s_aws.FluentBitKinesisOutput",
		reflect.TypeOf((*FluentBitKinesisOutput)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "autoRetryRequests", GoGetter: "AutoRetryRequests"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "endpoint", GoGetter: "Endpoint"},
			_jsii_.MemberProperty{JsiiProperty: "logKey", GoGetter: "LogKey"},
			_jsii_.MemberProperty{JsiiProperty: "match", GoGetter: "Match"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "pluginType", GoGetter: "PluginType"},
			_jsii_.MemberProperty{JsiiProperty: "region", GoGetter: "Region"},
			_jsii_.MemberMethod{JsiiMethod: "renderConfigFile", GoMethod: "RenderConfigFile"},
			_jsii_.MemberProperty{JsiiProperty: "role", GoGetter: "Role"},
			_jsii_.MemberProperty{JsiiProperty: "stream", GoGetter: "Stream"},
			_jsii_.MemberProperty{JsiiProperty: "stsEndpoint", GoGetter: "StsEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "timeKey", GoGetter: "TimeKey"},
			_jsii_.MemberProperty{JsiiProperty: "timeKeyFormat", GoGetter: "TimeKeyFormat"},
		},
		func() interface{} {
			j := jsiiProxy_FluentBitKinesisOutput{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_FluentBitOutputPluginBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.k8s_aws.FluentBitKinesisOutputOptions",
		reflect.TypeOf((*FluentBitKinesisOutputOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.k8s_aws.FluentBitKubernetesFilter",
		reflect.TypeOf((*FluentBitKubernetesFilter)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "annotations", GoGetter: "Annotations"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "bufferSize", GoGetter: "BufferSize"},
			_jsii_.MemberProperty{JsiiProperty: "cacheUseDockerId", GoGetter: "CacheUseDockerId"},
			_jsii_.MemberProperty{JsiiProperty: "dnsRetries", GoGetter: "DnsRetries"},
			_jsii_.MemberProperty{JsiiProperty: "dnsWaitTime", GoGetter: "DnsWaitTime"},
			_jsii_.MemberProperty{JsiiProperty: "dummyMeta", GoGetter: "DummyMeta"},
			_jsii_.MemberProperty{JsiiProperty: "k8sLoggingExclude", GoGetter: "K8sLoggingExclude"},
			_jsii_.MemberProperty{JsiiProperty: "k8sLoggingParser", GoGetter: "K8sLoggingParser"},
			_jsii_.MemberProperty{JsiiProperty: "keepLog", GoGetter: "KeepLog"},
			_jsii_.MemberProperty{JsiiProperty: "kubeCaFile", GoGetter: "KubeCaFile"},
			_jsii_.MemberProperty{JsiiProperty: "kubeCaPath", GoGetter: "KubeCaPath"},
			_jsii_.MemberProperty{JsiiProperty: "kubeletHost", GoGetter: "KubeletHost"},
			_jsii_.MemberProperty{JsiiProperty: "kubeletPort", GoGetter: "KubeletPort"},
			_jsii_.MemberProperty{JsiiProperty: "kubeMetaCacheTtl", GoGetter: "KubeMetaCacheTtl"},
			_jsii_.MemberProperty{JsiiProperty: "kubeMetaPreloadCacheDir", GoGetter: "KubeMetaPreloadCacheDir"},
			_jsii_.MemberProperty{JsiiProperty: "kubeTagPrefix", GoGetter: "KubeTagPrefix"},
			_jsii_.MemberProperty{JsiiProperty: "kubeTokenCommand", GoGetter: "KubeTokenCommand"},
			_jsii_.MemberProperty{JsiiProperty: "kubeTokenFile", GoGetter: "KubeTokenFile"},
			_jsii_.MemberProperty{JsiiProperty: "kubeTokenTtl", GoGetter: "KubeTokenTtl"},
			_jsii_.MemberProperty{JsiiProperty: "kubeUrl", GoGetter: "KubeUrl"},
			_jsii_.MemberProperty{JsiiProperty: "labels", GoGetter: "Labels"},
			_jsii_.MemberProperty{JsiiProperty: "match", GoGetter: "Match"},
			_jsii_.MemberProperty{JsiiProperty: "mergeLog", GoGetter: "MergeLog"},
			_jsii_.MemberProperty{JsiiProperty: "mergeLogKey", GoGetter: "MergeLogKey"},
			_jsii_.MemberProperty{JsiiProperty: "mergeLogTrim", GoGetter: "MergeLogTrim"},
			_jsii_.MemberProperty{JsiiProperty: "mergeParser", GoGetter: "MergeParser"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "pluginType", GoGetter: "PluginType"},
			_jsii_.MemberProperty{JsiiProperty: "regexParser", GoGetter: "RegexParser"},
			_jsii_.MemberMethod{JsiiMethod: "renderConfigFile", GoMethod: "RenderConfigFile"},
			_jsii_.MemberProperty{JsiiProperty: "tlsDebug", GoGetter: "TlsDebug"},
			_jsii_.MemberProperty{JsiiProperty: "tlsVerify", GoGetter: "TlsVerify"},
			_jsii_.MemberProperty{JsiiProperty: "useJournal", GoGetter: "UseJournal"},
			_jsii_.MemberProperty{JsiiProperty: "useKubelet", GoGetter: "UseKubelet"},
		},
		func() interface{} {
			j := jsiiProxy_FluentBitKubernetesFilter{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_FluentBitFilterPluginBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.k8s_aws.FluentBitKubernetesFilterOptions",
		reflect.TypeOf((*FluentBitKubernetesFilterOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.k8s_aws.FluentBitLogGroupOutput",
		reflect.TypeOf((*FluentBitLogGroupOutput)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "autoCreate", GoGetter: "AutoCreate"},
			_jsii_.MemberProperty{JsiiProperty: "logGroup", GoGetter: "LogGroup"},
			_jsii_.MemberProperty{JsiiProperty: "logGroupName", GoGetter: "LogGroupName"},
		},
		func() interface{} {
			return &jsiiProxy_FluentBitLogGroupOutput{}
		},
	)
	_jsii_.RegisterClass(
		"cdk-extensions.k8s_aws.FluentBitLogStreamOutput",
		reflect.TypeOf((*FluentBitLogStreamOutput)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "logStreamName", GoGetter: "LogStreamName"},
			_jsii_.MemberProperty{JsiiProperty: "logStreamPrefix", GoGetter: "LogStreamPrefix"},
		},
		func() interface{} {
			return &jsiiProxy_FluentBitLogStreamOutput{}
		},
	)
	_jsii_.RegisterClass(
		"cdk-extensions.k8s_aws.FluentBitLogfmtParser",
		reflect.TypeOf((*FluentBitLogfmtParser)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "format", GoGetter: "Format"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "pluginType", GoGetter: "PluginType"},
			_jsii_.MemberMethod{JsiiMethod: "renderConfigFile", GoMethod: "RenderConfigFile"},
			_jsii_.MemberProperty{JsiiProperty: "timeFormat", GoGetter: "TimeFormat"},
			_jsii_.MemberProperty{JsiiProperty: "timeKey", GoGetter: "TimeKey"},
			_jsii_.MemberProperty{JsiiProperty: "types", GoGetter: "Types"},
		},
		func() interface{} {
			j := jsiiProxy_FluentBitLogfmtParser{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_FluentBitParserPluginBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.k8s_aws.FluentBitLogfmtParserOptions",
		reflect.TypeOf((*FluentBitLogfmtParserOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.k8s_aws.FluentBitLtsvParser",
		reflect.TypeOf((*FluentBitLtsvParser)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "format", GoGetter: "Format"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "pluginType", GoGetter: "PluginType"},
			_jsii_.MemberMethod{JsiiMethod: "renderConfigFile", GoMethod: "RenderConfigFile"},
			_jsii_.MemberProperty{JsiiProperty: "timeFormat", GoGetter: "TimeFormat"},
			_jsii_.MemberProperty{JsiiProperty: "timeKey", GoGetter: "TimeKey"},
			_jsii_.MemberProperty{JsiiProperty: "types", GoGetter: "Types"},
		},
		func() interface{} {
			j := jsiiProxy_FluentBitLtsvParser{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_FluentBitParserPluginBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.k8s_aws.FluentBitLtsvParserOptions",
		reflect.TypeOf((*FluentBitLtsvParserOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.k8s_aws.FluentBitMatch",
		reflect.TypeOf((*FluentBitMatch)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "evaluator", GoGetter: "Evaluator"},
			_jsii_.MemberProperty{JsiiProperty: "pattern", GoGetter: "Pattern"},
			_jsii_.MemberMethod{JsiiMethod: "toObject", GoMethod: "ToObject"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			return &jsiiProxy_FluentBitMatch{}
		},
	)
	_jsii_.RegisterEnum(
		"cdk-extensions.k8s_aws.FluentBitMatchEvaluator",
		reflect.TypeOf((*FluentBitMatchEvaluator)(nil)).Elem(),
		map[string]interface{}{
			"GLOB": FluentBitMatchEvaluator_GLOB,
			"REGEX": FluentBitMatchEvaluator_REGEX,
		},
	)
	_jsii_.RegisterClass(
		"cdk-extensions.k8s_aws.FluentBitModifyFilter",
		reflect.TypeOf((*FluentBitModifyFilter)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addCondition", GoMethod: "AddCondition"},
			_jsii_.MemberMethod{JsiiMethod: "addOperation", GoMethod: "AddOperation"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "conditions", GoGetter: "Conditions"},
			_jsii_.MemberProperty{JsiiProperty: "match", GoGetter: "Match"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "operations", GoGetter: "Operations"},
			_jsii_.MemberProperty{JsiiProperty: "pluginType", GoGetter: "PluginType"},
			_jsii_.MemberMethod{JsiiMethod: "renderConfigFile", GoMethod: "RenderConfigFile"},
		},
		func() interface{} {
			j := jsiiProxy_FluentBitModifyFilter{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_FluentBitFilterPluginBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.k8s_aws.FluentBitModifyFilterOptions",
		reflect.TypeOf((*FluentBitModifyFilterOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.k8s_aws.FluentBitNestFilter",
		reflect.TypeOf((*FluentBitNestFilter)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "addPrefix", GoGetter: "AddPrefix"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "match", GoGetter: "Match"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "operation", GoGetter: "Operation"},
			_jsii_.MemberProperty{JsiiProperty: "pluginType", GoGetter: "PluginType"},
			_jsii_.MemberProperty{JsiiProperty: "removePrefix", GoGetter: "RemovePrefix"},
			_jsii_.MemberMethod{JsiiMethod: "renderConfigFile", GoMethod: "RenderConfigFile"},
		},
		func() interface{} {
			j := jsiiProxy_FluentBitNestFilter{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_FluentBitFilterPluginBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.k8s_aws.FluentBitNestFilterOptions",
		reflect.TypeOf((*FluentBitNestFilterOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.k8s_aws.FluentBitOpenSearchOutput",
		reflect.TypeOf((*FluentBitOpenSearchOutput)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "awsAuth", GoGetter: "AwsAuth"},
			_jsii_.MemberProperty{JsiiProperty: "awsExternalId", GoGetter: "AwsExternalId"},
			_jsii_.MemberProperty{JsiiProperty: "awsRegion", GoGetter: "AwsRegion"},
			_jsii_.MemberProperty{JsiiProperty: "awsRole", GoGetter: "AwsRole"},
			_jsii_.MemberProperty{JsiiProperty: "awsStsEndpoint", GoGetter: "AwsStsEndpoint"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "bufferSize", GoGetter: "BufferSize"},
			_jsii_.MemberProperty{JsiiProperty: "currentTimeIndex", GoGetter: "CurrentTimeIndex"},
			_jsii_.MemberProperty{JsiiProperty: "domain", GoGetter: "Domain"},
			_jsii_.MemberProperty{JsiiProperty: "generateId", GoGetter: "GenerateId"},
			_jsii_.MemberProperty{JsiiProperty: "httpPasswd", GoGetter: "HttpPasswd"},
			_jsii_.MemberProperty{JsiiProperty: "httpUser", GoGetter: "HttpUser"},
			_jsii_.MemberProperty{JsiiProperty: "idKey", GoGetter: "IdKey"},
			_jsii_.MemberProperty{JsiiProperty: "includeTagKey", GoGetter: "IncludeTagKey"},
			_jsii_.MemberProperty{JsiiProperty: "index", GoGetter: "Index"},
			_jsii_.MemberProperty{JsiiProperty: "logstashDateFormat", GoGetter: "LogstashDateFormat"},
			_jsii_.MemberProperty{JsiiProperty: "logstashFormat", GoGetter: "LogstashFormat"},
			_jsii_.MemberProperty{JsiiProperty: "logstashPrefix", GoGetter: "LogstashPrefix"},
			_jsii_.MemberProperty{JsiiProperty: "logstashPrefixKey", GoGetter: "LogstashPrefixKey"},
			_jsii_.MemberProperty{JsiiProperty: "match", GoGetter: "Match"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "path", GoGetter: "Path"},
			_jsii_.MemberProperty{JsiiProperty: "pipeline", GoGetter: "Pipeline"},
			_jsii_.MemberProperty{JsiiProperty: "pluginType", GoGetter: "PluginType"},
			_jsii_.MemberProperty{JsiiProperty: "port", GoGetter: "Port"},
			_jsii_.MemberMethod{JsiiMethod: "renderConfigFile", GoMethod: "RenderConfigFile"},
			_jsii_.MemberProperty{JsiiProperty: "replaceDots", GoGetter: "ReplaceDots"},
			_jsii_.MemberProperty{JsiiProperty: "suppressTypeName", GoGetter: "SuppressTypeName"},
			_jsii_.MemberProperty{JsiiProperty: "tagKey", GoGetter: "TagKey"},
			_jsii_.MemberProperty{JsiiProperty: "timeKey", GoGetter: "TimeKey"},
			_jsii_.MemberProperty{JsiiProperty: "timeKeyFormat", GoGetter: "TimeKeyFormat"},
			_jsii_.MemberProperty{JsiiProperty: "timeKeyNanos", GoGetter: "TimeKeyNanos"},
			_jsii_.MemberProperty{JsiiProperty: "traceError", GoGetter: "TraceError"},
			_jsii_.MemberProperty{JsiiProperty: "traceOutput", GoGetter: "TraceOutput"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "workers", GoGetter: "Workers"},
			_jsii_.MemberProperty{JsiiProperty: "writeOperation", GoGetter: "WriteOperation"},
		},
		func() interface{} {
			j := jsiiProxy_FluentBitOpenSearchOutput{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_FluentBitOutputPluginBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.k8s_aws.FluentBitOpenSearchOutputOptions",
		reflect.TypeOf((*FluentBitOpenSearchOutputOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.k8s_aws.FluentBitOutput",
		reflect.TypeOf((*FluentBitOutput)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_FluentBitOutput{}
		},
	)
	_jsii_.RegisterClass(
		"cdk-extensions.k8s_aws.FluentBitOutputPluginBase",
		reflect.TypeOf((*FluentBitOutputPluginBase)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "match", GoGetter: "Match"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "pluginType", GoGetter: "PluginType"},
			_jsii_.MemberMethod{JsiiMethod: "renderConfigFile", GoMethod: "RenderConfigFile"},
		},
		func() interface{} {
			j := jsiiProxy_FluentBitOutputPluginBase{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_FluentBitPlugin)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IFluentBitOutputPlugin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.k8s_aws.FluentBitOutputPluginCommonOptions",
		reflect.TypeOf((*FluentBitOutputPluginCommonOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.k8s_aws.FluentBitParser",
		reflect.TypeOf((*FluentBitParser)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_FluentBitParser{}
		},
	)
	_jsii_.RegisterClass(
		"cdk-extensions.k8s_aws.FluentBitParserFilter",
		reflect.TypeOf((*FluentBitParserFilter)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addParser", GoMethod: "AddParser"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "keyName", GoGetter: "KeyName"},
			_jsii_.MemberProperty{JsiiProperty: "match", GoGetter: "Match"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "parsers", GoGetter: "Parsers"},
			_jsii_.MemberProperty{JsiiProperty: "pluginType", GoGetter: "PluginType"},
			_jsii_.MemberProperty{JsiiProperty: "preserveKey", GoGetter: "PreserveKey"},
			_jsii_.MemberMethod{JsiiMethod: "renderConfigFile", GoMethod: "RenderConfigFile"},
			_jsii_.MemberProperty{JsiiProperty: "reserveData", GoGetter: "ReserveData"},
		},
		func() interface{} {
			j := jsiiProxy_FluentBitParserFilter{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_FluentBitFilterPluginBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.k8s_aws.FluentBitParserFilterOptions",
		reflect.TypeOf((*FluentBitParserFilterOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.k8s_aws.FluentBitParserPluginBase",
		reflect.TypeOf((*FluentBitParserPluginBase)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "format", GoGetter: "Format"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "pluginType", GoGetter: "PluginType"},
			_jsii_.MemberMethod{JsiiMethod: "renderConfigFile", GoMethod: "RenderConfigFile"},
		},
		func() interface{} {
			j := jsiiProxy_FluentBitParserPluginBase{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_FluentBitPlugin)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IFluentBitParserPlugin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.k8s_aws.FluentBitParserPluginCommonOptions",
		reflect.TypeOf((*FluentBitParserPluginCommonOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.k8s_aws.FluentBitPlugin",
		reflect.TypeOf((*FluentBitPlugin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "pluginType", GoGetter: "PluginType"},
			_jsii_.MemberMethod{JsiiMethod: "renderConfigFile", GoMethod: "RenderConfigFile"},
		},
		func() interface{} {
			j := jsiiProxy_FluentBitPlugin{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IFluentBitPlugin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.k8s_aws.FluentBitPluginCommonOptions",
		reflect.TypeOf((*FluentBitPluginCommonOptions)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cdk-extensions.k8s_aws.FluentBitPluginType",
		reflect.TypeOf((*FluentBitPluginType)(nil)).Elem(),
		map[string]interface{}{
			"FILTER": FluentBitPluginType_FILTER,
			"OUTPUT": FluentBitPluginType_OUTPUT,
			"PARSER": FluentBitPluginType_PARSER,
		},
	)
	_jsii_.RegisterClass(
		"cdk-extensions.k8s_aws.FluentBitRecordModifierFilter",
		reflect.TypeOf((*FluentBitRecordModifierFilter)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addAllow", GoMethod: "AddAllow"},
			_jsii_.MemberMethod{JsiiMethod: "addRecord", GoMethod: "AddRecord"},
			_jsii_.MemberMethod{JsiiMethod: "addRemove", GoMethod: "AddRemove"},
			_jsii_.MemberProperty{JsiiProperty: "allow", GoGetter: "Allow"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "match", GoGetter: "Match"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "pluginType", GoGetter: "PluginType"},
			_jsii_.MemberProperty{JsiiProperty: "records", GoGetter: "Records"},
			_jsii_.MemberProperty{JsiiProperty: "remove", GoGetter: "Remove"},
			_jsii_.MemberMethod{JsiiMethod: "renderConfigFile", GoMethod: "RenderConfigFile"},
		},
		func() interface{} {
			j := jsiiProxy_FluentBitRecordModifierFilter{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_FluentBitFilterPluginBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.k8s_aws.FluentBitRecordModifierFilterOptions",
		reflect.TypeOf((*FluentBitRecordModifierFilterOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.k8s_aws.FluentBitRegexParser",
		reflect.TypeOf((*FluentBitRegexParser)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "format", GoGetter: "Format"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "pluginType", GoGetter: "PluginType"},
			_jsii_.MemberProperty{JsiiProperty: "regex", GoGetter: "Regex"},
			_jsii_.MemberMethod{JsiiMethod: "renderConfigFile", GoMethod: "RenderConfigFile"},
			_jsii_.MemberProperty{JsiiProperty: "skipEmptyValues", GoGetter: "SkipEmptyValues"},
			_jsii_.MemberProperty{JsiiProperty: "timeFormat", GoGetter: "TimeFormat"},
			_jsii_.MemberProperty{JsiiProperty: "timeKey", GoGetter: "TimeKey"},
			_jsii_.MemberProperty{JsiiProperty: "types", GoGetter: "Types"},
		},
		func() interface{} {
			j := jsiiProxy_FluentBitRegexParser{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_FluentBitParserPluginBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.k8s_aws.FluentBitRegexParserOptions",
		reflect.TypeOf((*FluentBitRegexParserOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.k8s_aws.FluentBitRewriteTagFilter",
		reflect.TypeOf((*FluentBitRewriteTagFilter)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addRule", GoMethod: "AddRule"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "emitterMemBufLimit", GoGetter: "EmitterMemBufLimit"},
			_jsii_.MemberProperty{JsiiProperty: "emitterName", GoGetter: "EmitterName"},
			_jsii_.MemberProperty{JsiiProperty: "emitterStorageType", GoGetter: "EmitterStorageType"},
			_jsii_.MemberProperty{JsiiProperty: "match", GoGetter: "Match"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "pluginType", GoGetter: "PluginType"},
			_jsii_.MemberMethod{JsiiMethod: "renderConfigFile", GoMethod: "RenderConfigFile"},
			_jsii_.MemberProperty{JsiiProperty: "rules", GoGetter: "Rules"},
		},
		func() interface{} {
			j := jsiiProxy_FluentBitRewriteTagFilter{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_FluentBitFilterPluginBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.k8s_aws.FluentBitRewriteTagFilterOptions",
		reflect.TypeOf((*FluentBitRewriteTagFilterOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.k8s_aws.FluentBitThrottleFilter",
		reflect.TypeOf((*FluentBitThrottleFilter)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "interval", GoGetter: "Interval"},
			_jsii_.MemberProperty{JsiiProperty: "match", GoGetter: "Match"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "pluginType", GoGetter: "PluginType"},
			_jsii_.MemberProperty{JsiiProperty: "printStatus", GoGetter: "PrintStatus"},
			_jsii_.MemberProperty{JsiiProperty: "rate", GoGetter: "Rate"},
			_jsii_.MemberMethod{JsiiMethod: "renderConfigFile", GoMethod: "RenderConfigFile"},
			_jsii_.MemberProperty{JsiiProperty: "window", GoGetter: "Window"},
		},
		func() interface{} {
			j := jsiiProxy_FluentBitThrottleFilter{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_FluentBitFilterPluginBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.k8s_aws.FluentBitThrottleFilterOptions",
		reflect.TypeOf((*FluentBitThrottleFilterOptions)(nil)).Elem(),
	)
	_jsii_.RegisterInterface(
		"cdk-extensions.k8s_aws.IExternalDnsRegistry",
		reflect.TypeOf((*IExternalDnsRegistry)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "registryType", GoGetter: "RegistryType"},
		},
		func() interface{} {
			return &jsiiProxy_IExternalDnsRegistry{}
		},
	)
	_jsii_.RegisterInterface(
		"cdk-extensions.k8s_aws.IFluentBitFilterPlugin",
		reflect.TypeOf((*IFluentBitFilterPlugin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "pluginType", GoGetter: "PluginType"},
		},
		func() interface{} {
			j := jsiiProxy_IFluentBitFilterPlugin{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IFluentBitPlugin)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"cdk-extensions.k8s_aws.IFluentBitOutputPlugin",
		reflect.TypeOf((*IFluentBitOutputPlugin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "pluginType", GoGetter: "PluginType"},
		},
		func() interface{} {
			j := jsiiProxy_IFluentBitOutputPlugin{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IFluentBitPlugin)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"cdk-extensions.k8s_aws.IFluentBitParserPlugin",
		reflect.TypeOf((*IFluentBitParserPlugin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "format", GoGetter: "Format"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "pluginType", GoGetter: "PluginType"},
		},
		func() interface{} {
			j := jsiiProxy_IFluentBitParserPlugin{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IFluentBitPlugin)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"cdk-extensions.k8s_aws.IFluentBitPlugin",
		reflect.TypeOf((*IFluentBitPlugin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "pluginType", GoGetter: "PluginType"},
		},
		func() interface{} {
			return &jsiiProxy_IFluentBitPlugin{}
		},
	)
	_jsii_.RegisterInterface(
		"cdk-extensions.k8s_aws.INestFilterOperation",
		reflect.TypeOf((*INestFilterOperation)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "fields", GoGetter: "Fields"},
			_jsii_.MemberProperty{JsiiProperty: "operation", GoGetter: "Operation"},
		},
		func() interface{} {
			return &jsiiProxy_INestFilterOperation{}
		},
	)
	_jsii_.RegisterInterface(
		"cdk-extensions.k8s_aws.ISecretReference",
		reflect.TypeOf((*ISecretReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			return &jsiiProxy_ISecretReference{}
		},
	)
	_jsii_.RegisterInterface(
		"cdk-extensions.k8s_aws.ISecretStore",
		reflect.TypeOf((*ISecretStore)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "secretStoreName", GoGetter: "SecretStoreName"},
		},
		func() interface{} {
			j := jsiiProxy_ISecretStore{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIDependable)
			return &j
		},
	)
	_jsii_.RegisterEnum(
		"cdk-extensions.k8s_aws.KinesisFirehoseCompressionFormat",
		reflect.TypeOf((*KinesisFirehoseCompressionFormat)(nil)).Elem(),
		map[string]interface{}{
			"ARROW": KinesisFirehoseCompressionFormat_ARROW,
			"GZIP": KinesisFirehoseCompressionFormat_GZIP,
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.k8s_aws.LiftOptions",
		reflect.TypeOf((*LiftOptions)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cdk-extensions.k8s_aws.MetadataPolicy",
		reflect.TypeOf((*MetadataPolicy)(nil)).Elem(),
		map[string]interface{}{
			"FETCH": MetadataPolicy_FETCH,
			"NONE": MetadataPolicy_NONE,
		},
	)
	_jsii_.RegisterClass(
		"cdk-extensions.k8s_aws.ModifyCondition",
		reflect.TypeOf((*ModifyCondition)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "args", GoGetter: "Args"},
			_jsii_.MemberProperty{JsiiProperty: "condition", GoGetter: "Condition"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			return &jsiiProxy_ModifyCondition{}
		},
	)
	_jsii_.RegisterClass(
		"cdk-extensions.k8s_aws.ModifyOperation",
		reflect.TypeOf((*ModifyOperation)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "args", GoGetter: "Args"},
			_jsii_.MemberProperty{JsiiProperty: "operation", GoGetter: "Operation"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			return &jsiiProxy_ModifyOperation{}
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.k8s_aws.NamespacedExternalSecretOptions",
		reflect.TypeOf((*NamespacedExternalSecretOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.k8s_aws.NestFilterOperation",
		reflect.TypeOf((*NestFilterOperation)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "fields", GoGetter: "Fields"},
			_jsii_.MemberProperty{JsiiProperty: "operation", GoGetter: "Operation"},
		},
		func() interface{} {
			j := jsiiProxy_NestFilterOperation{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_INestFilterOperation)
			return &j
		},
	)
	_jsii_.RegisterEnum(
		"cdk-extensions.k8s_aws.NestFilterOperationType",
		reflect.TypeOf((*NestFilterOperationType)(nil)).Elem(),
		map[string]interface{}{
			"LIFT": NestFilterOperationType_LIFT,
			"NEST": NestFilterOperationType_NEST,
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.k8s_aws.NestOptions",
		reflect.TypeOf((*NestOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.k8s_aws.NoopRegistry",
		reflect.TypeOf((*NoopRegistry)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "registryType", GoGetter: "RegistryType"},
		},
		func() interface{} {
			j := jsiiProxy_NoopRegistry{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IExternalDnsRegistry)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"cdk-extensions.k8s_aws.OpenSearchOutputBufferSize",
		reflect.TypeOf((*OpenSearchOutputBufferSize)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_OpenSearchOutputBufferSize{}
		},
	)
	_jsii_.RegisterClass(
		"cdk-extensions.k8s_aws.ParserPluginDataType",
		reflect.TypeOf((*ParserPluginDataType)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
		},
		func() interface{} {
			return &jsiiProxy_ParserPluginDataType{}
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.k8s_aws.ResolvedFluentBitConfiguration",
		reflect.TypeOf((*ResolvedFluentBitConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.k8s_aws.RewriteTagRule",
		reflect.TypeOf((*RewriteTagRule)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.k8s_aws.Route53Dns",
		reflect.TypeOf((*Route53Dns)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDomainFilter", GoMethod: "AddDomainFilter"},
			_jsii_.MemberMethod{JsiiMethod: "addZoneTag", GoMethod: "AddZoneTag"},
			_jsii_.MemberProperty{JsiiProperty: "apiRetries", GoGetter: "ApiRetries"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "batchChangeSize", GoGetter: "BatchChangeSize"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "cluster", GoGetter: "Cluster"},
			_jsii_.MemberProperty{JsiiProperty: "domainFilter", GoGetter: "DomainFilter"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "evaluateTargetHealth", GoGetter: "EvaluateTargetHealth"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "logFormat", GoGetter: "LogFormat"},
			_jsii_.MemberProperty{JsiiProperty: "logLevel", GoGetter: "LogLevel"},
			_jsii_.MemberProperty{JsiiProperty: "namespace", GoGetter: "Namespace"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "preferCname", GoGetter: "PreferCname"},
			_jsii_.MemberProperty{JsiiProperty: "recordOwnershipRegistry", GoGetter: "RecordOwnershipRegistry"},
			_jsii_.MemberProperty{JsiiProperty: "region", GoGetter: "Region"},
			_jsii_.MemberProperty{JsiiProperty: "replicaCount", GoGetter: "ReplicaCount"},
			_jsii_.MemberProperty{JsiiProperty: "serviceAccount", GoGetter: "ServiceAccount"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "syncPolicy", GoGetter: "SyncPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "zoneTags", GoGetter: "ZoneTags"},
			_jsii_.MemberProperty{JsiiProperty: "zoneType", GoGetter: "ZoneType"},
		},
		func() interface{} {
			j := jsiiProxy_Route53Dns{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.k8s_aws.Route53DnsOptions",
		reflect.TypeOf((*Route53DnsOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.k8s_aws.Route53DnsProps",
		reflect.TypeOf((*Route53DnsProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.k8s_aws.SecretFieldReference",
		reflect.TypeOf((*SecretFieldReference)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.k8s_aws.SecretReferenceConfiguration",
		reflect.TypeOf((*SecretReferenceConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.k8s_aws.SecretsManagerReference",
		reflect.TypeOf((*SecretsManagerReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addFieldMapping", GoMethod: "AddFieldMapping"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "fields", GoGetter: "Fields"},
			_jsii_.MemberProperty{JsiiProperty: "secret", GoGetter: "Secret"},
		},
		func() interface{} {
			j := jsiiProxy_SecretsManagerReference{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ISecretReference)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.k8s_aws.SecretsManagerReferenceOptions",
		reflect.TypeOf((*SecretsManagerReferenceOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.k8s_aws.SecretsManagerSecretStore",
		reflect.TypeOf((*SecretsManagerSecretStore)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addSecret", GoMethod: "AddSecret"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "cluster", GoGetter: "Cluster"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "manifest", GoGetter: "Manifest"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "namespace", GoGetter: "Namespace"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "secretStoreName", GoGetter: "SecretStoreName"},
			_jsii_.MemberProperty{JsiiProperty: "service", GoGetter: "Service"},
			_jsii_.MemberProperty{JsiiProperty: "serviceAccount", GoGetter: "ServiceAccount"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_SecretsManagerSecretStore{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_AwsSecretStore)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.k8s_aws.SecretsManagerSecretStoreProps",
		reflect.TypeOf((*SecretsManagerSecretStoreProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.k8s_aws.SsmParameterReference",
		reflect.TypeOf((*SsmParameterReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addFieldMapping", GoMethod: "AddFieldMapping"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "fields", GoGetter: "Fields"},
			_jsii_.MemberProperty{JsiiProperty: "parameter", GoGetter: "Parameter"},
		},
		func() interface{} {
			j := jsiiProxy_SsmParameterReference{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ISecretReference)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.k8s_aws.SsmParameterReferenceOptions",
		reflect.TypeOf((*SsmParameterReferenceOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.k8s_aws.SsmParameterSecretStore",
		reflect.TypeOf((*SsmParameterSecretStore)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addSecret", GoMethod: "AddSecret"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "cluster", GoGetter: "Cluster"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "manifest", GoGetter: "Manifest"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "namespace", GoGetter: "Namespace"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "secretStoreName", GoGetter: "SecretStoreName"},
			_jsii_.MemberProperty{JsiiProperty: "service", GoGetter: "Service"},
			_jsii_.MemberProperty{JsiiProperty: "serviceAccount", GoGetter: "ServiceAccount"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_SsmParameterSecretStore{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_AwsSecretStore)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.k8s_aws.SsmParameterSecretStoreProps",
		reflect.TypeOf((*SsmParameterSecretStoreProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.k8s_aws.TxtRegistry",
		reflect.TypeOf((*TxtRegistry)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "ownerId", GoGetter: "OwnerId"},
			_jsii_.MemberProperty{JsiiProperty: "prefix", GoGetter: "Prefix"},
			_jsii_.MemberProperty{JsiiProperty: "registryType", GoGetter: "RegistryType"},
		},
		func() interface{} {
			j := jsiiProxy_TxtRegistry{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IExternalDnsRegistry)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.k8s_aws.TxtRegistryOptions",
		reflect.TypeOf((*TxtRegistryOptions)(nil)).Elem(),
	)
}
