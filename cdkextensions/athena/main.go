package athena

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"cdk-extensions.athena.AddNamedQueryOptions",
		reflect.TypeOf((*AddNamedQueryOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.athena.AnalyticsEngine",
		reflect.TypeOf((*AnalyticsEngine)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_AnalyticsEngine{}
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.athena.AnalyticsEngineBindProps",
		reflect.TypeOf((*AnalyticsEngineBindProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.athena.AnalyticsEngineConfiguration",
		reflect.TypeOf((*AnalyticsEngineConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.athena.AnalyticsEngineOutputOptions",
		reflect.TypeOf((*AnalyticsEngineOutputOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.athena.AnalyticsEngineVersion",
		reflect.TypeOf((*AnalyticsEngineVersion)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
		},
		func() interface{} {
			return &jsiiProxy_AnalyticsEngineVersion{}
		},
	)
	_jsii_.RegisterClass(
		"cdk-extensions.athena.ApacheSparkEngineOptions",
		reflect.TypeOf((*ApacheSparkEngineOptions)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "encryptionKey", GoGetter: "EncryptionKey"},
			_jsii_.MemberProperty{JsiiProperty: "engineVersion", GoGetter: "EngineVersion"},
			_jsii_.MemberProperty{JsiiProperty: "output", GoGetter: "Output"},
			_jsii_.MemberProperty{JsiiProperty: "publishMetrics", GoGetter: "PublishMetrics"},
			_jsii_.MemberProperty{JsiiProperty: "role", GoGetter: "Role"},
		},
		func() interface{} {
			return &jsiiProxy_ApacheSparkEngineOptions{}
		},
	)
	_jsii_.RegisterClass(
		"cdk-extensions.athena.ApacheSparkEngineVersion",
		reflect.TypeOf((*ApacheSparkEngineVersion)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
		},
		func() interface{} {
			j := jsiiProxy_ApacheSparkEngineVersion{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_AnalyticsEngineVersion)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"cdk-extensions.athena.ApacheSparkOutputEncryption",
		reflect.TypeOf((*ApacheSparkOutputEncryption)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_ApacheSparkOutputEncryption{}
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.athena.AthenaResultEncryptionConfiguration",
		reflect.TypeOf((*AthenaResultEncryptionConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.athena.AthenaResultKmsEncryptionOptions",
		reflect.TypeOf((*AthenaResultKmsEncryptionOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.athena.AthenaSqlEngineOptions",
		reflect.TypeOf((*AthenaSqlEngineOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.athena.AthenaSqlEngineVersion",
		reflect.TypeOf((*AthenaSqlEngineVersion)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
		},
		func() interface{} {
			j := jsiiProxy_AthenaSqlEngineVersion{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_AnalyticsEngineVersion)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"cdk-extensions.athena.AthenaSqlOutputEncryption",
		reflect.TypeOf((*AthenaSqlOutputEncryption)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_AthenaSqlOutputEncryption{}
		},
	)
	_jsii_.RegisterInterface(
		"cdk-extensions.athena.IAnalyticsEngine",
		reflect.TypeOf((*IAnalyticsEngine)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			return &jsiiProxy_IAnalyticsEngine{}
		},
	)
	_jsii_.RegisterInterface(
		"cdk-extensions.athena.IAthenaResultEncryption",
		reflect.TypeOf((*IAthenaResultEncryption)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			return &jsiiProxy_IAthenaResultEncryption{}
		},
	)
	_jsii_.RegisterInterface(
		"cdk-extensions.athena.IWorkGroup",
		reflect.TypeOf((*IWorkGroup)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "workGroupArn", GoGetter: "WorkGroupArn"},
			_jsii_.MemberProperty{JsiiProperty: "workGroupCreationTime", GoGetter: "WorkGroupCreationTime"},
			_jsii_.MemberProperty{JsiiProperty: "workGroupEffectiveEngineVersion", GoGetter: "WorkGroupEffectiveEngineVersion"},
			_jsii_.MemberProperty{JsiiProperty: "workGroupName", GoGetter: "WorkGroupName"},
		},
		func() interface{} {
			return &jsiiProxy_IWorkGroup{}
		},
	)
	_jsii_.RegisterClass(
		"cdk-extensions.athena.NamedQuery",
		reflect.TypeOf((*NamedQuery)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "database", GoGetter: "Database"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "namedQueryId", GoGetter: "NamedQueryId"},
			_jsii_.MemberProperty{JsiiProperty: "namedQueryName", GoGetter: "NamedQueryName"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "queryString", GoGetter: "QueryString"},
			_jsii_.MemberProperty{JsiiProperty: "resource", GoGetter: "Resource"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "workGroup", GoGetter: "WorkGroup"},
		},
		func() interface{} {
			j := jsiiProxy_NamedQuery{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.athena.NamedQueryProps",
		reflect.TypeOf((*NamedQueryProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.athena.WorkGroup",
		reflect.TypeOf((*WorkGroup)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addNamedQuery", GoMethod: "AddNamedQuery"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "engine", GoGetter: "Engine"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "recursiveDelete", GoGetter: "RecursiveDelete"},
			_jsii_.MemberProperty{JsiiProperty: "resource", GoGetter: "Resource"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "state", GoGetter: "State"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "workGroupArn", GoGetter: "WorkGroupArn"},
			_jsii_.MemberProperty{JsiiProperty: "workGroupCreationTime", GoGetter: "WorkGroupCreationTime"},
			_jsii_.MemberProperty{JsiiProperty: "workGroupEffectiveEngineVersion", GoGetter: "WorkGroupEffectiveEngineVersion"},
			_jsii_.MemberProperty{JsiiProperty: "workGroupName", GoGetter: "WorkGroupName"},
		},
		func() interface{} {
			j := jsiiProxy_WorkGroup{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IWorkGroup)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.athena.WorkGroupAttributes",
		reflect.TypeOf((*WorkGroupAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.athena.WorkGroupOptions",
		reflect.TypeOf((*WorkGroupOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.athena.WorkGroupProps",
		reflect.TypeOf((*WorkGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.athena.WorkGroupState",
		reflect.TypeOf((*WorkGroupState)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
		},
		func() interface{} {
			return &jsiiProxy_WorkGroupState{}
		},
	)
}
