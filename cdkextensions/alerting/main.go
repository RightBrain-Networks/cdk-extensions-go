package alerting

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"cdk-extensions.alerting.AddReferenceProps",
		reflect.TypeOf((*AddReferenceProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.alerting.AppendDelimiter",
		reflect.TypeOf((*AppendDelimiter)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "delimiter", GoGetter: "Delimiter"},
		},
		func() interface{} {
			return &jsiiProxy_AppendDelimiter{}
		},
	)
	_jsii_.RegisterClass(
		"cdk-extensions.alerting.DescriptionBuilder",
		reflect.TypeOf((*DescriptionBuilder)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addIterator", GoMethod: "AddIterator"},
			_jsii_.MemberMethod{JsiiMethod: "addReference", GoMethod: "AddReference"},
			_jsii_.MemberMethod{JsiiMethod: "addSection", GoMethod: "AddSection"},
			_jsii_.MemberMethod{JsiiMethod: "buildId", GoMethod: "BuildId"},
			_jsii_.MemberProperty{JsiiProperty: "classifier", GoGetter: "Classifier"},
			_jsii_.MemberProperty{JsiiProperty: "initialDescription", GoGetter: "InitialDescription"},
			_jsii_.MemberMethod{JsiiMethod: "initialize", GoMethod: "Initialize"},
			_jsii_.MemberProperty{JsiiProperty: "initializeNode", GoGetter: "InitializeNode"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "registerBuilder", GoMethod: "RegisterBuilder"},
			_jsii_.MemberMethod{JsiiMethod: "registerChainable", GoMethod: "RegisterChainable"},
			_jsii_.MemberMethod{JsiiMethod: "render", GoMethod: "Render"},
			_jsii_.MemberMethod{JsiiMethod: "setDelimiter", GoMethod: "SetDelimiter"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "write", GoMethod: "Write"},
		},
		func() interface{} {
			j := jsiiProxy_DescriptionBuilder{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IDelayedChainable)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IDescriptionBuilderComponent)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"cdk-extensions.alerting.DescriptionBuilderIterator",
		reflect.TypeOf((*DescriptionBuilderIterator)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addIterator", GoMethod: "AddIterator"},
			_jsii_.MemberMethod{JsiiMethod: "addReference", GoMethod: "AddReference"},
			_jsii_.MemberProperty{JsiiProperty: "arrayRef", GoGetter: "ArrayRef"},
			_jsii_.MemberMethod{JsiiMethod: "buildId", GoMethod: "BuildId"},
			_jsii_.MemberProperty{JsiiProperty: "classifier", GoGetter: "Classifier"},
			_jsii_.MemberProperty{JsiiProperty: "fieldDelimiter", GoGetter: "FieldDelimiter"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "recordDelimiter", GoGetter: "RecordDelimiter"},
			_jsii_.MemberMethod{JsiiMethod: "registerBuilder", GoMethod: "RegisterBuilder"},
			_jsii_.MemberMethod{JsiiMethod: "registerChainable", GoMethod: "RegisterChainable"},
			_jsii_.MemberMethod{JsiiMethod: "render", GoMethod: "Render"},
			_jsii_.MemberProperty{JsiiProperty: "resultPath", GoGetter: "ResultPath"},
			_jsii_.MemberProperty{JsiiProperty: "sectionDelimiter", GoGetter: "SectionDelimiter"},
			_jsii_.MemberMethod{JsiiMethod: "setDelimiter", GoMethod: "SetDelimiter"},
			_jsii_.MemberProperty{JsiiProperty: "title", GoGetter: "Title"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "write", GoMethod: "Write"},
		},
		func() interface{} {
			j := jsiiProxy_DescriptionBuilderIterator{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IDelayedChainable)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IDescriptionBuilderComponent)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.alerting.DescriptionBuilderIteratorProps",
		reflect.TypeOf((*DescriptionBuilderIteratorProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.alerting.DescriptionBuilderProps",
		reflect.TypeOf((*DescriptionBuilderProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.alerting.DescriptionBuilderSection",
		reflect.TypeOf((*DescriptionBuilderSection)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addIterator", GoMethod: "AddIterator"},
			_jsii_.MemberMethod{JsiiMethod: "addReference", GoMethod: "AddReference"},
			_jsii_.MemberMethod{JsiiMethod: "addReferenceCheck", GoMethod: "AddReferenceCheck"},
			_jsii_.MemberMethod{JsiiMethod: "buildId", GoMethod: "BuildId"},
			_jsii_.MemberProperty{JsiiProperty: "classifier", GoGetter: "Classifier"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "refs", GoGetter: "Refs"},
			_jsii_.MemberMethod{JsiiMethod: "registerBuilder", GoMethod: "RegisterBuilder"},
			_jsii_.MemberMethod{JsiiMethod: "registerChainable", GoMethod: "RegisterChainable"},
			_jsii_.MemberMethod{JsiiMethod: "render", GoMethod: "Render"},
			_jsii_.MemberMethod{JsiiMethod: "setDelimiter", GoMethod: "SetDelimiter"},
			_jsii_.MemberProperty{JsiiProperty: "title", GoGetter: "Title"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "write", GoMethod: "Write"},
		},
		func() interface{} {
			j := jsiiProxy_DescriptionBuilderSection{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IDelayedChainable)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IDescriptionBuilderComponent)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.alerting.DescriptionBuilderSectionProps",
		reflect.TypeOf((*DescriptionBuilderSectionProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.alerting.EcrImageScanSeverity",
		reflect.TypeOf((*EcrImageScanSeverity)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "priority", GoGetter: "Priority"},
			_jsii_.MemberProperty{JsiiProperty: "standardized", GoGetter: "Standardized"},
		},
		func() interface{} {
			return &jsiiProxy_EcrImageScanSeverity{}
		},
	)
	_jsii_.RegisterClass(
		"cdk-extensions.alerting.EcrImageScanSeverityConfiguration",
		reflect.TypeOf((*EcrImageScanSeverityConfiguration)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_EcrImageScanSeverityConfiguration{}
		},
	)
	_jsii_.RegisterClass(
		"cdk-extensions.alerting.EcrScanFinding",
		reflect.TypeOf((*EcrScanFinding)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberMethod{JsiiMethod: "buildDescription", GoMethod: "BuildDescription"},
			_jsii_.MemberMethod{JsiiMethod: "buildLogging", GoMethod: "BuildLogging"},
			_jsii_.MemberMethod{JsiiMethod: "buildSummary", GoMethod: "BuildSummary"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "handler", GoGetter: "Handler"},
			_jsii_.MemberProperty{JsiiProperty: "logging", GoGetter: "Logging"},
			_jsii_.MemberProperty{JsiiProperty: "matchType", GoGetter: "MatchType"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "registerIssueTrigger", GoMethod: "RegisterIssueTrigger"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "timeout", GoGetter: "Timeout"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "triggers", GoGetter: "Triggers"},
		},
		func() interface{} {
			j := jsiiProxy_EcrScanFinding{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IssuePluginBase)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IIssueParser)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.alerting.EcrScanFindingEventOptions",
		reflect.TypeOf((*EcrScanFindingEventOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.alerting.EcrScanFindingProps",
		reflect.TypeOf((*EcrScanFindingProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.alerting.GuardDutyFinding",
		reflect.TypeOf((*GuardDutyFinding)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDefaultTrigger", GoMethod: "AddDefaultTrigger"},
			_jsii_.MemberMethod{JsiiMethod: "addSectionField", GoMethod: "AddSectionField"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberMethod{JsiiMethod: "buildLogging", GoMethod: "BuildLogging"},
			_jsii_.MemberMethod{JsiiMethod: "buildSeverityMap", GoMethod: "BuildSeverityMap"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "handler", GoGetter: "Handler"},
			_jsii_.MemberProperty{JsiiProperty: "logging", GoGetter: "Logging"},
			_jsii_.MemberProperty{JsiiProperty: "matchType", GoGetter: "MatchType"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "registerIssueTrigger", GoMethod: "RegisterIssueTrigger"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "timeout", GoGetter: "Timeout"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "triggers", GoGetter: "Triggers"},
		},
		func() interface{} {
			j := jsiiProxy_GuardDutyFinding{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IssueParserPluginBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.alerting.GuardDutyFindingProps",
		reflect.TypeOf((*GuardDutyFindingProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.alerting.GuardDutyFindingRuleOptions",
		reflect.TypeOf((*GuardDutyFindingRuleOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.alerting.GuardDutySeverity",
		reflect.TypeOf((*GuardDutySeverity)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "buildCondition", GoMethod: "BuildCondition"},
			_jsii_.MemberProperty{JsiiProperty: "lowerBound", GoGetter: "LowerBound"},
			_jsii_.MemberProperty{JsiiProperty: "standardized", GoGetter: "Standardized"},
			_jsii_.MemberProperty{JsiiProperty: "upperBound", GoGetter: "UpperBound"},
		},
		func() interface{} {
			return &jsiiProxy_GuardDutySeverity{}
		},
	)
	_jsii_.RegisterInterface(
		"cdk-extensions.alerting.IDelayedChainable",
		reflect.TypeOf((*IDelayedChainable)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "render", GoMethod: "Render"},
		},
		func() interface{} {
			return &jsiiProxy_IDelayedChainable{}
		},
	)
	_jsii_.RegisterInterface(
		"cdk-extensions.alerting.IDescriptionBuilderComponent",
		reflect.TypeOf((*IDescriptionBuilderComponent)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addIterator", GoMethod: "AddIterator"},
			_jsii_.MemberMethod{JsiiMethod: "addReference", GoMethod: "AddReference"},
			_jsii_.MemberProperty{JsiiProperty: "classifier", GoGetter: "Classifier"},
			_jsii_.MemberMethod{JsiiMethod: "render", GoMethod: "Render"},
			_jsii_.MemberMethod{JsiiMethod: "setDelimiter", GoMethod: "SetDelimiter"},
			_jsii_.MemberMethod{JsiiMethod: "write", GoMethod: "Write"},
		},
		func() interface{} {
			j := jsiiProxy_IDescriptionBuilderComponent{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IDelayedChainable)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"cdk-extensions.alerting.IEcrImageScanSeverityConfiguration",
		reflect.TypeOf((*IEcrImageScanSeverityConfiguration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "levels", GoGetter: "Levels"},
		},
		func() interface{} {
			return &jsiiProxy_IEcrImageScanSeverityConfiguration{}
		},
	)
	_jsii_.RegisterInterface(
		"cdk-extensions.alerting.IIssueHandler",
		reflect.TypeOf((*IIssueHandler)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "handler", GoGetter: "Handler"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IIssueHandler{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"cdk-extensions.alerting.IIssueParser",
		reflect.TypeOf((*IIssueParser)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "handler", GoGetter: "Handler"},
			_jsii_.MemberProperty{JsiiProperty: "matchType", GoGetter: "MatchType"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IIssueParser{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"cdk-extensions.alerting.ISecurityHubSeverityConfiguration",
		reflect.TypeOf((*ISecurityHubSeverityConfiguration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "levels", GoGetter: "Levels"},
		},
		func() interface{} {
			return &jsiiProxy_ISecurityHubSeverityConfiguration{}
		},
	)
	_jsii_.RegisterClass(
		"cdk-extensions.alerting.InspectorFinding",
		reflect.TypeOf((*InspectorFinding)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDefaultTrigger", GoMethod: "AddDefaultTrigger"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberMethod{JsiiMethod: "buildDescription", GoMethod: "BuildDescription"},
			_jsii_.MemberMethod{JsiiMethod: "buildLogging", GoMethod: "BuildLogging"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "handler", GoGetter: "Handler"},
			_jsii_.MemberProperty{JsiiProperty: "logging", GoGetter: "Logging"},
			_jsii_.MemberProperty{JsiiProperty: "matchType", GoGetter: "MatchType"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "registerIssueTrigger", GoMethod: "RegisterIssueTrigger"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "timeout", GoGetter: "Timeout"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "triggers", GoGetter: "Triggers"},
		},
		func() interface{} {
			j := jsiiProxy_InspectorFinding{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IssueParserPluginBase)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IIssueParser)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.alerting.InspectorFindingEventOptions",
		reflect.TypeOf((*InspectorFindingEventOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.alerting.InspectorFindingProps",
		reflect.TypeOf((*InspectorFindingProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.alerting.InspectorSeverity",
		reflect.TypeOf((*InspectorSeverity)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "buildCondition", GoMethod: "BuildCondition"},
			_jsii_.MemberProperty{JsiiProperty: "original", GoGetter: "Original"},
			_jsii_.MemberProperty{JsiiProperty: "priority", GoGetter: "Priority"},
			_jsii_.MemberProperty{JsiiProperty: "standardized", GoGetter: "Standardized"},
		},
		func() interface{} {
			return &jsiiProxy_InspectorSeverity{}
		},
	)
	_jsii_.RegisterClass(
		"cdk-extensions.alerting.IssueHander",
		reflect.TypeOf((*IssueHander)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_IssueHander{}
		},
	)
	_jsii_.RegisterClass(
		"cdk-extensions.alerting.IssueHandlerOverride",
		reflect.TypeOf((*IssueHandlerOverride)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "handler", GoGetter: "Handler"},
			_jsii_.MemberProperty{JsiiProperty: "overrides", GoGetter: "Overrides"},
		},
		func() interface{} {
			return &jsiiProxy_IssueHandlerOverride{}
		},
	)
	_jsii_.RegisterClass(
		"cdk-extensions.alerting.IssueManager",
		reflect.TypeOf((*IssueManager)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addEventRules", GoMethod: "AddEventRules"},
			_jsii_.MemberMethod{JsiiMethod: "addHandler", GoMethod: "AddHandler"},
			_jsii_.MemberMethod{JsiiMethod: "addIssueParser", GoMethod: "AddIssueParser"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "timeout", GoGetter: "Timeout"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_IssueManager{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.alerting.IssueManagerProps",
		reflect.TypeOf((*IssueManagerProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.alerting.IssueParserPluginBase",
		reflect.TypeOf((*IssueParserPluginBase)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDefaultTrigger", GoMethod: "AddDefaultTrigger"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberMethod{JsiiMethod: "buildLogging", GoMethod: "BuildLogging"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "handler", GoGetter: "Handler"},
			_jsii_.MemberProperty{JsiiProperty: "logging", GoGetter: "Logging"},
			_jsii_.MemberProperty{JsiiProperty: "matchType", GoGetter: "MatchType"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "timeout", GoGetter: "Timeout"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "triggers", GoGetter: "Triggers"},
		},
		func() interface{} {
			j := jsiiProxy_IssueParserPluginBase{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IssuePluginBase)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IIssueParser)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.alerting.IssueParserPluginBaseProps",
		reflect.TypeOf((*IssueParserPluginBaseProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.alerting.IssuePluginBase",
		reflect.TypeOf((*IssuePluginBase)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "buildLogging", GoMethod: "BuildLogging"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "logging", GoGetter: "Logging"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_IssuePluginBase{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.alerting.IssuePluginBaseProps",
		reflect.TypeOf((*IssuePluginBaseProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.alerting.IssueTrigger",
		reflect.TypeOf((*IssueTrigger)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "eventPattern", GoGetter: "EventPattern"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "overrides", GoGetter: "Overrides"},
			_jsii_.MemberProperty{JsiiProperty: "parser", GoGetter: "Parser"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_IssueTrigger{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.alerting.IssueTriggerProps",
		reflect.TypeOf((*IssueTriggerProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.alerting.JiraTicket",
		reflect.TypeOf((*JiraTicket)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "apiDestination", GoGetter: "ApiDestination"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "assignee", GoGetter: "Assignee"},
			_jsii_.MemberMethod{JsiiMethod: "buildEventOverrides", GoMethod: "BuildEventOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "buildLogging", GoMethod: "BuildLogging"},
			_jsii_.MemberMethod{JsiiMethod: "buildSeverityMap", GoMethod: "BuildSeverityMap"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "credentials", GoGetter: "Credentials"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "eventBus", GoGetter: "EventBus"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "handler", GoGetter: "Handler"},
			_jsii_.MemberProperty{JsiiProperty: "issueType", GoGetter: "IssueType"},
			_jsii_.MemberProperty{JsiiProperty: "jiraUrl", GoGetter: "JiraUrl"},
			_jsii_.MemberProperty{JsiiProperty: "logging", GoGetter: "Logging"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "priorityMap", GoGetter: "PriorityMap"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "timeout", GoGetter: "Timeout"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_JiraTicket{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IssuePluginBase)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IIssueHandler)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.alerting.JiraTicketOverrideOptions",
		reflect.TypeOf((*JiraTicketOverrideOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.alerting.JiraTicketPriorityMap",
		reflect.TypeOf((*JiraTicketPriorityMap)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.alerting.JiraTicketProps",
		reflect.TypeOf((*JiraTicketProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.alerting.OpenSearchEvent",
		reflect.TypeOf((*OpenSearchEvent)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDefaultTrigger", GoMethod: "AddDefaultTrigger"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberMethod{JsiiMethod: "buildDescription", GoMethod: "BuildDescription"},
			_jsii_.MemberMethod{JsiiMethod: "buildLogging", GoMethod: "BuildLogging"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "handler", GoGetter: "Handler"},
			_jsii_.MemberProperty{JsiiProperty: "logging", GoGetter: "Logging"},
			_jsii_.MemberProperty{JsiiProperty: "matchType", GoGetter: "MatchType"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "registerIssueTrigger", GoMethod: "RegisterIssueTrigger"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "timeout", GoGetter: "Timeout"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "triggers", GoGetter: "Triggers"},
		},
		func() interface{} {
			j := jsiiProxy_OpenSearchEvent{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IssueParserPluginBase)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IIssueParser)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.alerting.OpenSearchEventProps",
		reflect.TypeOf((*OpenSearchEventProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.alerting.OpenSearchEventRuleOptions",
		reflect.TypeOf((*OpenSearchEventRuleOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.alerting.OpenSearchEventSeverity",
		reflect.TypeOf((*OpenSearchEventSeverity)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "buildCondition", GoMethod: "BuildCondition"},
			_jsii_.MemberProperty{JsiiProperty: "original", GoGetter: "Original"},
			_jsii_.MemberProperty{JsiiProperty: "priority", GoGetter: "Priority"},
			_jsii_.MemberProperty{JsiiProperty: "standardized", GoGetter: "Standardized"},
		},
		func() interface{} {
			return &jsiiProxy_OpenSearchEventSeverity{}
		},
	)
	_jsii_.RegisterClass(
		"cdk-extensions.alerting.OpenSearchEventType",
		reflect.TypeOf((*OpenSearchEventType)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "detailType", GoGetter: "DetailType"},
			_jsii_.MemberProperty{JsiiProperty: "eventName", GoGetter: "EventName"},
		},
		func() interface{} {
			return &jsiiProxy_OpenSearchEventType{}
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.alerting.OpenSearchEventTypeProps",
		reflect.TypeOf((*OpenSearchEventTypeProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.alerting.SecurityHubFinding",
		reflect.TypeOf((*SecurityHubFinding)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberMethod{JsiiMethod: "buildDescription", GoMethod: "BuildDescription"},
			_jsii_.MemberMethod{JsiiMethod: "buildLogging", GoMethod: "BuildLogging"},
			_jsii_.MemberMethod{JsiiMethod: "buildRemediation", GoMethod: "BuildRemediation"},
			_jsii_.MemberMethod{JsiiMethod: "buildResources", GoMethod: "BuildResources"},
			_jsii_.MemberMethod{JsiiMethod: "buildSeverityMap", GoMethod: "BuildSeverityMap"},
			_jsii_.MemberMethod{JsiiMethod: "buildUrl", GoMethod: "BuildUrl"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "handler", GoGetter: "Handler"},
			_jsii_.MemberProperty{JsiiProperty: "logging", GoGetter: "Logging"},
			_jsii_.MemberProperty{JsiiProperty: "matchType", GoGetter: "MatchType"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "registerIssueTrigger", GoMethod: "RegisterIssueTrigger"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "timeout", GoGetter: "Timeout"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "triggers", GoGetter: "Triggers"},
		},
		func() interface{} {
			j := jsiiProxy_SecurityHubFinding{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IssuePluginBase)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IIssueParser)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.alerting.SecurityHubFindingEventOptions",
		reflect.TypeOf((*SecurityHubFindingEventOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.alerting.SecurityHubFindingProps",
		reflect.TypeOf((*SecurityHubFindingProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.alerting.SecurityHubSeverity",
		reflect.TypeOf((*SecurityHubSeverity)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "lowerBound", GoGetter: "LowerBound"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "standardized", GoGetter: "Standardized"},
			_jsii_.MemberProperty{JsiiProperty: "upperBound", GoGetter: "UpperBound"},
		},
		func() interface{} {
			return &jsiiProxy_SecurityHubSeverity{}
		},
	)
	_jsii_.RegisterClass(
		"cdk-extensions.alerting.SecurityHubSeverityConfiguration",
		reflect.TypeOf((*SecurityHubSeverityConfiguration)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_SecurityHubSeverityConfiguration{}
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.alerting.SetDelimiterProps",
		reflect.TypeOf((*SetDelimiterProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.alerting.StateMachineLogging",
		reflect.TypeOf((*StateMachineLogging)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.alerting.WriteProps",
		reflect.TypeOf((*WriteProps)(nil)).Elem(),
	)
}
