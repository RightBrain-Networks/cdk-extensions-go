package aps

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"cdk-extensions.aps.AlertManagerConfiguration",
		reflect.TypeOf((*AlertManagerConfiguration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addInhibitRule", GoMethod: "AddInhibitRule"},
			_jsii_.MemberMethod{JsiiMethod: "addReciever", GoMethod: "AddReciever"},
			_jsii_.MemberMethod{JsiiMethod: "addTemplate", GoMethod: "AddTemplate"},
			_jsii_.MemberMethod{JsiiMethod: "addTimeInterval", GoMethod: "AddTimeInterval"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "defaultReceiver", GoGetter: "DefaultReceiver"},
			_jsii_.MemberProperty{JsiiProperty: "defaultRoute", GoGetter: "DefaultRoute"},
			_jsii_.MemberProperty{JsiiProperty: "defaultTopic", GoGetter: "DefaultTopic"},
			_jsii_.MemberMethod{JsiiMethod: "generateTemplateName", GoMethod: "GenerateTemplateName"},
			_jsii_.MemberProperty{JsiiProperty: "inhibitRules", GoGetter: "InhibitRules"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "receivers", GoGetter: "Receivers"},
			_jsii_.MemberMethod{JsiiMethod: "renderAlertManagerConfig", GoMethod: "RenderAlertManagerConfig"},
			_jsii_.MemberMethod{JsiiMethod: "renderTemplates", GoMethod: "RenderTemplates"},
			_jsii_.MemberProperty{JsiiProperty: "templates", GoGetter: "Templates"},
			_jsii_.MemberProperty{JsiiProperty: "timeIntervals", GoGetter: "TimeIntervals"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AlertManagerConfiguration{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IAlertManagerConfiguration)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.aps.AlertManagerConfigurationDetails",
		reflect.TypeOf((*AlertManagerConfigurationDetails)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.aps.AlertManagerConfigurationProps",
		reflect.TypeOf((*AlertManagerConfigurationProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.aps.AlertManagerDestination",
		reflect.TypeOf((*AlertManagerDestination)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_AlertManagerDestination{}
		},
	)
	_jsii_.RegisterClass(
		"cdk-extensions.aps.AlertManagerDestinationCategory",
		reflect.TypeOf((*AlertManagerDestinationCategory)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
		},
		func() interface{} {
			return &jsiiProxy_AlertManagerDestinationCategory{}
		},
	)
	_jsii_.RegisterClass(
		"cdk-extensions.aps.AlertManagerInhibitRule",
		reflect.TypeOf((*AlertManagerInhibitRule)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addEqualLabel", GoMethod: "AddEqualLabel"},
			_jsii_.MemberMethod{JsiiMethod: "addSourceMatcher", GoMethod: "AddSourceMatcher"},
			_jsii_.MemberMethod{JsiiMethod: "addTargetMatcher", GoMethod: "AddTargetMatcher"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "equalLabels", GoGetter: "EqualLabels"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "sourceMatchers", GoGetter: "SourceMatchers"},
			_jsii_.MemberProperty{JsiiProperty: "targetMatchers", GoGetter: "TargetMatchers"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AlertManagerInhibitRule{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"cdk-extensions.aps.AlertManagerInhibitRuleProps",
		reflect.TypeOf((*AlertManagerInhibitRuleProps)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "equalLabels", GoGetter: "EqualLabels"},
			_jsii_.MemberProperty{JsiiProperty: "sourceMatchers", GoGetter: "SourceMatchers"},
			_jsii_.MemberProperty{JsiiProperty: "targetMatchers", GoGetter: "TargetMatchers"},
		},
		func() interface{} {
			return &jsiiProxy_AlertManagerInhibitRuleProps{}
		},
	)
	_jsii_.RegisterClass(
		"cdk-extensions.aps.AlertManagerMatcher",
		reflect.TypeOf((*AlertManagerMatcher)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "expression", GoGetter: "Expression"},
		},
		func() interface{} {
			return &jsiiProxy_AlertManagerMatcher{}
		},
	)
	_jsii_.RegisterClass(
		"cdk-extensions.aps.AlertManagerReceiver",
		reflect.TypeOf((*AlertManagerReceiver)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDestination", GoMethod: "AddDestination"},
			_jsii_.MemberMethod{JsiiMethod: "addSnsTopic", GoMethod: "AddSnsTopic"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "destinations", GoGetter: "Destinations"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AlertManagerReceiver{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.aps.AlertManagerReceiverProps",
		reflect.TypeOf((*AlertManagerReceiverProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.aps.AlertManagerRoute",
		reflect.TypeOf((*AlertManagerRoute)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "activeTimeIntervals", GoGetter: "ActiveTimeIntervals"},
			_jsii_.MemberMethod{JsiiMethod: "addActiveTimeInterval", GoMethod: "AddActiveTimeInterval"},
			_jsii_.MemberMethod{JsiiMethod: "addChild", GoMethod: "AddChild"},
			_jsii_.MemberMethod{JsiiMethod: "addGroupByLabel", GoMethod: "AddGroupByLabel"},
			_jsii_.MemberMethod{JsiiMethod: "addMatcher", GoMethod: "AddMatcher"},
			_jsii_.MemberMethod{JsiiMethod: "addMuteTimeInterval", GoMethod: "AddMuteTimeInterval"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "children", GoGetter: "Children"},
			_jsii_.MemberProperty{JsiiProperty: "continueMatching", GoGetter: "ContinueMatching"},
			_jsii_.MemberProperty{JsiiProperty: "groupByLabels", GoGetter: "GroupByLabels"},
			_jsii_.MemberProperty{JsiiProperty: "groupInterval", GoGetter: "GroupInterval"},
			_jsii_.MemberProperty{JsiiProperty: "groupWait", GoGetter: "GroupWait"},
			_jsii_.MemberProperty{JsiiProperty: "matchers", GoGetter: "Matchers"},
			_jsii_.MemberProperty{JsiiProperty: "muteTimeIntervals", GoGetter: "MuteTimeIntervals"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "receiver", GoGetter: "Receiver"},
			_jsii_.MemberProperty{JsiiProperty: "repeatInterval", GoGetter: "RepeatInterval"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AlertManagerRoute{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.aps.AlertManagerRouteProps",
		reflect.TypeOf((*AlertManagerRouteProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.aps.AlertManagerSnsDestination",
		reflect.TypeOf((*AlertManagerSnsDestination)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addAttribute", GoMethod: "AddAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "apiUrl", GoGetter: "ApiUrl"},
			_jsii_.MemberProperty{JsiiProperty: "attributes", GoGetter: "Attributes"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "category", GoGetter: "Category"},
			_jsii_.MemberProperty{JsiiProperty: "message", GoGetter: "Message"},
			_jsii_.MemberProperty{JsiiProperty: "sendResolved", GoGetter: "SendResolved"},
			_jsii_.MemberProperty{JsiiProperty: "subject", GoGetter: "Subject"},
			_jsii_.MemberProperty{JsiiProperty: "topic", GoGetter: "Topic"},
		},
		func() interface{} {
			j := jsiiProxy_AlertManagerSnsDestination{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IAlertManagerDestination)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.aps.AlertManagerSnsDestinationOptions",
		reflect.TypeOf((*AlertManagerSnsDestinationOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.aps.AlertManagerSnsDestinationProps",
		reflect.TypeOf((*AlertManagerSnsDestinationProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.aps.AlertManagerTemplate",
		reflect.TypeOf((*AlertManagerTemplate)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "content", GoGetter: "Content"},
		},
		func() interface{} {
			return &jsiiProxy_AlertManagerTemplate{}
		},
	)
	_jsii_.RegisterClass(
		"cdk-extensions.aps.AlertingRule",
		reflect.TypeOf((*AlertingRule)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addAnnotation", GoMethod: "AddAnnotation"},
			_jsii_.MemberMethod{JsiiMethod: "addLabel", GoMethod: "AddLabel"},
			_jsii_.MemberProperty{JsiiProperty: "alert", GoGetter: "Alert"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "expression", GoGetter: "Expression"},
			_jsii_.MemberProperty{JsiiProperty: "period", GoGetter: "Period"},
		},
		func() interface{} {
			j := jsiiProxy_AlertingRule{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IPrometheusRule)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.aps.AlertingRuleProps",
		reflect.TypeOf((*AlertingRuleProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.aps.DayOfMonthRange",
		reflect.TypeOf((*DayOfMonthRange)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.aps.DefaultRouteOptions",
		reflect.TypeOf((*DefaultRouteOptions)(nil)).Elem(),
	)
	_jsii_.RegisterInterface(
		"cdk-extensions.aps.IAlertManagerConfiguration",
		reflect.TypeOf((*IAlertManagerConfiguration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			return &jsiiProxy_IAlertManagerConfiguration{}
		},
	)
	_jsii_.RegisterInterface(
		"cdk-extensions.aps.IAlertManagerDestination",
		reflect.TypeOf((*IAlertManagerDestination)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "category", GoGetter: "Category"},
		},
		func() interface{} {
			return &jsiiProxy_IAlertManagerDestination{}
		},
	)
	_jsii_.RegisterInterface(
		"cdk-extensions.aps.IPrometheusRule",
		reflect.TypeOf((*IPrometheusRule)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			return &jsiiProxy_IPrometheusRule{}
		},
	)
	_jsii_.RegisterInterface(
		"cdk-extensions.aps.IRuleGroupConfiguration",
		reflect.TypeOf((*IRuleGroupConfiguration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
		},
		func() interface{} {
			j := jsiiProxy_IRuleGroupConfiguration{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"cdk-extensions.aps.IRuleGroupsNamespace",
		reflect.TypeOf((*IRuleGroupsNamespace)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "rulesGroupsNamespaceArn", GoGetter: "RulesGroupsNamespaceArn"},
		},
		func() interface{} {
			return &jsiiProxy_IRuleGroupsNamespace{}
		},
	)
	_jsii_.RegisterInterface(
		"cdk-extensions.aps.IWorkspace",
		reflect.TypeOf((*IWorkspace)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "workspaceArn", GoGetter: "WorkspaceArn"},
			_jsii_.MemberProperty{JsiiProperty: "workspaceId", GoGetter: "WorkspaceId"},
			_jsii_.MemberProperty{JsiiProperty: "workspacePrometheusEndpoint", GoGetter: "WorkspacePrometheusEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "workspaceQueryUrl", GoGetter: "WorkspaceQueryUrl"},
			_jsii_.MemberProperty{JsiiProperty: "workspaceRemoteWriteUrl", GoGetter: "WorkspaceRemoteWriteUrl"},
		},
		func() interface{} {
			return &jsiiProxy_IWorkspace{}
		},
	)
	_jsii_.RegisterEnum(
		"cdk-extensions.aps.MatchOperator",
		reflect.TypeOf((*MatchOperator)(nil)).Elem(),
		map[string]interface{}{
			"EQUALS": MatchOperator_EQUALS,
			"NOT_EQUALS": MatchOperator_NOT_EQUALS,
			"RE_EQUALS": MatchOperator_RE_EQUALS,
			"RE_NOT_EQUALS": MatchOperator_RE_NOT_EQUALS,
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.aps.MonthRange",
		reflect.TypeOf((*MonthRange)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.aps.PrometheusRule",
		reflect.TypeOf((*PrometheusRule)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_PrometheusRule{}
		},
	)
	_jsii_.RegisterClass(
		"cdk-extensions.aps.RecordingRule",
		reflect.TypeOf((*RecordingRule)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addLabel", GoMethod: "AddLabel"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "expression", GoGetter: "Expression"},
			_jsii_.MemberProperty{JsiiProperty: "labels", GoGetter: "Labels"},
			_jsii_.MemberProperty{JsiiProperty: "record", GoGetter: "Record"},
		},
		func() interface{} {
			j := jsiiProxy_RecordingRule{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IPrometheusRule)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.aps.RecordingRuleProps",
		reflect.TypeOf((*RecordingRuleProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.aps.RuleGroup",
		reflect.TypeOf((*RuleGroup)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addAlertingRule", GoMethod: "AddAlertingRule"},
			_jsii_.MemberMethod{JsiiMethod: "addRecordingRule", GoMethod: "AddRecordingRule"},
			_jsii_.MemberMethod{JsiiMethod: "addRule", GoMethod: "AddRule"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "interval", GoGetter: "Interval"},
			_jsii_.MemberProperty{JsiiProperty: "limit", GoGetter: "Limit"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_RuleGroup{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"cdk-extensions.aps.RuleGroupConfiguration",
		reflect.TypeOf((*RuleGroupConfiguration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addRuleGroup", GoMethod: "AddRuleGroup"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "ruleGroups", GoGetter: "RuleGroups"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_RuleGroupConfiguration{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IRuleGroupConfiguration)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.aps.RuleGroupConfigurationDetails",
		reflect.TypeOf((*RuleGroupConfigurationDetails)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.aps.RuleGroupConfigurationProps",
		reflect.TypeOf((*RuleGroupConfigurationProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.aps.RuleGroupProps",
		reflect.TypeOf((*RuleGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.aps.RuleGroupsNamespace",
		reflect.TypeOf((*RuleGroupsNamespace)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addRuleGroup", GoMethod: "AddRuleGroup"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "configuration", GoGetter: "Configuration"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "resource", GoGetter: "Resource"},
			_jsii_.MemberProperty{JsiiProperty: "rulesGroupsNamespaceArn", GoGetter: "RulesGroupsNamespaceArn"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "workspace", GoGetter: "Workspace"},
		},
		func() interface{} {
			j := jsiiProxy_RuleGroupsNamespace{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IRuleGroupsNamespace)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.aps.RuleGroupsNamespaceProps",
		reflect.TypeOf((*RuleGroupsNamespaceProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.aps.TimeInterval",
		reflect.TypeOf((*TimeInterval)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addInterval", GoMethod: "AddInterval"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "intervals", GoGetter: "Intervals"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_TimeInterval{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"cdk-extensions.aps.TimeIntervalEntry",
		reflect.TypeOf((*TimeIntervalEntry)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDaysOfTheMonth", GoMethod: "AddDaysOfTheMonth"},
			_jsii_.MemberMethod{JsiiMethod: "addMonth", GoMethod: "AddMonth"},
			_jsii_.MemberMethod{JsiiMethod: "addTimes", GoMethod: "AddTimes"},
			_jsii_.MemberMethod{JsiiMethod: "addWeekday", GoMethod: "AddWeekday"},
			_jsii_.MemberMethod{JsiiMethod: "addYears", GoMethod: "AddYears"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "daysOfTheMonth", GoGetter: "DaysOfTheMonth"},
			_jsii_.MemberProperty{JsiiProperty: "months", GoGetter: "Months"},
			_jsii_.MemberProperty{JsiiProperty: "times", GoGetter: "Times"},
			_jsii_.MemberProperty{JsiiProperty: "timeZone", GoGetter: "TimeZone"},
			_jsii_.MemberProperty{JsiiProperty: "weekdays", GoGetter: "Weekdays"},
			_jsii_.MemberProperty{JsiiProperty: "years", GoGetter: "Years"},
		},
		func() interface{} {
			return &jsiiProxy_TimeIntervalEntry{}
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.aps.TimeIntervalEntryProps",
		reflect.TypeOf((*TimeIntervalEntryProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.aps.TimeIntervalProps",
		reflect.TypeOf((*TimeIntervalProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.aps.TimeRange",
		reflect.TypeOf((*TimeRange)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cdk-extensions.aps.Weekday",
		reflect.TypeOf((*Weekday)(nil)).Elem(),
		map[string]interface{}{
			"SUNDAY": Weekday_SUNDAY,
			"MONDAY": Weekday_MONDAY,
			"TUESDAY": Weekday_TUESDAY,
			"WEDNESDAY": Weekday_WEDNESDAY,
			"THURSDAY": Weekday_THURSDAY,
			"FRIDAY": Weekday_FRIDAY,
			"SATURDAY": Weekday_SATURDAY,
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.aps.WeekdayRange",
		reflect.TypeOf((*WeekdayRange)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.aps.Workspace",
		reflect.TypeOf((*Workspace)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "alerting", GoGetter: "Alerting"},
			_jsii_.MemberProperty{JsiiProperty: "alertManagerConfiguration", GoGetter: "AlertManagerConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "alertTopic", GoGetter: "AlertTopic"},
			_jsii_.MemberProperty{JsiiProperty: "alias", GoGetter: "Alias"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "logging", GoGetter: "Logging"},
			_jsii_.MemberProperty{JsiiProperty: "logGroup", GoGetter: "LogGroup"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "resource", GoGetter: "Resource"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "workspaceArn", GoGetter: "WorkspaceArn"},
			_jsii_.MemberProperty{JsiiProperty: "workspaceId", GoGetter: "WorkspaceId"},
			_jsii_.MemberProperty{JsiiProperty: "workspacePrometheusEndpoint", GoGetter: "WorkspacePrometheusEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "workspaceQueryUrl", GoGetter: "WorkspaceQueryUrl"},
			_jsii_.MemberProperty{JsiiProperty: "workspaceRemoteWriteUrl", GoGetter: "WorkspaceRemoteWriteUrl"},
		},
		func() interface{} {
			j := jsiiProxy_Workspace{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IWorkspace)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.aps.WorkspaceAlertingOptions",
		reflect.TypeOf((*WorkspaceAlertingOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.aps.WorkspaceAttributes",
		reflect.TypeOf((*WorkspaceAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.aps.WorkspaceLoggingOptions",
		reflect.TypeOf((*WorkspaceLoggingOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.aps.WorkspaceProps",
		reflect.TypeOf((*WorkspaceProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.aps.YearRange",
		reflect.TypeOf((*YearRange)(nil)).Elem(),
	)
}
